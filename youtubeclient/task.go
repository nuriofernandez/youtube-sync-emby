package youtubeclient

import (
	"fmt"
	"time"
)

var channels = map[string]int{
	"@spicy4tuna":         5,
	"@juanrallo":          0,
	"@SoloFonseca":        0,
	"@VisualEconomik":     0,
	"@NOTICIASILUSTRADAS": 0,
	"@HRom":               0,
	"@vicesat":            0,
	"@Fireship":           0,
	"@KiraSensei1":        0,
}

func FetchVideos() []string {
	videoLinks := make([]string, 0)

	for channelTag, durationLenghtMinimum := range channels {
		content := scrap("https://www.youtube.com/" + channelTag + "/videos")
		videos := videoExtractor(content)

		// Download thumbnails
		for videoLink, duration := range videos {
			if len(duration) <= durationLenghtMinimum {
				fmt.Println("[YT Client] '" + channelTag + "' Skipping video '" + videoLink + "' due to duration limit. (" + duration + ")")
				continue
			}

			fmt.Println("[YT Client] '" + channelTag + "' queueing video '" + videoLink + "' ...")
			videoLinks = append(videoLinks, videoLink)
		}

		fmt.Println("[YT Client] '" + channelTag + "' fetching is completed. Waiting 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	return videoLinks
}
