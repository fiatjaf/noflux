// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/html"
	"github.com/fiatjaf/noflux/internal/http/route"
)

func (h *handler) removeAPIKey(w http.ResponseWriter, r *http.Request) {
	keyID := request.RouteInt64Param(r, "keyID")
	err := h.store.RemoveAPIKey(request.UserID(r), keyID)
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	html.Redirect(w, r, route.Path(h.router, "apiKeys"))
}
