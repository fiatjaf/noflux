// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package rss // import "github.com/fiatjaf/noflux/internal/reader/rss"

import (
	"fmt"
	"io"

	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/reader/xml"
)

// Parse returns a normalized feed struct from a RSS feed.
func Parse(baseURL string, data io.ReadSeeker) (*model.Feed, error) {
	rssFeed := new(RSS)
	decoder := xml.NewXMLDecoder(data)
	decoder.DefaultSpace = "rss"
	if err := decoder.Decode(rssFeed); err != nil {
		return nil, fmt.Errorf("rss: unable to parse feed: %w", err)
	}
	return NewRSSAdapter(rssFeed).BuildFeed(baseURL), nil
}
