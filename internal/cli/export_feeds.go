// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package cli // import "github.com/fiatjaf/noflux/internal/cli"

import (
	"fmt"

	"github.com/fiatjaf/noflux/internal/reader/opml"
	"github.com/fiatjaf/noflux/internal/storage"
)

func exportUserFeeds(store *storage.Storage, username string) {
	user, err := store.UserByUsername(username)
	if err != nil {
		printErrorAndExit(fmt.Errorf("unable to find user: %w", err))
	}

	if user == nil {
		printErrorAndExit(fmt.Errorf("user %q not found", username))
	}

	opmlHandler := opml.NewHandler(store)
	opmlExport, err := opmlHandler.Export(user.ID)
	if err != nil {
		printErrorAndExit(fmt.Errorf("unable to export feeds: %w", err))
	}

	fmt.Println(opmlExport)
}
