// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

/*
Package client implements a client library for the Noflux REST API.

# Examples

This code snippet fetch the list of users:

	import (
		noflux "github.com/fiatjaf/noflux/client"
	)

	client := noflux.NewClient("https://api.example.org", "admin", "secret")
	users, err := client.Users()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users, err)

This one discover subscriptions on a website:

	subscriptions, err := client.Discover("https://example.org/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(subscriptions)
*/
package client // import "github.com/fiatjaf/noflux/client"
