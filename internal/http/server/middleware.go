// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package httpd // import "github.com/fiatjaf/noflux/internal/http/server"

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/fiatjaf/noflux/internal/config"
	"github.com/fiatjaf/noflux/internal/http/request"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := request.FindClientIP(r)
		ctx := r.Context()
		ctx = context.WithValue(ctx, request.ClientIPContextKey, clientIP)

		if r.Header.Get("X-Forwarded-Proto") == "https" {
			config.Opts.HTTPS = true
		}

		t1 := time.Now()
		defer func() {
			slog.Debug("Incoming request",
				slog.String("client_ip", clientIP),
				slog.Group("request",
					slog.String("method", r.Method),
					slog.String("uri", r.RequestURI),
					slog.String("protocol", r.Proto),
					slog.Duration("execution_time", time.Since(t1)),
				),
			)
		}()

		if config.Opts.HTTPS && config.Opts.HasHSTS() {
			w.Header().Set("Strict-Transport-Security", "max-age=31536000")
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
