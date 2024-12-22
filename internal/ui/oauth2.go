// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"context"

	"github.com/fiatjaf/noflux/internal/config"
	"github.com/fiatjaf/noflux/internal/oauth2"
)

func getOAuth2Manager(ctx context.Context) *oauth2.Manager {
	return oauth2.NewManager(
		ctx,
		config.Opts.OAuth2ClientID(),
		config.Opts.OAuth2ClientSecret(),
		config.Opts.OAuth2RedirectURL(),
		config.Opts.OIDCDiscoveryEndpoint(),
	)
}
