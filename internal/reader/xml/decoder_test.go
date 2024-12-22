// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package xml // import "github.com/fiatjaf/noflux/internal/reader/xml"

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

func TestXMLDocumentWithIllegalUnicodeCharacters(t *testing.T) {
	type myxml struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Title   string   `xml:"title"`
	}

	expected := "Title & 中文标题"
	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><rss version="2.0"><title>Title & 中文%s标题</title></rss>`, "\x10")
	reader := strings.NewReader(data)

	var x myxml

	decoder := NewXMLDecoder(reader)
	err := decoder.Decode(&x)
	if err != nil {
		t.Error(err)
		return
	}
	if x.Title != expected {
		t.Errorf("Incorrect entry title, expected: %s, got: %s", expected, x.Title)
	}
}

func TestXMLDocumentWindows251EncodedWithIllegalCharacters(t *testing.T) {
	type myxml struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Title   string   `xml:"title"`
	}

	expected := "Title & 中文标题"
	data := fmt.Sprintf(`<?xml version="1.0" encoding="windows-1251"?><rss version="2.0"><title>Title & 中文%s标题</title></rss>`, "\x10")
	reader := strings.NewReader(data)

	var x myxml

	decoder := NewXMLDecoder(reader)
	err := decoder.Decode(&x)
	if err != nil {
		t.Error(err)
		return
	}
	if x.Title != expected {
		t.Errorf("Incorrect entry title, expected: %s, got: %s", expected, x.Title)
	}
}

func TestXMLDocumentWithIncorrectEncodingField(t *testing.T) {
	type myxml struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Title   string   `xml:"title"`
	}

	expected := "Title & 中文标题"
	data := fmt.Sprintf(`<?xml version="1.0" encoding="invalid"?><rss version="2.0"><title>Title & 中文%s标题</title></rss>`, "\x10")
	reader := strings.NewReader(data)

	var x myxml

	decoder := NewXMLDecoder(reader)
	err := decoder.Decode(&x)
	if err != nil {
		t.Error(err)
		return
	}
	if x.Title != expected {
		t.Errorf("Incorrect entry title, expected: %s, got: %s", expected, x.Title)
	}
}
