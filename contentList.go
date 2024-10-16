package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Lists struct {
	Horizontal []HorizontalLists `json:"horizontalLists"`
	Content    []Content         `json:"content"`
}

type HorizontalLists struct {
	Title   string    `json:"title"`
	Content []Content `json:"content"`
}
type SearchResult struct {
	Content []Content `json:"content"`
}

type Content struct {
	Type              string           `json:"type"`
	YoutubeId         string           `json:"youtubeId"`
	Title             string           `json:"title"`
	Description       string           `json:"description"`
	Thumbnail         string           `json:"thumbnail"`
	Hls               string           `json:"hls"`
	UploadedDate      string           `json:"uploadedDate"`
	Duration          int              `json:"duration"`
	ViewsCount        int              `json:"viewsCount"`
	WatchedTime       int              `json:"watchedTime"`
	VideosCount       int              `json:"videosCount"`
	PlaylistVideos    []PlaylistVideos `json:"playlistVideos"`
	ChannelYoutubeId  string           `json:"channelYoutubeId"`
	ChannelName       string           `json:"channelName"`
	ChannelAvatarUrl  string           `json:"channelAvatarUrl"`
	IsSubscribed      bool             `json:"isSubscribed"`
	SubscribersCount  int              `json:"subscribersCount"`
	Source            string           `json:"source"`
	ChannelIsVerified bool             `json:"channelIsVerified"`
}

type PlaylistVideos struct {
	YoutubeId string `json:"youtubeId"`
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
}

func UnmarshalList() {
	var list SearchResult

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error!", r)
		}
	}()

	jsonFile, err := os.ReadFile("jsonFiles/test2.json")
	if err != nil {
		panic(err)
	}

	er := json.Unmarshal([]byte(jsonFile), &list)
	if er != nil {
		panic(er)
	}
	fmt.Println("OK")

	// fmt.Println("Titles: ")
	// for i := range list.Content {
	// 	fmt.Println(list.Content[i].Title)
	// }

	fmt.Println("\nplayVideos: ")
	empty := 0
	notEmpty := 0
	for _, val := range list.Content {

		if len(val.PlaylistVideos) == 0 {
			empty++
		}
		if len(val.PlaylistVideos) > 0 {
			notEmpty++
			fmt.Println(val.PlaylistVideos)
		}

	}
	fmt.Printf("%d empty slices\n", empty)
	fmt.Printf("%d not empty slices\n", notEmpty)

}
