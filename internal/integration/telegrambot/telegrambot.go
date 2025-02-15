// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package telegrambot // import "github.com/fiatjaf/noflux/internal/integration/telegrambot"

import (
	"fmt"

	"github.com/fiatjaf/noflux/internal/model"
)

func PushEntry(feed *model.Feed, entry *model.Entry, botToken, chatID string, topicID *int64, disableWebPagePreview, disableNotification bool, disableButtons bool) error {
	formattedText := fmt.Sprintf(
		`<b>%s</b> - <a href=%q>%s</a>`,
		feed.Title,
		entry.URL,
		entry.Title,
	)

	message := &MessageRequest{
		ChatID:                chatID,
		Text:                  formattedText,
		ParseMode:             HTMLFormatting,
		DisableWebPagePreview: disableWebPagePreview,
		DisableNotification:   disableNotification,
	}

	if topicID != nil {
		message.MessageThreadID = *topicID
	}

	if !disableButtons {
		var markupRow []*InlineKeyboardButton

		websiteURLButton := InlineKeyboardButton{Text: "Go to website", URL: feed.SiteURL}
		markupRow = append(markupRow, &websiteURLButton)

		articleURLButton := InlineKeyboardButton{Text: "Go to article", URL: entry.URL}
		markupRow = append(markupRow, &articleURLButton)

		if entry.CommentsURL != "" {
			commentURLButton := InlineKeyboardButton{Text: "Comments", URL: entry.CommentsURL}
			markupRow = append(markupRow, &commentURLButton)
		}

		message.ReplyMarkup = &InlineKeyboard{}
		message.ReplyMarkup.InlineKeyboard = append(message.ReplyMarkup.InlineKeyboard, markupRow)
	}

	client := NewClient(botToken, chatID)
	_, err := client.SendMessage(message)
	return err
}
