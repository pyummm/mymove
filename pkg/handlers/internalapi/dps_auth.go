package internalapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/dps_auth"
	"github.com/transcom/mymove/pkg/handlers"
)

// DPSAuthCookieHandler handles the authentication process for DPS
type DPSAuthCookieHandler struct {
	handlers.HandlerContext
}

// Handle sets the cookie and begins the authentication process for DPS
func (h DPSAuthCookieHandler) Handle(params dps_auth.SetDPSAuthCookieParams) middleware.Responder {
	return dps_auth.NewSetDPSAuthCookieOK()
}
