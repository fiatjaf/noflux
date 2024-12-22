// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package rss // import "github.com/fiatjaf/noflux/internal/reader/rss"

// FeedBurnerItemElement represents FeedBurner XML elements.
type FeedBurnerItemElement struct {
	FeedBurnerLink          string `xml:"http://rssnamespace.org/feedburner/ext/1.0 origLink"`
	FeedBurnerEnclosureLink string `xml:"http://rssnamespace.org/feedburner/ext/1.0 origEnclosureLink"`
}
