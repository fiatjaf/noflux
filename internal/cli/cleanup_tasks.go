// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package cli // import "github.com/fiatjaf/noflux/internal/cli"

import (
	"log/slog"
	"time"

	"github.com/fiatjaf/noflux/internal/config"
	"github.com/fiatjaf/noflux/internal/metric"
	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/storage"
)

func runCleanupTasks(store *storage.Storage) {
	nbSessions := store.CleanOldSessions(config.Opts.CleanupRemoveSessionsDays())
	nbUserSessions := store.CleanOldUserSessions(config.Opts.CleanupRemoveSessionsDays())
	slog.Info("Sessions cleanup completed",
		slog.Int64("application_sessions_removed", nbSessions),
		slog.Int64("user_sessions_removed", nbUserSessions),
	)

	startTime := time.Now()
	if rowsAffected, err := store.ArchiveEntries(model.EntryStatusRead, config.Opts.CleanupArchiveReadDays(), config.Opts.CleanupArchiveBatchSize()); err != nil {
		slog.Error("Unable to archive read entries", slog.Any("error", err))
	} else {
		slog.Info("Archiving read entries completed",
			slog.Int64("read_entries_archived", rowsAffected),
		)

		if config.Opts.HasMetricsCollector() {
			metric.ArchiveEntriesDuration.WithLabelValues(model.EntryStatusRead).Observe(time.Since(startTime).Seconds())
		}
	}

	startTime = time.Now()
	if rowsAffected, err := store.ArchiveEntries(model.EntryStatusUnread, config.Opts.CleanupArchiveUnreadDays(), config.Opts.CleanupArchiveBatchSize()); err != nil {
		slog.Error("Unable to archive unread entries", slog.Any("error", err))
	} else {
		slog.Info("Archiving unread entries completed",
			slog.Int64("unread_entries_archived", rowsAffected),
		)

		if config.Opts.HasMetricsCollector() {
			metric.ArchiveEntriesDuration.WithLabelValues(model.EntryStatusUnread).Observe(time.Since(startTime).Seconds())
		}
	}
}
