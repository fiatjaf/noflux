// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package cli // import "github.com/fiatjaf/noflux/internal/cli"

import (
	"log/slog"
	"time"

	"github.com/fiatjaf/noflux/internal/config"
	"github.com/fiatjaf/noflux/internal/storage"
	"github.com/fiatjaf/noflux/internal/worker"
)

func runScheduler(store *storage.Storage, pool *worker.Pool) {
	slog.Debug(`Starting background scheduler...`)

	go feedScheduler(
		store,
		pool,
		config.Opts.PollingFrequency(),
		config.Opts.BatchSize(),
		config.Opts.PollingParsingErrorLimit(),
	)

	go cleanupScheduler(
		store,
		config.Opts.CleanupFrequencyHours(),
	)
}

func feedScheduler(store *storage.Storage, pool *worker.Pool, frequency, batchSize, errorLimit int) {
	for range time.Tick(time.Duration(frequency) * time.Minute) {
		// Generate a batch of feeds for any user that has feeds to refresh.
		batchBuilder := store.NewBatchBuilder()
		batchBuilder.WithBatchSize(batchSize)
		batchBuilder.WithErrorLimit(errorLimit)
		batchBuilder.WithoutDisabledFeeds()
		batchBuilder.WithNextCheckExpired()

		if jobs, err := batchBuilder.FetchJobs(); err != nil {
			slog.Error("Unable to fetch jobs from database", slog.Any("error", err))
		} else if len(jobs) > 0 {
			slog.Info("Created a batch of feeds",
				slog.Int("nb_jobs", len(jobs)),
			)
			pool.Push(jobs)
		}
	}
}

func cleanupScheduler(store *storage.Storage, frequency int) {
	for range time.Tick(time.Duration(frequency) * time.Hour) {
		runCleanupTasks(store)
	}
}
