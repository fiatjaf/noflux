// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/json"
	"github.com/fiatjaf/noflux/internal/integration"
	"github.com/fiatjaf/noflux/internal/model"
)

func (h *handler) saveEntry(w http.ResponseWriter, r *http.Request) {
	entryID := request.RouteInt64Param(r, "entryID")
	builder := h.store.NewEntryQueryBuilder(request.UserID(r))
	builder.WithEntryID(entryID)
	builder.WithoutStatus(model.EntryStatusRemoved)

	entry, err := builder.GetEntry()
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	if entry == nil {
		json.NotFound(w, r)
		return
	}

	userIntegrations, err := h.store.Integration(request.UserID(r))
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	go integration.SendEntry(entry, userIntegrations)

	json.Created(w, r, map[string]string{"message": "saved"})
}
