// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package parser // import "github.com/fiatjaf/noflux/internal/reader/parser"

import (
	"errors"
	"io"

	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/reader/atom"
	"github.com/fiatjaf/noflux/internal/reader/json"
	"github.com/fiatjaf/noflux/internal/reader/rdf"
	"github.com/fiatjaf/noflux/internal/reader/rss"
)

var ErrFeedFormatNotDetected = errors.New("parser: unable to detect feed format")

// ParseFeed analyzes the input data and returns a normalized feed object.
func ParseFeed(baseURL string, r io.ReadSeeker) (*model.Feed, error) {
	r.Seek(0, io.SeekStart)
	format, version := DetectFeedFormat(r)
	switch format {
	case FormatAtom:
		r.Seek(0, io.SeekStart)
		return atom.Parse(baseURL, r, version)
	case FormatRSS:
		r.Seek(0, io.SeekStart)
		return rss.Parse(baseURL, r)
	case FormatJSON:
		r.Seek(0, io.SeekStart)
		return json.Parse(baseURL, r)
	case FormatRDF:
		r.Seek(0, io.SeekStart)
		return rdf.Parse(baseURL, r)
	default:
		return nil, ErrFeedFormatNotDetected
	}
}
