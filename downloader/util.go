package downloader

import "regexp"

// SanitizeFilename makes a safe filename
func SanitizeFilename(name string) string {
	re := regexp.MustCompile(`[<>:"/\\|?*]`)
	return re.ReplaceAllString(name, "_")
}
