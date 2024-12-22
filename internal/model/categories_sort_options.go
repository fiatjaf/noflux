// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package model // import "github.com/fiatjaf/noflux/internal/model"

func CategoriesSortingOptions() map[string]string {
	return map[string]string{
		"unread_count": "form.prefs.select.unread_count",
		"alphabetical": "form.prefs.select.alphabetical",
	}
}
