package dpsapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/transcom/mymove/pkg/gen/dpsapi/dpsoperations/dps"
	"github.com/transcom/mymove/pkg/gen/dpsmessages"
	"github.com/transcom/mymove/pkg/handlers"
)

// GetUserHandler returns Orders by uuid
type GetUserHandler struct {
	handlers.HandlerContext
}

// Handle simply returns a NotImplementedError
func (h GetUserHandler) Handle(params dps.GetUserParams) middleware.Responder {
	token := params.Token
	fmt.Println(token)
	prefix := "mymove:"
	if !strings.HasPrefix(token, prefix) {
		return dps.NewGetUserInternalServerError()
	}

	decryptedToken, err := decrypt(token[len(prefix):])
	if err != nil {
		return dps.NewGetUserInternalServerError()
	}

	var values map[string]string
	err = json.Unmarshal(decryptedToken, &values)
	if err != nil {
		return dps.NewGetUserInternalServerError()
	}

	userID := values["user_id"]
	fmt.Println(userID)

	return dps.NewGetUserOK().WithPayload(getPayload(userID))
}

func getPayload(userID string) *dpsmessages.AuthenticationUserPayload {
	// TODO: Add real data
	affiliation := dpsmessages.AffiliationArmy
	middleName := "M"
	suffix := "III"
	telephone := "(555) 555-5555"
	payload := dpsmessages.AuthenticationUserPayload{
		Affiliation:          &affiliation,
		Email:                "test@example.com",
		FirstName:            "Jane",
		MiddleName:           &middleName,
		LastName:             "Doe",
		Suffix:               &suffix,
		LoginGovID:           strfmt.UUID(userID),
		SocialSecurityNumber: "555-55-5555",
		Telephone:            &telephone,
	}
	return &payload
}

func decrypt(data string) ([]byte, error) {
	// TODO: decrypt
	return base64.StdEncoding.DecodeString(data)

	/*
		    key := os.Getenv("DPS_AUTH_COOKIE_SECRET_KEY")
				var plaintext []byte
				c, err := aes.NewCipher([]byte(key))
				if err != nil {
					return plaintext, err
				}

				gcm, err := cipher.NewGCM(c)
				if err != nil {
					return plaintext, err
				}

				nonceSize := gcm.NonceSize()
				nonce, ciphertext := data[:nonceSize], data[nonceSize:]
				cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
				if err != nil {
					fmt.Println(err)
					return plaintext, err
				}

				plaintext, err = gcm.Open(nil, []byte(nonce), cipherBytes, nil)
				if err != nil {
					fmt.Println(err)
					return plaintext, err
				}

				return plaintext, nil
	*/
}
