// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package api // import "github.com/fiatjaf/noflux/internal/api"

import (
	json_parser "encoding/json"
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/json"
	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/validator"
)

func (h *handler) getEnclosureByID(w http.ResponseWriter, r *http.Request) {
	enclosureID := request.RouteInt64Param(r, "enclosureID")

	enclosure, err := h.store.GetEnclosure(enclosureID)
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	if enclosure == nil {
		json.NotFound(w, r)
		return
	}

	userID := request.UserID(r)
	if enclosure.UserID != userID {
		json.NotFound(w, r)
		return
	}

	enclosure.ProxifyEnclosureURL(h.router)

	json.OK(w, r, enclosure)
}

func (h *handler) updateEnclosureByID(w http.ResponseWriter, r *http.Request) {
	enclosureID := request.RouteInt64Param(r, "enclosureID")

	var enclosureUpdateRequest model.EnclosureUpdateRequest
	if err := json_parser.NewDecoder(r.Body).Decode(&enclosureUpdateRequest); err != nil {
		json.BadRequest(w, r, err)
		return
	}

	if err := validator.ValidateEnclosureUpdateRequest(&enclosureUpdateRequest); err != nil {
		json.BadRequest(w, r, err)
		return
	}

	enclosure, err := h.store.GetEnclosure(enclosureID)
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	if enclosure == nil {
		json.NotFound(w, r)
		return
	}

	userID := request.UserID(r)
	if enclosure.UserID != userID {
		json.NotFound(w, r)
		return
	}

	enclosure.MediaProgression = enclosureUpdateRequest.MediaProgression
	if err := h.store.UpdateEnclosure(enclosure); err != nil {
		json.ServerError(w, r, err)
		return
	}

	json.NoContent(w, r)
}
