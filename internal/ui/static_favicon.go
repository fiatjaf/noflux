// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"net/http"
	"time"

	"github.com/fiatjaf/noflux/internal/http/response"
	"github.com/fiatjaf/noflux/internal/http/response/html"
	"github.com/fiatjaf/noflux/internal/ui/static"
)

func (h *handler) showFavicon(w http.ResponseWriter, r *http.Request) {
	etag, err := static.GetBinaryFileChecksum("favicon.ico")
	if err != nil {
		html.NotFound(w, r)
		return
	}

	response.New(w, r).WithCaching(etag, 48*time.Hour, func(b *response.Builder) {
		blob, err := static.LoadBinaryFile("favicon.ico")
		if err != nil {
			html.ServerError(w, r, err)
			return
		}

		b.WithHeader("Content-Type", "image/x-icon")
		b.WithoutCompression()
		b.WithBody(blob)
		b.Write()
	})
}
