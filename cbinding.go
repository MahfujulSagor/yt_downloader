package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"github/MahfujulSagor/yt_downloader/downloader"
)

//export GetVideoInfo
func GetVideoInfo(cUrl *C.char) *C.char {
	url := C.GoString(cUrl)
	info, err := downloader.FetchVideoInfo(url)
	if err != nil {
		return C.CString(`{"error":"` + err.Error() + `"}`)
	}
	b, _ := json.Marshal(info)
	return C.CString(string(b))
}

//export StartDownload
func StartDownload(cUrl, cPath *C.char) C.int {
	url := C.GoString(cUrl)
	path := C.GoString(cPath)
	id := downloader.StartDownload(url, path)
	return C.int(id)
}

//export GetStatus
func GetStatus(id C.int) *C.char {
	status := downloader.GetDownloadStatus(int(id))
	b, _ := json.Marshal(status)
	return C.CString(string(b))
}

//export MergeFiles
func MergeFiles(cVideo, cAudio, cOut *C.char) C.int {
	video := C.GoString(cVideo)
	audio := C.GoString(cAudio)
	out := C.GoString(cOut)
	if err := downloader.MergeAudioVideo(video, audio, out); err != nil {
		return 0
	}
	return 1
}

func main() {}
