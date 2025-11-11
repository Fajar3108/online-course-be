package helpers

import "strings"

func Slug(text string) string {
	slug := strings.ToLower(text)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	slug = strings.Trim(slug, "-")

	return slug
}
