// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package cli // import "github.com/fiatjaf/noflux/internal/cli"

import (
	"fmt"
	"runtime"

	"github.com/fiatjaf/noflux/internal/version"
)

func info() {
	fmt.Println("Version:", version.Version)
	fmt.Println("Commit:", version.Commit)
	fmt.Println("Build Date:", version.BuildDate)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Compiler:", runtime.Compiler)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("OS:", runtime.GOOS)
}
