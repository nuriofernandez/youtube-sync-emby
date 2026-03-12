package main

import (
	"YoutubeDownloader/mutubeclient"
	"YoutubeDownloader/youtubeclient"
	"fmt"
	"time"
)

func main() {
	fmt.Println("[YoutubeDownloader] Starting...")

	fmt.Println("[YoutubeDownloader] Fetching videos...")
	videos := youtubeclient.FetchVideos()

	fmt.Println("[YoutubeDownloader] All videos are now queued.")

	fmt.Println("[YoutubeDownloader] Starting thumbnail downloading...")
	for i, videoLink := range videos {
		fmt.Printf("[YoutubeDownloader] (%d/%d) Downloading thumbnail '%s' ...\n", i+1, len(videos), videoLink)
		mutubeclient.FetchThumbnail("https://www.youtube.com" + videoLink)
	}

	fmt.Println("[YoutubeDownloader] Thumbnail downloading completed! Waiting 1 minute...")
	time.Sleep(1 * time.Minute)

	fmt.Println("[YoutubeDownloader] Starting thumbnail refreshing...")
	for i, videoLink := range videos {
		fmt.Printf("[YoutubeDownloader] (%d/%d) Refreshing thumbnail '%s' ...\n", i+1, len(videos), videoLink)
		mutubeclient.RefreshThumbnail("https://www.youtube.com" + videoLink)
	}

	fmt.Println("[YoutubeDownloader] Thumbnail downloading completed! Waiting 20 seconds...")
	time.Sleep(20 * time.Second)

	fmt.Println("[YoutubeDownloader] Starting video downloading...")
	for i, videoLink := range videos {
		fmt.Printf("[YoutubeDownloader] (%d/%d) Downloading video '%s' ...\n", i+1, len(videos), videoLink)
		mutubeclient.FetchVideo("https://www.youtube.com" + videoLink)
	}

	fmt.Println("[YoutubeDownloader] Video downloading completed!")
}
