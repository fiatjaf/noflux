// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package oauth2 // import "github.com/fiatjaf/noflux/internal/oauth2"

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/fiatjaf/noflux/internal/model"
)

// Provider is an interface for OAuth2 providers.
type Provider interface {
	GetConfig() *oauth2.Config
	GetUserExtraKey() string
	GetProfile(ctx context.Context, code, codeVerifier string) (*Profile, error)
	PopulateUserCreationWithProfileID(user *model.UserCreationRequest, profile *Profile)
	PopulateUserWithProfileID(user *model.User, profile *Profile)
	UnsetUserProfileID(user *model.User)
}
