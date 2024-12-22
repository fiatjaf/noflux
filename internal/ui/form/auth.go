// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package form // import "github.com/fiatjaf/noflux/internal/ui/form"

import (
	"net/http"
	"strings"

	"github.com/fiatjaf/noflux/internal/locale"
)

// AuthForm represents the authentication form.
type AuthForm struct {
	Username string
	Password string
}

// Validate makes sure the form values are valid.
func (a AuthForm) Validate() *locale.LocalizedError {
	if a.Username == "" || a.Password == "" {
		return locale.NewLocalizedError("error.fields_mandatory")
	}

	return nil
}

// NewAuthForm returns a new AuthForm.
func NewAuthForm(r *http.Request) *AuthForm {
	return &AuthForm{
		Username: strings.TrimSpace(r.FormValue("username")),
		Password: strings.TrimSpace(r.FormValue("password")),
	}
}
