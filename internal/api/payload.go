// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package api // import "github.com/fiatjaf/noflux/internal/api"

import (
	"github.com/fiatjaf/noflux/internal/model"
)

type feedIconResponse struct {
	ID       int64  `json:"id"`
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

type entriesResponse struct {
	Total   int           `json:"total"`
	Entries model.Entries `json:"entries"`
}

type feedCreationResponse struct {
	FeedID int64 `json:"feed_id"`
}

type versionResponse struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Compiler  string `json:"compiler"`
	Arch      string `json:"arch"`
	OS        string `json:"os"`
}
