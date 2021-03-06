package models_test

import (
	"fmt"

	"github.com/gobuffalo/uuid"

	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ModelSuite) Test_BackupContactCreate() {
	t := suite.T()

	serviceMember := testdatagen.MakeDefaultServiceMember(suite.db)

	newContact := models.BackupContact{
		ServiceMemberID: serviceMember.ID,
		ServiceMember:   serviceMember,
		Name:            "name",
		Email:           "email@example.com",
		Permission:      internalmessages.BackupContactPermissionEDIT,
	}

	verrs, err := suite.db.ValidateAndCreate(&newContact)

	if err != nil {
		fmt.Println(err)
		t.Fatal("could not save BackupContact", err)
	}

	if verrs.Count() != 0 {
		t.Errorf("did not expect validation errors: %v", verrs)
	}
}

func (suite *ModelSuite) Test_BackupContactValidations() {
	contact := &models.BackupContact{}

	var expErrors = map[string][]string{
		"name":       {"Name can not be blank."},
		"email":      {"Email can not be blank."},
		"permission": {"Permission can not be blank."},
	}

	suite.verifyValidationErrors(contact, expErrors)
}

func (suite *ModelSuite) Test_FetchBackupContact() {
	t := suite.T()

	serviceMember1 := testdatagen.MakeDefaultServiceMember(suite.db)
	serviceMember2 := testdatagen.MakeDefaultServiceMember(suite.db)

	backupContact := models.BackupContact{
		ServiceMemberID: serviceMember1.ID,
		Name:            "name",
		Email:           "email@example.com",
		Permission:      internalmessages.BackupContactPermissionEDIT,
	}
	suite.mustSave(&backupContact)

	session := &auth.Session{
		UserID:          serviceMember1.UserID,
		ServiceMemberID: serviceMember1.ID,
		ApplicationName: auth.MyApp,
	}
	shouldSucceed, err := models.FetchBackupContact(suite.db, session, backupContact.ID)
	if err != nil || !uuid.Equal(backupContact.ID, shouldSucceed.ID) {
		t.Errorf("failed retrieving own backup contact: %v", err)
	}

	session.UserID = serviceMember2.UserID
	session.ServiceMemberID = serviceMember2.ID
	_, err = models.FetchBackupContact(suite.db, session, backupContact.ID)
	if err == nil {
		t.Error("should have failed getting other user's contact")
	}
}
