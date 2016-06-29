package main

import "github.com/felipejfc/go-torrent-finder"

func main() {
	t := finder.GetTorrentFinder()
	t.SearchTorrents("game of thrones s06e10")
}
