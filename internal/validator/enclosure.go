// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package validator

import (
	"fmt"

	"github.com/fiatjaf/noflux/internal/model"
)

func ValidateEnclosureUpdateRequest(request *model.EnclosureUpdateRequest) error {
	if request.MediaProgression < 0 {
		return fmt.Errorf(`media progression must an positive integer`)
	}

	return nil
}
