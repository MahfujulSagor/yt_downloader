package downloader

import (
	"fmt"

	"github.com/kkdai/youtube/v2"
)

// VideoInfo holds video metadata
type VideoInfo struct {
	Title    string
	Author   string
	Duration string
	URL      string
}

// FetchVideoInfo retrieves metadata from YouTube
func FetchVideoInfo(url string) (*VideoInfo, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch video: %w", err)
	}

	info := &VideoInfo{
		Title:    video.Title,
		Author:   video.Author,
		Duration: video.Duration.String(),
		URL:      url,
	}
	return info, nil
}
