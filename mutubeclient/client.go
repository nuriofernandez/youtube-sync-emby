package mutubeclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func download(link, format string) {
	url := "http://emby.nurio.me:8081/add"
	method := "POST"

	payload := strings.NewReader(`{
    "url": "` + link + `",
    "quality": "best",
    "format": "` + format + `",
    "auto_start": true,
    "split_by_chapters": false,
    "chapter_template": "%(title)s - %(section_number)02d - %(section_title)s.%(ext)s",
    "subtitle_format": "srt",
    "subtitle_language": "en",
    "subtitle_mode": "prefer_manual"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("[MuTube] '" + format + "' @ " + strings.Split(link, "watch?v=")[1] + " " + string(body))
}
