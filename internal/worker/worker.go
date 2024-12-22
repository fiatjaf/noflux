// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package worker // import "github.com/fiatjaf/noflux/internal/worker"

import (
	"log/slog"
	"time"

	"github.com/fiatjaf/noflux/internal/config"
	"github.com/fiatjaf/noflux/internal/metric"
	"github.com/fiatjaf/noflux/internal/model"
	feedHandler "github.com/fiatjaf/noflux/internal/reader/handler"
	"github.com/fiatjaf/noflux/internal/storage"
)

// Worker refreshes a feed in the background.
type Worker struct {
	id    int
	store *storage.Storage
}

// Run wait for a job and refresh the given feed.
func (w *Worker) Run(c <-chan model.Job) {
	slog.Debug("Worker started",
		slog.Int("worker_id", w.id),
	)

	for {
		job := <-c
		slog.Debug("Job received by worker",
			slog.Int("worker_id", w.id),
			slog.Int64("user_id", job.UserID),
			slog.Int64("feed_id", job.FeedID),
		)

		startTime := time.Now()
		localizedError := feedHandler.RefreshFeed(w.store, job.UserID, job.FeedID, false)

		if config.Opts.HasMetricsCollector() {
			status := "success"
			if localizedError != nil {
				status = "error"
			}
			metric.BackgroundFeedRefreshDuration.WithLabelValues(status).Observe(time.Since(startTime).Seconds())
		}

		if localizedError != nil {
			slog.Warn("Unable to refresh a feed",
				slog.Int64("user_id", job.UserID),
				slog.Int64("feed_id", job.FeedID),
				slog.Any("error", localizedError.Error()),
			)
		}
	}
}
