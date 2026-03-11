package mutubeclient

func Queue(link string) {
	download(link, "mp4")
	download(link, "thumbnail")
}

func RefreshThumbnail(link string) {
	download(link, "thumbnail")
}
