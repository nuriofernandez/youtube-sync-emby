package youtubeclient

import (
	"fmt"
	"regexp"
	"strings"
)

func videoExtractor(str string) map[string]string {
	videos := make(map[string]string)

	split := strings.Split(str, "lengthText")
	for _, v := range split {

		// Content parts
		i := strings.Split(v, "webCommandMetadata")

		// Split
		durationHtml := i[0]
		videoLinkHtml := strings.Split(i[1], "rootVe")[0]

		// Extraction
		link := extractWatch(videoLinkHtml)
		if len(link) == 0 {
			continue
		}
		duration := extractDuration(durationHtml)
		if len(duration) == 0 {
			continue
		}

		fmt.Println("[Scrapper] Video '" + link + "' with duration '" + duration + "'.")
		videos[link] = duration
	}

	return videos
}

var watchRegExp = regexp.MustCompile(`(?m)\/watch\?v=\w{10,}`)

func extractWatch(html string) string {
	allLinks := watchRegExp.FindAllString(html, 1)
	if len(allLinks) == 0 {
		return ""
	}
	return allLinks[0]
}

var durationRegExp = regexp.MustCompile(`(?m)\d{1,2}:\d{2}(:\d{2})?`)

func extractDuration(html string) string {
	allDurations := durationRegExp.FindAllString(html, 1)
	if len(allDurations) == 0 {
		return ""
	}
	return allDurations[0]
}
