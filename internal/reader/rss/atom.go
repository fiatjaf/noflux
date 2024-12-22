// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package rss // import "github.com/fiatjaf/noflux/internal/reader/rss"

import (
	"github.com/fiatjaf/noflux/internal/reader/atom"
)

type AtomAuthor struct {
	Author atom.AtomPerson `xml:"http://www.w3.org/2005/Atom author"`
}

func (a *AtomAuthor) PersonName() string {
	return a.Author.PersonName()
}

type AtomLinks struct {
	Links []*atom.AtomLink `xml:"http://www.w3.org/2005/Atom link"`
}
