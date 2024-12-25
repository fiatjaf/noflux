package nostr

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fiatjaf/noflux/internal/model"
	"github.com/fiatjaf/noflux/internal/reader/processor"
	"github.com/fiatjaf/noflux/internal/storage"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip05"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/nbd-wtf/go-nostr/nip23"
	"github.com/nbd-wtf/go-nostr/sdk"
	"github.com/nbd-wtf/go-nostr/sdk/hints/sqlh"
)

var NostrSdk *sdk.System

func Initialize(db *sql.DB) error {
	hdb, err := sqlh.NewSQLHints(db, "postgres")
	if err != nil {
		return fmt.Errorf("failed to initialize nostr hints db: %w", err)
	}

	NostrSdk = sdk.NewSystem(
		sdk.WithHintsDB(hdb),
	)

	return nil
}

func GetIcon(feed *model.Feed) (bool, string) {
	yes, profile := IsItNostr(feed.FeedURL)
	if yes {
		return true, profile.Picture
	}
	return false, ""
}

func CreateFeed(
	store *storage.Storage,
	userID int64,
	req *model.FeedCreationRequest,
	profile *sdk.ProfileMetadata,
) (*model.Feed, error) {
	ctx := context.Background()
	subscription := &model.Feed{}
	nprofile := profile.Nprofile(ctx, NostrSdk, 3)

	subscription.Title = profile.Name
	subscription.UserID = userID
	subscription.UserAgent = req.UserAgent
	subscription.Cookie = req.Cookie
	subscription.Username = req.Username
	subscription.Password = req.Password
	subscription.Crawler = req.Crawler
	subscription.FetchViaProxy = req.FetchViaProxy
	subscription.HideGlobally = req.HideGlobally
	subscription.FeedURL = fmt.Sprintf("nostr:%s", nprofile)
	subscription.SiteURL = fmt.Sprintf("https://njump.me/%s", nprofile)
	subscription.WithCategoryID(req.CategoryID)
	subscription.CheckedNow()

	if err := store.CreateFeed(subscription); err != nil {
		return nil, err
	}

	if err := RefreshFeed(store, userID, subscription, profile, false); err != nil {
		return nil, err
	}

	return subscription, nil
}

func RefreshFeed(
	store *storage.Storage,
	userID int64,
	originalFeed *model.Feed,
	profile *sdk.ProfileMetadata,
	forceRefresh bool,
) error {
	ctx := context.Background()

	relays := NostrSdk.FetchOutboxRelays(ctx, profile.PubKey, 3)
	evchan := NostrSdk.Pool.SubManyEose(ctx, relays, nostr.Filters{
		{
			Authors: []string{profile.PubKey},
			Kinds:   []int{nostr.KindArticle},
			Limit:   300,
		},
	})
	updatedFeed := originalFeed
	for event := range evchan {
		eventTime := event.CreatedAt.Time()
		publishedAt := eventTime
		if paTag := event.Tags.GetFirst([]string{"published_at", ""}); paTag != nil && len(*paTag) >= 2 {
			i, err := strconv.ParseInt((*paTag)[1], 10, 64)
			if err == nil {
				publishedAt = time.Unix(i, 0)
			}
		}

		naddr, err := nip19.EncodeEntity(event.PubKey, event.Kind, event.Tags.GetD(), relays)
		if err != nil {
			continue
		}

		title := ""
		titleTag := event.Tags.GetFirst([]string{"title", ""})
		if titleTag != nil && len(*titleTag) >= 2 {
			title = (*titleTag)[1]
		}

		entry := &model.Entry{
			Date:      publishedAt,
			CreatedAt: publishedAt,
			ChangedAt: eventTime,
			Title:     title,
			Content:   replaceNostrURLsWithHTMLTags(nip23.MarkdownToHTML(event.Content)),
			URL:       fmt.Sprintf("https://njump.me/%s", naddr),
			Hash:      fmt.Sprintf("nostr:%s:%s", event.PubKey, event.Tags.GetD()),
		}

		updatedFeed.Entries = append(updatedFeed.Entries, entry)
	}

	processor.ProcessFeedEntries(store, updatedFeed, userID, forceRefresh)

	_, err := store.RefreshFeedEntries(originalFeed.UserID, originalFeed.ID, updatedFeed.Entries, forceRefresh)
	if err != nil {
		return fmt.Errorf("failed to store refreshed items: %w", err)
	}

	return nil
}

func IsItNostr(candidateUrl string) (bool, *sdk.ProfileMetadata) {
	url := candidateUrl
	ctx := context.Background()

	// check for nostr url prefixes
	hasNostrPrefix := false
	if strings.HasPrefix(url, "nostr://") {
		hasNostrPrefix = true
		url = url[8:]
	} else if strings.HasPrefix(url, "nostr:") {
		hasNostrPrefix = true
		url = url[6:]
	}

	// check for npub or nprofile
	if prefix, _, err := nip19.Decode(url); err == nil {
		if prefix == "nprofile" || prefix == "npub" {
			profile, err := NostrSdk.FetchProfileFromInput(ctx, url)
			if err != nil {
				return false, nil
			}
			return true, &profile
		}
	}

	// only do nip05 check when nostr prefix
	if hasNostrPrefix && nip05.IsValidIdentifier(url) {
		profile, err := NostrSdk.FetchProfileFromInput(ctx, url)
		if err != nil {
			return false, nil
		}
		return true, &profile
	}

	return false, nil
}
