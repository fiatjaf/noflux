// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package json // import "github.com/fiatjaf/noflux/internal/reader/json"

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/fiatjaf/noflux/internal/model"
)

// Parse returns a normalized feed struct from a JSON feed.
func Parse(baseURL string, data io.Reader) (*model.Feed, error) {
	jsonFeed := new(JSONFeed)
	if err := json.NewDecoder(data).Decode(&jsonFeed); err != nil {
		return nil, fmt.Errorf("json: unable to parse feed: %w", err)
	}

	return NewJSONAdapter(jsonFeed).BuildFeed(baseURL), nil
}
