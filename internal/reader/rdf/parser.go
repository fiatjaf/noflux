// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package rdf // import "github.com/fiatjaf/noflux/internal/reader/rdf"

import (
	"fmt"
	"io"

	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/reader/xml"
)

// Parse returns a normalized feed struct from a RDF feed.
func Parse(baseURL string, data io.ReadSeeker) (*model.Feed, error) {
	xmlFeed := new(RDF)
	if err := xml.NewXMLDecoder(data).Decode(xmlFeed); err != nil {
		return nil, fmt.Errorf("rdf: unable to parse feed: %w", err)
	}

	return NewRDFAdapter(xmlFeed).BuildFeed(baseURL), nil
}
