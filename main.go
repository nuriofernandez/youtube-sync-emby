package main

import (
	"YoutubeDownloader/youtubeclient"
	"fmt"
	"time"
)

func main() {
	youtubeclient.Run()

	fmt.Println("Waiting a minute to retry...")
	time.Sleep(1 * time.Minute)

	youtubeclient.Run()
}
