// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "github.com/fiatjaf/noflux/internal/ui"

import (
	"github.com/fiatjaf/noflux/internal/storage"
	"github.com/fiatjaf/noflux/internal/template"
	"github.com/fiatjaf/noflux/internal/worker"

	"github.com/gorilla/mux"
)

type handler struct {
	router *mux.Router
	store  *storage.Storage
	tpl    *template.Engine
	pool   *worker.Pool
}
