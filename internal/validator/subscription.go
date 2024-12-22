// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package validator // import "github.com/fiatjaf/noflux/internal/validator"

import (
	"github.com/fiatjaf/noflux/internal/locale"
	"github.com/fiatjaf/noflux/internal/model"
)

// ValidateSubscriptionDiscovery validates subscription discovery requests.
func ValidateSubscriptionDiscovery(request *model.SubscriptionDiscoveryRequest) *locale.LocalizedError {
	if !IsValidURL(request.URL) {
		return locale.NewLocalizedError("error.invalid_site_url")
	}

	return nil
}
