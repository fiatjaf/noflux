// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package cli // import "github.com/fiatjaf/noflux/internal/cli"

import (
	"fmt"

	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/storage"
	"github.com/fiatjaf/noflux/internal/validator"
)

func resetPassword(store *storage.Storage) {
	username, password := askCredentials()
	user, err := store.UserByUsername(username)
	if err != nil {
		printErrorAndExit(err)
	}

	if user == nil {
		printErrorAndExit(fmt.Errorf("user not found"))
	}

	userModificationRequest := &model.UserModificationRequest{
		Password: &password,
	}
	if validationErr := validator.ValidateUserModification(store, user.ID, userModificationRequest); validationErr != nil {
		printErrorAndExit(validationErr.Error())
	}

	user.Password = password
	if err := store.UpdateUser(user); err != nil {
		printErrorAndExit(err)
	}

	fmt.Println("Password changed!")
}
