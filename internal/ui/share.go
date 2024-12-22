// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"
	"time"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response"
	"github.com/fiatjaf/noflux/internal/http/response/html"
	"github.com/fiatjaf/noflux/internal/http/route"
	"github.com/fiatjaf/noflux/internal/storage"
	"github.com/fiatjaf/noflux/internal/ui/session"
	"github.com/fiatjaf/noflux/internal/ui/view"
)

func (h *handler) createSharedEntry(w http.ResponseWriter, r *http.Request) {
	entryID := request.RouteInt64Param(r, "entryID")
	shareCode, err := h.store.EntryShareCode(request.UserID(r), entryID)
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	html.Redirect(w, r, route.Path(h.router, "sharedEntry", "shareCode", shareCode))
}

func (h *handler) unshareEntry(w http.ResponseWriter, r *http.Request) {
	entryID := request.RouteInt64Param(r, "entryID")
	if err := h.store.UnshareEntry(request.UserID(r), entryID); err != nil {
		html.ServerError(w, r, err)
		return
	}

	html.Redirect(w, r, route.Path(h.router, "sharedEntries"))
}

func (h *handler) sharedEntry(w http.ResponseWriter, r *http.Request) {
	shareCode := request.RouteStringParam(r, "shareCode")
	if shareCode == "" {
		html.NotFound(w, r)
		return
	}

	etag := shareCode
	response.New(w, r).WithCaching(etag, 72*time.Hour, func(b *response.Builder) {
		builder := storage.NewAnonymousQueryBuilder(h.store)
		builder.WithShareCode(shareCode)

		entry, err := builder.GetEntry()
		if err != nil || entry == nil {
			html.NotFound(w, r)
			return
		}

		sess := session.New(h.store, request.SessionID(r))
		view := view.New(h.tpl, r, sess)
		view.Set("entry", entry)

		b.WithHeader("Content-Type", "text/html; charset=utf-8")
		b.WithBody(view.Render("entry"))
		b.Write()
	})
}
