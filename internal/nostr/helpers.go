package nostr

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/puzpuzpuz/xsync/v3"
)

var nostrEveryMatcher = regexp.MustCompile(`nostr:((npub|note|nevent|nprofile|naddr)1[a-z0-9]+)\b`)

func replaceNostrURLsWithHTMLTags(input string) string {
	// match and replace npup1, nprofile1, note1, nevent1, etc
	names := xsync.NewMapOf[string, string]()
	wg := sync.WaitGroup{}

	// first we run it without waiting for the results of getting the name as they will be async
	for _, match := range nostrEveryMatcher.FindAllString(input, len(input)+1) {
		nip19 := match[len("nostr:"):]

		if strings.HasPrefix(nip19, "npub1") || strings.HasPrefix(nip19, "nprofile1") {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
			defer cancel()
			wg.Add(1)
			go func() {
				metadata, _ := NostrSdk.FetchProfileFromInput(ctx, nip19)
				if metadata.Name != "" {
					names.Store(nip19, metadata.Name)
				}
				wg.Done()
			}()
		}
	}

	// in the second time now that we got all the names we actually perform replacement
	wg.Wait()
	return nostrEveryMatcher.ReplaceAllStringFunc(input, func(match string) string {
		nip19 := match[len("nostr:"):]
		firstChars := nip19[:8]
		lastChars := nip19[len(nip19)-4:]

		if strings.HasPrefix(nip19, "npub1") || strings.HasPrefix(nip19, "nprofile1") {
			name, _ := names.Load(nip19)
			if name != "" {
				return fmt.Sprintf(`<a href="https://njump.me/%s">%s (%s…%s)</a>`, nip19, name, firstChars, lastChars)
			} else {
				return fmt.Sprintf(`<a href="https://njump.me/%s">%s…%s</a>`, nip19, firstChars, lastChars)
			}
		} else {
			return fmt.Sprintf(`<a href="https://njump.me/%s">%s…%s</a>`, nip19, firstChars, lastChars)
		}
	})
}
