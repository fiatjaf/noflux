// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/request"
	"github.com/fiatjaf/noflux/internal/http/response/html"
	"github.com/fiatjaf/noflux/internal/http/response/xml"
	"github.com/fiatjaf/noflux/internal/reader/opml"
)

func (h *handler) exportFeeds(w http.ResponseWriter, r *http.Request) {
	opmlExport, err := opml.NewHandler(h.store).Export(request.UserID(r))
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	xml.Attachment(w, r, "feeds.opml", opmlExport)
}
