// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/html"
	"github.com/fiatjaf/noflux/internal/http/route"
	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/storage"
	"github.com/fiatjaf/noflux/internal/ui/session"
	"github.com/fiatjaf/noflux/internal/ui/view"
)

func (h *handler) showReadEntryPage(w http.ResponseWriter, r *http.Request) {
	user, err := h.store.UserByID(request.UserID(r))
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	entryID := request.RouteInt64Param(r, "entryID")
	builder := h.store.NewEntryQueryBuilder(user.ID)
	builder.WithEntryID(entryID)
	builder.WithoutStatus(model.EntryStatusRemoved)

	entry, err := builder.GetEntry()
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	if entry == nil {
		html.NotFound(w, r)
		return
	}

	entryPaginationBuilder := storage.NewEntryPaginationBuilder(h.store, user.ID, entry.ID, "changed_at", "desc")
	entryPaginationBuilder.WithStatus(model.EntryStatusRead)
	prevEntry, nextEntry, err := entryPaginationBuilder.Entries()
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	nextEntryRoute := ""
	if nextEntry != nil {
		nextEntryRoute = route.Path(h.router, "readEntry", "entryID", nextEntry.ID)
	}

	prevEntryRoute := ""
	if prevEntry != nil {
		prevEntryRoute = route.Path(h.router, "readEntry", "entryID", prevEntry.ID)
	}

	sess := session.New(h.store, request.SessionID(r))
	view := view.New(h.tpl, r, sess)
	view.Set("entry", entry)
	view.Set("prevEntry", prevEntry)
	view.Set("nextEntry", nextEntry)
	view.Set("nextEntryRoute", nextEntryRoute)
	view.Set("prevEntryRoute", prevEntryRoute)
	view.Set("menu", "history")
	view.Set("user", user)
	view.Set("countUnread", h.store.CountUnreadEntries(user.ID))
	view.Set("countErrorFeeds", h.store.CountUserFeedsWithErrors(user.ID))
	view.Set("hasSaveEntry", h.store.HasSaveEntry(user.ID))

	html.OK(w, r, view.Render("entry"))
}
