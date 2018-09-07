package dpsapi

import (
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/transcom/mymove/pkg/gen/dpsapi"
	dpsops "github.com/transcom/mymove/pkg/gen/dpsapi/dpsoperations"
	"github.com/transcom/mymove/pkg/handlers"
)

// NewDPSAPIHandler returns a handler for the DPS API
func NewDPSAPIHandler(context handlers.HandlerContext) http.Handler {

	// Wire up the handlers to the ordersAPIMux
	dpsSpec, err := loads.Analyzed(dpsapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	dpsAPI := dpsops.NewMymoveAPI(dpsSpec)
	dpsAPI.DpsGetUserHandler = GetUserHandler{context}
	return dpsAPI.Serve(nil)
}
