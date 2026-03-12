package mutubeclient

var store *Store

func registerStorage() {
	if store == nil {
		tmpStore, err := NewStore("cached_links.txt")
		if err != nil {
			panic(err)
		}
		store = tmpStore
	}
}

func Queue(link string) {
	registerStorage()

	if store.Exists(link + "-mp4") {
		return
	}

	download(link, "mp4")
	download(link, "thumbnail")

	store.Save(link + "-mp4")
}

func RefreshThumbnail(link string) {
	registerStorage()

	if store.Exists(link + "-thumbnail") {
		return
	}

	download(link, "thumbnail")
	store.Save(link + "-thumbnail")
}
