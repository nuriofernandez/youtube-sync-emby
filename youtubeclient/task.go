package youtubeclient

import (
	"YoutubeDownloader/mutubeclient"
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

func Run() {
	for channelTag, durationLenghtMinimum := range channels {
		content := scrap("https://www.youtube.com/" + channelTag + "/videos")
		videos := videoExtractor(content)

		// Download thumbnails
		for videoLink, duration := range videos {
			if len(duration) <= durationLenghtMinimum {
				fmt.Println("[Task] '" + channelTag + "' Skipping video '" + videoLink + "' due to duration limit. (" + duration + ")")
				continue
			}

			fmt.Println("[Task] '" + channelTag + "' Downloading thumbnail '" + videoLink + "' ...")
			mutubeclient.RefreshThumbnail("https://www.youtube.com" + videoLink)
		}

		time.Sleep(10 * time.Second)

		// Download video
		for videoLink, duration := range videos {
			if len(duration) <= durationLenghtMinimum {
				fmt.Println("[Task] '" + channelTag + "' Skipping video '" + videoLink + "' due to duration limit. (" + duration + ")")
				continue
			}

			fmt.Println("[Task] '" + channelTag + "' Downloading video '" + videoLink + "' ...")
			mutubeclient.Queue("https://www.youtube.com" + videoLink)
		}
	}

}
