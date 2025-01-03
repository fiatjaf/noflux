// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package api // import "github.com/fiatjaf/noflux/internal/api"

import (
	json_parser "encoding/json"
	"net/http"

	"github.com/fiatjaf/noflux/internal/config"
	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/json"
	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/reader/fetcher"
	"github.com/fiatjaf/noflux/internal/reader/subscription"
	"github.com/fiatjaf/noflux/internal/validator"
)

func (h *handler) discoverSubscriptions(w http.ResponseWriter, r *http.Request) {
	var subscriptionDiscoveryRequest model.SubscriptionDiscoveryRequest
	if err := json_parser.NewDecoder(r.Body).Decode(&subscriptionDiscoveryRequest); err != nil {
		json.BadRequest(w, r, err)
		return
	}

	if validationErr := validator.ValidateSubscriptionDiscovery(&subscriptionDiscoveryRequest); validationErr != nil {
		json.BadRequest(w, r, validationErr.Error())
		return
	}

	var rssbridgeURL string
	intg, err := h.store.Integration(request.UserID(r))
	if err == nil && intg != nil && intg.RSSBridgeEnabled {
		rssbridgeURL = intg.RSSBridgeURL
	}

	requestBuilder := fetcher.NewRequestBuilder()
	requestBuilder.WithTimeout(config.Opts.HTTPClientTimeout())
	requestBuilder.WithProxy(config.Opts.HTTPClientProxy())
	requestBuilder.WithUserAgent(subscriptionDiscoveryRequest.UserAgent, config.Opts.HTTPClientUserAgent())
	requestBuilder.WithCookie(subscriptionDiscoveryRequest.Cookie)
	requestBuilder.WithUsernameAndPassword(subscriptionDiscoveryRequest.Username, subscriptionDiscoveryRequest.Password)
	requestBuilder.UseProxy(subscriptionDiscoveryRequest.FetchViaProxy)
	requestBuilder.IgnoreTLSErrors(subscriptionDiscoveryRequest.AllowSelfSignedCertificates)
	requestBuilder.DisableHTTP2(subscriptionDiscoveryRequest.DisableHTTP2)

	subscriptions, localizedError := subscription.NewSubscriptionFinder(requestBuilder).FindSubscriptions(
		subscriptionDiscoveryRequest.URL,
		rssbridgeURL,
	)

	if localizedError != nil {
		json.ServerError(w, r, localizedError.Error())
		return
	}

	if len(subscriptions) == 0 {
		json.NotFound(w, r)
		return
	}

	json.OK(w, r, subscriptions)
}
