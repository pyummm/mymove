package server

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	idleTimeout   = 120 * time.Second // 2 minutes
	maxHeaderSize = 1 * 1000 * 1000   // 1 Megabyte
)

var supportedProtocols = []string{"h2"}

// ErrMissingCACert represents an error caused by server config that requires
// certificate verification, but is missing a CA certificate
var ErrMissingCACert = errors.New("missing required CA certificate")

// ErrUnparseableCACert represents an error cause by a misconfigured CA certificate
// that was unable to be parsed.
var ErrUnparseableCACert = errors.New("unable to parse CA certificate")

type serverFunc func(server *http.Server) error

// TLSCert encapsulates a public certificate and private key.
// Each are represented as a slice of bytes.
type TLSCert struct {
	CertPEMBlock []byte
	KeyPEMBlock  []byte
}

// Server represents an http or https listening server. HTTPS listeners support
// requiring client authentication with a provided CA.
type Server struct {
	CACertPEMBlock []byte
	ClientAuthType tls.ClientAuthType
	HTTPHandler    http.Handler
	ListenAddress  string
	Logger         *zap.Logger
	Port           string
	TLSCerts       []TLSCert
	PrefixToDrop   string // if len(PrefixToDrop) > 0, then all log lines with the prefix will be dropped
}

type loggerWriter struct {
	prefixToDrop []byte
	logFunc      func(msg string, fields ...zap.Field)
}

func (lw *loggerWriter) Write(p []byte) (int, error) {
	p = bytes.TrimSpace(p)
	if len(lw.prefixToDrop) > 0 {
		if reflect.DeepEqual(p[:len(lw.prefixToDrop)], lw.prefixToDrop) {
			return 0, nil
		}
	}
	lw.logFunc(string(p))
	return len(p), nil
}

func levelToFunc(logger *zap.Logger, lvl zapcore.Level) (func(string, ...zap.Field), error) {
	switch lvl {
	case zap.DebugLevel:
		return logger.Debug, nil
	case zap.InfoLevel:
		return logger.Info, nil
	case zap.WarnLevel:
		return logger.Warn, nil
	case zap.ErrorLevel:
		return logger.Error, nil
	case zap.DPanicLevel:
		return logger.DPanic, nil
	case zap.PanicLevel:
		return logger.Panic, nil
	case zap.FatalLevel:
		return logger.Fatal, nil
	}
	return nil, fmt.Errorf("unrecognized level: %q", lvl)
}

func newStdLogAt(l *zap.Logger, level zapcore.Level, prefixToDrop string) (*log.Logger, error) {
	logger := l.WithOptions(zap.AddCallerSkip(2 + 2))
	logFunc, err := levelToFunc(logger, level)
	if err != nil {
		return nil, err
	}
	return log.New(&loggerWriter{prefixToDrop: []byte(prefixToDrop), logFunc: logFunc}, "", 0), nil
}

// addr generates an address:port string to be used in defining an http.Server
func addr(listenAddress, port string) string {
	return fmt.Sprintf("%s:%s", listenAddress, port)
}

// stdLogError creates a *log.logger based off an existing zap.Logger instance.
// Some libraries call log.logger directly, which isn't structured as JSON. This method
// Will reformat log calls as zap.Error logs.
func stdLogError(logger *zap.Logger, prefixToDrop string) (*log.Logger, error) {
	standardLog, err := newStdLogAt(logger, zapcore.ErrorLevel, prefixToDrop)
	if err != nil {
		return nil, err
	}
	return standardLog, nil
}

// serverConfig generates a *http.Server with a structured error logger.
func (s Server) serverConfig(tlsConfig *tls.Config, prefixToDrop string) (*http.Server, error) {
	// By detault http.Server will use the standard logging library which isn't
	// structured JSON. This will pass zap.Logger with log level error
	standardLog, err := stdLogError(s.Logger, prefixToDrop)
	if err != nil {
		s.Logger.Error("failed to create an error logger", zap.Error(err))
		return nil, errors.Wrap(err, "Faile")
	}

	serverConfig := &http.Server{
		Addr:           addr(s.ListenAddress, s.Port),
		ErrorLog:       standardLog,
		Handler:        s.HTTPHandler,
		IdleTimeout:    idleTimeout,
		MaxHeaderBytes: maxHeaderSize,
		TLSConfig:      tlsConfig,
	}
	return serverConfig, err
}

// tlsConfig generates a new *tls.Config. It will
func (s Server) tlsConfig() (*tls.Config, error) {
	var caCerts *x509.CertPool
	var tlsCerts []tls.Certificate
	var err error

	// Load client Certificate Authority (CA) if we are requiring client
	// cert authentication.
	if s.ClientAuthType == tls.VerifyClientCertIfGiven ||
		s.ClientAuthType == tls.RequireAndVerifyClientCert {
		if s.CACertPEMBlock == nil {
			return nil, ErrMissingCACert

		}
		caCerts = x509.NewCertPool()
		ok := caCerts.AppendCertsFromPEM(s.CACertPEMBlock)
		if !ok {
			return nil, ErrUnparseableCACert
		}
	}

	// Parse and append all of the TLSCerts to the tls.Config
	for _, cert := range s.TLSCerts {
		parsedCert, err := tls.X509KeyPair(cert.CertPEMBlock, cert.KeyPEMBlock)
		if err != nil {
			s.Logger.Error("failed to parse tls certificate", zap.Error(err))
			return nil, err
		}
		tlsCerts = append(tlsCerts, parsedCert)
	}

	tlsConfig := &tls.Config{
		ClientCAs:    caCerts,
		Certificates: tlsCerts,
		NextProtos:   supportedProtocols,
		ClientAuth:   s.ClientAuthType,
	}

	// Map certificates with the CommonName / DNSNames to support
	// Subject Name Indication (SNI). In other words this will tell
	// the TLS listener to sever the appropriate certificate matching
	// the requested hostname.
	tlsConfig.BuildNameToCertificate()

	return tlsConfig, err
}

// ListenAndServeTLS returns a TLS Listener function for serving HTTPS requests
func (s Server) ListenAndServeTLS() error {
	var serverFunc serverFunc
	var server *http.Server
	var tlsConfig *tls.Config
	var err error

	tlsConfig, err = s.tlsConfig()
	if err != nil {
		s.Logger.Error("failed to generate a TLS config", zap.Error(err))
		return err
	}

	server, err = s.serverConfig(tlsConfig, s.PrefixToDrop)
	if err != nil {
		s.Logger.Error("failed to generate a TLS server config", zap.Error(err))
		return err
	}

	s.Logger.Info("start https listener",
		zap.Duration("idle-timeout", server.IdleTimeout),
		zap.Any("listen-address", s.ListenAddress),
		zap.Int("max-header-bytes", server.MaxHeaderBytes),
		zap.String("port", s.Port),
	)

	serverFunc = func(httpServer *http.Server) error {
		tlsListener, err := tls.Listen("tcp",
			server.Addr,
			tlsConfig)
		if err != nil {
			return err
		}
		defer tlsListener.Close()
		return server.Serve(tlsListener)
	}

	return serverFunc(server)
}

// ListenAndServe returns an HTTP ListenAndServe function for serving HTTP requests
func (s Server) ListenAndServe() error {
	var serverFunc serverFunc
	var server *http.Server
	var tlsConfig *tls.Config
	var err error

	server, err = s.serverConfig(tlsConfig, s.PrefixToDrop)
	if err != nil {
		s.Logger.Error("failed to generate a server config", zap.Error(err))
		return err
	}

	s.Logger.Info("start http listener",
		zap.Duration("idle-timeout", server.IdleTimeout),
		zap.Any("listen-address", s.ListenAddress),
		zap.Int("max-header-bytes", server.MaxHeaderBytes),
		zap.String("port", s.Port),
	)

	serverFunc = func(httpServer *http.Server) error {
		return httpServer.ListenAndServe()
	}

	return serverFunc(server)
}
