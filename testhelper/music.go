package testhelper

import "github.com/breaking-fullstack/forever-server/entity"

var MusicList = []entity.Music{
	{
		ID:    "1",
		Title: "Foo Bars",
		URL:   "http://foo.bars/music",
	},
	{
		ID:    "2",
		Title: "Baz Bass",
		URL:   "http://baz.music",
	},
}
