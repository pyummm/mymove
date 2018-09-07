package internalapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/dps_auth"
	"github.com/transcom/mymove/pkg/handlers"
)

const cookieExpiresInHours = 1

// SetDPSAuthCookieOKResponder is a custom responder that sets the DPS authentication cookie
// when writing the response
type SetDPSAuthCookieOKResponder struct {
	cookie http.Cookie
}

// NewSetDPSAuthCookieOKResponder creates a new SetDPSAuthCookieOKResponder
func NewSetDPSAuthCookieOKResponder(cookie http.Cookie) *SetDPSAuthCookieOKResponder {
	return &SetDPSAuthCookieOKResponder{cookie: cookie}
}

// WriteResponse to the client
func (o *SetDPSAuthCookieOKResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	http.SetCookie(rw, &o.cookie)

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses
	rw.WriteHeader(200)
}

// DPSAuthCookieHandler handles the authentication process for DPS
type DPSAuthCookieHandler struct {
	handlers.HandlerContext
}

// Handle sets the cookie necessary for beginning the authentication process for DPS
func (h DPSAuthCookieHandler) Handle(params dps_auth.SetDPSAuthCookieParams) middleware.Responder {
	cookieName := "DPS"
	if params.CookieName != nil {
		cookieName = *params.CookieName
	}

	cookieValue, err := h.encryptedCookieValue(params)
	if err != nil {
		fmt.Println(err)
		return dps_auth.NewSetDPSAuthCookieInternalServerError()
	}

	cookie := http.Cookie{Name: cookieName, Value: "mymove:" + cookieValue}
	return NewSetDPSAuthCookieOKResponder(cookie)
}

func (h DPSAuthCookieHandler) encryptedCookieValue(params dps_auth.SetDPSAuthCookieParams) (string, error) {
	session := auth.SessionFromRequestContext(params.HTTPRequest)

	expirationTime := time.Now().Add(time.Hour * time.Duration(cookieExpiresInHours)).Unix()
	fmt.Println(session.UserID.String())
	value := map[string]string{
		"user_id":    session.UserID.String(),
		"expires_at": strconv.FormatInt(expirationTime, 10),
	}

	valueJSON, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return encrypt(valueJSON)
}

func encrypt(data []byte) (string, error) {
	// TODO: encrypt
	return base64.StdEncoding.EncodeToString(data), nil

	/*
		key := os.Getenv("DPS_AUTH_COOKIE_SECRET_KEY")
		c, err := aes.NewCipher([]byte(key))
		if err != nil {
			return "", err
		}

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			return "", err
		}

		nonce := make([]byte, gcm.NonceSize())
		fmt.Print("NNN: ")
		fmt.Println(nonce)
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			return "", err
		}

		encoded := gcm.Seal(nonce, nonce, data, nil)
		return base64.StdEncoding.EncodeToString(encoded), nil
	*/
}
