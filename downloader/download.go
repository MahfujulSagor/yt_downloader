package downloader

import (
	"io"
	"net/http"
	"os"
	"sync"
)

type DownloadStatus struct {
	Progress int    // 0-100
	Error    string // error message if any
	Done     bool   // finished?
}

var (
	downloads   = make(map[int]*DownloadStatus)
	downloadsMu sync.Mutex
	nextID      = 1
)

// StartDownload runs async download
func StartDownload(url, filepath string) int {
	downloadsMu.Lock()
	id := nextID
	nextID++
	status := &DownloadStatus{Progress: 0}
	downloads[id] = status
	downloadsMu.Unlock()

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			status.Error = err.Error()
			status.Done = true
			return
		}
		defer resp.Body.Close()

		out, err := os.Create(filepath)
		if err != nil {
			status.Error = err.Error()
			status.Done = true
			return
		}
		defer out.Close()

		buf := make([]byte, 64*1024)
		var total, written int64
		if resp.ContentLength > 0 {
			total = resp.ContentLength
		}

		for {
			n, err := resp.Body.Read(buf)
			if n > 0 {
				w, _ := out.Write(buf[:n])
				written += int64(w)
				if total > 0 {
					status.Progress = int((written * 100) / total)
				}
			}
			if err != nil {
				if err == io.EOF {
					status.Progress = 100
					status.Done = true
				} else {
					status.Error = err.Error()
					status.Done = true
				}
				break
			}
		}
	}()

	return id
}

func GetDownloadStatus(id int) DownloadStatus {
	downloadsMu.Lock()
	defer downloadsMu.Unlock()
	if st, ok := downloads[id]; ok {
		return *st
	}
	return DownloadStatus{Error: "invalid id", Done: true}
}
