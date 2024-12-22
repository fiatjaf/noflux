// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package model // import "github.com/fiatjaf/noflux/internal/model"

import (
	"fmt"
	"time"

	"github.com/fiatjaf/noflux/internal/timezone"
)

// UserSession represents a user session in the system.
type UserSession struct {
	ID        int64
	UserID    int64
	Token     string
	CreatedAt time.Time
	UserAgent string
	IP        string
}

func (u *UserSession) String() string {
	return fmt.Sprintf(`ID=%q, UserID=%q, IP=%q, Token=%q`, u.ID, u.UserID, u.IP, u.Token)
}

// UseTimezone converts creation date to the given timezone.
func (u *UserSession) UseTimezone(tz string) {
	u.CreatedAt = timezone.Convert(tz, u.CreatedAt)
}

// UserSessions represents a list of sessions.
type UserSessions []*UserSession

// UseTimezone converts creation date of all sessions to the given timezone.
func (u UserSessions) UseTimezone(tz string) {
	for _, session := range u {
		session.UseTimezone(tz)
	}
}
