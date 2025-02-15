// SPDX-FileCopyrightText: Copyright The Noflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package integration // import "github.com/fiatjaf/noflux/internal/integration"

import (
	"log/slog"

	"github.com/fiatjaf/noflux/internal/config"
	"github.com/fiatjaf/noflux/internal/integration/apprise"
	"github.com/fiatjaf/noflux/internal/integration/betula"
	"github.com/fiatjaf/noflux/internal/integration/cubox"
	"github.com/fiatjaf/noflux/internal/integration/discord"
	"github.com/fiatjaf/noflux/internal/integration/espial"
	"github.com/fiatjaf/noflux/internal/integration/instapaper"
	"github.com/fiatjaf/noflux/internal/integration/linkace"
	"github.com/fiatjaf/noflux/internal/integration/linkding"
	"github.com/fiatjaf/noflux/internal/integration/linkwarden"
	"github.com/fiatjaf/noflux/internal/integration/matrixbot"
	"github.com/fiatjaf/noflux/internal/integration/notion"
	"github.com/fiatjaf/noflux/internal/integration/ntfy"
	"github.com/fiatjaf/noflux/internal/integration/nunuxkeeper"
	"github.com/fiatjaf/noflux/internal/integration/omnivore"
	"github.com/fiatjaf/noflux/internal/integration/pinboard"
	"github.com/fiatjaf/noflux/internal/integration/pocket"
	"github.com/fiatjaf/noflux/internal/integration/raindrop"
	"github.com/fiatjaf/noflux/internal/integration/readeck"
	"github.com/fiatjaf/noflux/internal/integration/readwise"
	"github.com/fiatjaf/noflux/internal/integration/shaarli"
	"github.com/fiatjaf/noflux/internal/integration/shiori"
	"github.com/fiatjaf/noflux/internal/integration/telegrambot"
	"github.com/fiatjaf/noflux/internal/integration/wallabag"
	"github.com/fiatjaf/noflux/internal/integration/webhook"
	"github.com/fiatjaf/noflux/internal/model"
)

// SendEntry sends the entry to third-party providers when the user click on "Save".
func SendEntry(entry *model.Entry, userIntegrations *model.Integration) {
	if userIntegrations.BetulaEnabled {
		slog.Debug("Sending entry to Betula",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := betula.NewClient(userIntegrations.BetulaURL, userIntegrations.BetulaToken)
		err := client.CreateBookmark(
			entry.URL,
			entry.Title,
			entry.Tags,
		)
		if err != nil {
			slog.Error("Unable to send entry to Betula",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.PinboardEnabled {
		slog.Debug("Sending entry to Pinboard",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := pinboard.NewClient(userIntegrations.PinboardToken)
		err := client.CreateBookmark(
			entry.URL,
			entry.Title,
			userIntegrations.PinboardTags,
			userIntegrations.PinboardMarkAsUnread,
		)
		if err != nil {
			slog.Error("Unable to send entry to Pinboard",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.InstapaperEnabled {
		slog.Debug("Sending entry to Instapaper",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := instapaper.NewClient(userIntegrations.InstapaperUsername, userIntegrations.InstapaperPassword)
		if err := client.AddURL(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Instapaper",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.WallabagEnabled {
		slog.Debug("Sending entry to Wallabag",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := wallabag.NewClient(
			userIntegrations.WallabagURL,
			userIntegrations.WallabagClientID,
			userIntegrations.WallabagClientSecret,
			userIntegrations.WallabagUsername,
			userIntegrations.WallabagPassword,
			userIntegrations.WallabagOnlyURL,
		)

		if err := client.CreateEntry(entry.URL, entry.Title, entry.Content); err != nil {
			slog.Error("Unable to send entry to Wallabag",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.NotionEnabled {
		slog.Debug("Sending entry to Notion",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := notion.NewClient(
			userIntegrations.NotionToken,
			userIntegrations.NotionPageID,
		)
		if err := client.UpdateDocument(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Notion",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.NunuxKeeperEnabled {
		slog.Debug("Sending entry to NunuxKeeper",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := nunuxkeeper.NewClient(
			userIntegrations.NunuxKeeperURL,
			userIntegrations.NunuxKeeperAPIKey,
		)

		if err := client.AddEntry(entry.URL, entry.Title, entry.Content); err != nil {
			slog.Error("Unable to send entry to NunuxKeeper",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.EspialEnabled {
		slog.Debug("Sending entry to Espial",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := espial.NewClient(
			userIntegrations.EspialURL,
			userIntegrations.EspialAPIKey,
		)

		if err := client.CreateLink(entry.URL, entry.Title, userIntegrations.EspialTags); err != nil {
			slog.Error("Unable to send entry to Espial",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.PocketEnabled {
		slog.Debug("Sending entry to Pocket",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := pocket.NewClient(config.Opts.PocketConsumerKey(userIntegrations.PocketConsumerKey), userIntegrations.PocketAccessToken)
		if err := client.AddURL(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Pocket",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.LinkAceEnabled {
		slog.Debug("Sending entry to LinkAce",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := linkace.NewClient(
			userIntegrations.LinkAceURL,
			userIntegrations.LinkAceAPIKey,
			userIntegrations.LinkAceTags,
			userIntegrations.LinkAcePrivate,
			userIntegrations.LinkAceCheckDisabled,
		)
		if err := client.AddURL(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to LinkAce",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.LinkdingEnabled {
		slog.Debug("Sending entry to Linkding",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := linkding.NewClient(
			userIntegrations.LinkdingURL,
			userIntegrations.LinkdingAPIKey,
			userIntegrations.LinkdingTags,
			userIntegrations.LinkdingMarkAsUnread,
		)
		if err := client.CreateBookmark(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Linkding",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.LinkwardenEnabled {
		slog.Debug("Sending entry to linkwarden",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := linkwarden.NewClient(
			userIntegrations.LinkwardenURL,
			userIntegrations.LinkwardenAPIKey,
		)
		if err := client.CreateBookmark(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Linkwarden",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.ReadeckEnabled {
		slog.Debug("Sending entry to Readeck",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := readeck.NewClient(
			userIntegrations.ReadeckURL,
			userIntegrations.ReadeckAPIKey,
			userIntegrations.ReadeckLabels,
			userIntegrations.ReadeckOnlyURL,
		)
		if err := client.CreateBookmark(entry.URL, entry.Title, entry.Content); err != nil {
			slog.Error("Unable to send entry to Readeck",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.ReadwiseEnabled {
		slog.Debug("Sending entry to Readwise",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := readwise.NewClient(
			userIntegrations.ReadwiseAPIKey,
		)

		if err := client.CreateDocument(entry.URL); err != nil {
			slog.Error("Unable to send entry to Readwise",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.CuboxEnabled {
		slog.Debug("Sending entry to Cubox",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := cubox.NewClient(userIntegrations.CuboxAPILink)

		if err := client.SaveLink(entry.URL); err != nil {
			slog.Error("Unable to send entry to Cubox",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.ShioriEnabled {
		slog.Debug("Sending entry to Shiori",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := shiori.NewClient(
			userIntegrations.ShioriURL,
			userIntegrations.ShioriUsername,
			userIntegrations.ShioriPassword,
		)

		if err := client.CreateBookmark(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Shiori",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.ShaarliEnabled {
		slog.Debug("Sending entry to Shaarli",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := shaarli.NewClient(
			userIntegrations.ShaarliURL,
			userIntegrations.ShaarliAPISecret,
		)

		if err := client.CreateLink(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Shaarli",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.WebhookEnabled {
		slog.Debug("Sending entry to Webhook",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
			slog.String("webhook_url", userIntegrations.WebhookURL),
		)

		webhookClient := webhook.NewClient(userIntegrations.WebhookURL, userIntegrations.WebhookSecret)
		if err := webhookClient.SendSaveEntryWebhookEvent(entry); err != nil {
			slog.Error("Unable to send entry to Webhook",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.String("webhook_url", userIntegrations.WebhookURL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.OmnivoreEnabled {
		slog.Debug("Sending entry to Omnivore",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := omnivore.NewClient(userIntegrations.OmnivoreAPIKey, userIntegrations.OmnivoreURL)
		if err := client.SaveUrl(entry.URL); err != nil {
			slog.Error("Unable to send entry to Omnivore",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.RaindropEnabled {
		slog.Debug("Sending entry to Raindrop",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int64("entry_id", entry.ID),
			slog.String("entry_url", entry.URL),
		)

		client := raindrop.NewClient(userIntegrations.RaindropToken, userIntegrations.RaindropCollectionID, userIntegrations.RaindropTags)
		if err := client.CreateRaindrop(entry.URL, entry.Title); err != nil {
			slog.Error("Unable to send entry to Raindrop",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int64("entry_id", entry.ID),
				slog.String("entry_url", entry.URL),
				slog.Any("error", err),
			)
		}
	}
}

// PushEntries pushes a list of entries to activated third-party providers during feed refreshes.
func PushEntries(feed *model.Feed, entries model.Entries, userIntegrations *model.Integration) {
	if userIntegrations.MatrixBotEnabled {
		slog.Debug("Sending new entries to Matrix",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int("nb_entries", len(entries)),
			slog.Int64("feed_id", feed.ID),
		)

		err := matrixbot.PushEntries(
			feed,
			entries,
			userIntegrations.MatrixBotURL,
			userIntegrations.MatrixBotUser,
			userIntegrations.MatrixBotPassword,
			userIntegrations.MatrixBotChatID,
		)
		if err != nil {
			slog.Error("Unable to send new entries to Matrix",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int("nb_entries", len(entries)),
				slog.Int64("feed_id", feed.ID),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.WebhookEnabled {
		slog.Debug("Sending new entries to Webhook",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int("nb_entries", len(entries)),
			slog.Int64("feed_id", feed.ID),
			slog.String("webhook_url", userIntegrations.WebhookURL),
		)

		webhookClient := webhook.NewClient(userIntegrations.WebhookURL, userIntegrations.WebhookSecret)
		if err := webhookClient.SendNewEntriesWebhookEvent(feed, entries); err != nil {
			slog.Debug("Unable to send new entries to Webhook",
				slog.Int64("user_id", userIntegrations.UserID),
				slog.Int("nb_entries", len(entries)),
				slog.Int64("feed_id", feed.ID),
				slog.String("webhook_url", userIntegrations.WebhookURL),
				slog.Any("error", err),
			)
		}
	}

	if userIntegrations.NtfyEnabled && feed.NtfyEnabled {
		slog.Debug("Sending new entries to Ntfy",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int("nb_entries", len(entries)),
			slog.Int64("feed_id", feed.ID),
		)

		client := ntfy.NewClient(
			userIntegrations.NtfyURL,
			userIntegrations.NtfyTopic,
			userIntegrations.NtfyAPIToken,
			userIntegrations.NtfyUsername,
			userIntegrations.NtfyPassword,
			userIntegrations.NtfyIconURL,
			userIntegrations.NtfyInternalLinks,
			feed.NtfyPriority,
		)

		if err := client.SendMessages(feed, entries); err != nil {
			slog.Warn("Unable to send new entries to Ntfy", slog.Any("error", err))
		}
	}

	if userIntegrations.AppriseEnabled {
		slog.Debug("Sending new entries to Apprise",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int("nb_entries", len(entries)),
			slog.Int64("feed_id", feed.ID),
		)

		appriseServiceURLs := userIntegrations.AppriseServicesURL
		if feed.AppriseServiceURLs != "" {
			appriseServiceURLs = feed.AppriseServiceURLs
		}

		client := apprise.NewClient(
			appriseServiceURLs,
			userIntegrations.AppriseURL,
		)

		if err := client.SendNotification(feed, entries); err != nil {
			slog.Warn("Unable to send new entries to Apprise", slog.Any("error", err))
		}
	}

	if userIntegrations.DiscordEnabled {
		slog.Debug("Sending new entries to Discord",
			slog.Int64("user_id", userIntegrations.UserID),
			slog.Int("nb_entries", len(entries)),
			slog.Int64("feed_id", feed.ID),
		)

		client := discord.NewClient(
			userIntegrations.DiscordWebhookLink,
		)

		if err := client.SendDiscordMsg(feed, entries); err != nil {
			slog.Warn("Unable to send new entries to Discord", slog.Any("error", err))
		}
	}

	// Integrations that only support sending individual entries
	if userIntegrations.TelegramBotEnabled {
		for _, entry := range entries {
			if userIntegrations.TelegramBotEnabled {
				slog.Debug("Sending a new entry to Telegram",
					slog.Int64("user_id", userIntegrations.UserID),
					slog.Int64("entry_id", entry.ID),
					slog.String("entry_url", entry.URL),
				)

				if err := telegrambot.PushEntry(
					feed,
					entry,
					userIntegrations.TelegramBotToken,
					userIntegrations.TelegramBotChatID,
					userIntegrations.TelegramBotTopicID,
					userIntegrations.TelegramBotDisableWebPagePreview,
					userIntegrations.TelegramBotDisableNotification,
					userIntegrations.TelegramBotDisableButtons,
				); err != nil {
					slog.Error("Unable to send entry to Telegram",
						slog.Int64("user_id", userIntegrations.UserID),
						slog.Int64("entry_id", entry.ID),
						slog.String("entry_url", entry.URL),
						slog.Any("error", err),
					)
				}
			}
		}
	}
}
