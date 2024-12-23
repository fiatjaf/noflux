// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/html"
	"github.com/fiatjaf/noflux/internal/http/route"
	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/ui/form"
	"github.com/fiatjaf/noflux/internal/ui/session"
	"github.com/fiatjaf/noflux/internal/ui/view"
	"github.com/fiatjaf/noflux/internal/validator"
)

func (h *handler) saveCategory(w http.ResponseWriter, r *http.Request) {
	loggedUser, err := h.store.UserByID(request.UserID(r))
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	categoryForm := form.NewCategoryForm(r)

	sess := session.New(h.store, request.SessionID(r))
	view := view.New(h.tpl, r, sess)
	view.Set("form", categoryForm)
	view.Set("menu", "categories")
	view.Set("user", loggedUser)
	view.Set("countUnread", h.store.CountUnreadEntries(loggedUser.ID))
	view.Set("countErrorFeeds", h.store.CountUserFeedsWithErrors(loggedUser.ID))

	categoryRequest := &model.CategoryRequest{Title: categoryForm.Title}

	if validationErr := validator.ValidateCategoryCreation(h.store, loggedUser.ID, categoryRequest); validationErr != nil {
		view.Set("errorMessage", validationErr.Translate(loggedUser.Language))
		html.OK(w, r, view.Render("create_category"))
		return
	}

	if _, err = h.store.CreateCategory(loggedUser.ID, categoryRequest); err != nil {
		html.ServerError(w, r, err)
		return
	}

	html.Redirect(w, r, route.Path(h.router, "categories"))
}
