// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/html"
	"github.com/fiatjaf/noflux/internal/http/route"
	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/ui/session"
	"github.com/fiatjaf/noflux/internal/ui/view"
)

func (h *handler) showSearchPage(w http.ResponseWriter, r *http.Request) {
	user, err := h.store.UserByID(request.UserID(r))
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	searchQuery := request.QueryStringParam(r, "q", "")
	offset := request.QueryIntParam(r, "offset", 0)

	var entries model.Entries
	var entriesCount int

	if searchQuery != "" {
		builder := h.store.NewEntryQueryBuilder(user.ID)
		builder.WithSearchQuery(searchQuery)
		builder.WithoutStatus(model.EntryStatusRemoved)
		builder.WithOffset(offset)
		builder.WithLimit(user.EntriesPerPage)

		entries, err = builder.GetEntries()
		if err != nil {
			html.ServerError(w, r, err)
			return
		}

		entriesCount, err = builder.CountEntries()
		if err != nil {
			html.ServerError(w, r, err)
			return
		}
	}

	sess := session.New(h.store, request.SessionID(r))
	view := view.New(h.tpl, r, sess)
	pagination := getPagination(route.Path(h.router, "search"), entriesCount, offset, user.EntriesPerPage)
	pagination.SearchQuery = searchQuery

	view.Set("searchQuery", searchQuery)
	view.Set("entries", entries)
	view.Set("total", entriesCount)
	view.Set("pagination", pagination)
	view.Set("menu", "search")
	view.Set("user", user)
	view.Set("countUnread", h.store.CountUnreadEntries(user.ID))
	view.Set("countErrorFeeds", h.store.CountUserFeedsWithErrors(user.ID))
	view.Set("hasSaveEntry", h.store.HasSaveEntry(user.ID))

	html.OK(w, r, view.Render("search"))
}
