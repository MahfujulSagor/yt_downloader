package downloader

import (
	"fmt"
	"os/exec"
)

// MergeAudioVideo merges video+audio into one file using ffmpeg
func MergeAudioVideo(videoPath, audioPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-y",
		"-i", videoPath,
		"-i", audioPath,
		"-c:v", "copy",
		"-c:a", "aac",
		outputPath,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("merge failed: %v, output: %s", err, string(output))
	}
	return nil
}
