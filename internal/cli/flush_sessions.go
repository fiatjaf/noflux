// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package cli // import "github.com/fiatjaf/noflux/internal/cli"

import (
	"fmt"

	"github.com/fiatjaf/noflux/internal/storage"
)

func flushSessions(store *storage.Storage) {
	fmt.Println("Flushing all sessions (disconnect users)")
	if err := store.FlushAllSessions(); err != nil {
		printErrorAndExit(err)
	}
}
