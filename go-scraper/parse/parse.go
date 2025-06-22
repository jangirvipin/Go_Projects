package parse

import "strings"

func ValidLinksOnly(link string) bool {
	if strings.HasPrefix(link, "#") || strings.HasPrefix(link, "mailto") || link == "" {
		return false
	}
	if !strings.HasPrefix(link, "https://www.theguardian.com") {
		return false
	}
	if strings.Contains(link, "/info") || strings.Contains(link, "/help") || strings.Contains(link, "/index") {
		return false
	}
	return true
}

func Normalize(link string) string {
	if strings.HasPrefix(link, "/") {
		return "https://www.theguardian.com" + link
	}
	return link
}
