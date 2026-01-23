package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("regexp examples starting...")

	text := "Email: test.user+label@example.co.uk, Phone: +1-555-1234, URL: https://example.com/path?q=1"

	// 1) Compile vs MustCompile
	emailRe := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	fmt.Println("email match:", emailRe.MatchString(text))

	// Safe compile when pattern may come from user
	pattern := `\+?\d[\d\-\s]+\d`
	phoneRe, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("phone compile error:", err)
	} else {
		fmt.Println("phone found:", phoneRe.FindString(text))
	}

	// 2) FindAllString
	emails := emailRe.FindAllString(text, -1)
	fmt.Println("emails found:", emails)

	// 3) FindStringSubmatch + named groups
	urlRe := regexp.MustCompile(`(?P<proto>https?)://(?P<host>[^/]+)(?P<path>/[^?\s]*)?(?:\?(?P<query>[^\s]+))?`)
	m := urlRe.FindStringSubmatch(text)
	if m != nil {
		names := urlRe.SubexpNames()
		fmt.Println("URL named groups:")
		for i, n := range names {
			if i == 0 || n == "" { // 0 is full match
				continue
			}
			fmt.Printf("  %s = %s\n", n, m[i])
		}
	}

	// 4) ReplaceAllString and ReplaceAllStringFunc
	redacted := emailRe.ReplaceAllString(text, "[REDACTED_EMAIL]")
	fmt.Println("redacted:", redacted)

	upperURLs := urlRe.ReplaceAllStringFunc(text, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println("URL uppercased in text:", upperURLs)

	// 5) Split
	splitter := regexp.MustCompile(`\W+`)
	tokens := splitter.Split("one,two;three four", -1)
	fmt.Println("tokens:", tokens)

	// 6) QuoteMeta — escape regex metacharacters when building patterns dynamically
	raw := "+1.2*(test)?"
	quoted := regexp.QuoteMeta(raw)
	fmt.Println("quoted meta:", quoted)

	// 7) FindAllStringIndex and Submatch indexes
	idx := emailRe.FindAllStringIndex(text, -1)
	fmt.Println("email index ranges:", idx)

	// Notes: prefer MustCompile for static patterns; cache compiled regexes for reuse.
	// Avoid catastrophic backtracking by keeping patterns simple and anchored when possible.

	fmt.Println("regexp examples done")
}
