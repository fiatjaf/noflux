// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package form // import "github.com/fiatjaf/noflux/internal/ui/form"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/locale"
	"github.com/fiatjaf/noflux/internal/model"
)

// UserForm represents the user form.
type UserForm struct {
	Username     string
	Password     string
	Confirmation string
	IsAdmin      bool
}

// ValidateCreation validates user creation.
func (u UserForm) ValidateCreation() *locale.LocalizedError {
	if u.Username == "" || u.Password == "" || u.Confirmation == "" {
		return locale.NewLocalizedError("error.fields_mandatory")
	}

	if u.Password != u.Confirmation {
		return locale.NewLocalizedError("error.different_passwords")
	}

	return nil
}

// ValidateModification validates user modification.
func (u UserForm) ValidateModification() *locale.LocalizedError {
	if u.Username == "" {
		return locale.NewLocalizedError("error.user_mandatory_fields")
	}

	if u.Password != "" {
		if u.Password != u.Confirmation {
			return locale.NewLocalizedError("error.different_passwords")
		}

		if len(u.Password) < 6 {
			return locale.NewLocalizedError("error.password_min_length")
		}
	}

	return nil
}

// Merge updates the fields of the given user.
func (u UserForm) Merge(user *model.User) *model.User {
	user.Username = u.Username
	user.IsAdmin = u.IsAdmin

	if u.Password != "" {
		user.Password = u.Password
	}

	return user
}

// NewUserForm returns a new UserForm.
func NewUserForm(r *http.Request) *UserForm {
	return &UserForm{
		Username:     r.FormValue("username"),
		Password:     r.FormValue("password"),
		Confirmation: r.FormValue("confirmation"),
		IsAdmin:      r.FormValue("is_admin") == "1",
	}
}
