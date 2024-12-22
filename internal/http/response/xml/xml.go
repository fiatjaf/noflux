// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package xml // import "github.com/fiatjaf/noflux/internal/http/response/xml"

import (
	"net/http"

	"github.com/fiatjaf/noflux/internal/http/response"
)

// OK writes a standard XML response with a status 200 OK.
func OK(w http.ResponseWriter, r *http.Request, body interface{}) {
	builder := response.New(w, r)
	builder.WithHeader("Content-Type", "text/xml; charset=utf-8")
	builder.WithBody(body)
	builder.Write()
}

// Attachment forces the XML document to be downloaded by the web browser.
func Attachment(w http.ResponseWriter, r *http.Request, filename string, body interface{}) {
	builder := response.New(w, r)
	builder.WithHeader("Content-Type", "text/xml; charset=utf-8")
	builder.WithAttachment(filename)
	builder.WithBody(body)
	builder.Write()
}
