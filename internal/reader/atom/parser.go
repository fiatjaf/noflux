// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package atom // import "github.com/fiatjaf/noflux/internal/reader/atom"

import (
	"fmt"
	"io"

	"github.com/fiatjaf/noflux/internal/model"
	xml_decoder "github.com/fiatjaf/noflux/internal/reader/xml"
)

// Parse returns a normalized feed struct from a Atom feed.
func Parse(baseURL string, r io.ReadSeeker, version string) (*model.Feed, error) {
	switch version {
	case "0.3":
		atomFeed := new(Atom03Feed)
		if err := xml_decoder.NewXMLDecoder(r).Decode(atomFeed); err != nil {
			return nil, fmt.Errorf("atom: unable to parse Atom 0.3 feed: %w", err)
		}
		return NewAtom03Adapter(atomFeed).BuildFeed(baseURL), nil
	default:
		atomFeed := new(Atom10Feed)
		if err := xml_decoder.NewXMLDecoder(r).Decode(atomFeed); err != nil {
			return nil, fmt.Errorf("atom: unable to parse Atom 1.0 feed: %w", err)
		}
		return NewAtom10Adapter(atomFeed).BuildFeed(baseURL), nil
	}
}
