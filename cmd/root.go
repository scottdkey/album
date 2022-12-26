/*
Copyright © 2022 Scot Key scottdkey@gmail.com

*/

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Album struct {
	AlbumId      int    `json:"albumId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

var rootCmd = &cobra.Command{
	Use:   "album",
	Short: "get photo album info by album number",
	Long:  "will call the jsonplaceholder.typicode.com/photos api to retrieve and return name and URL via get REST Call using format https://jsonplaceholder.typicode.com/photos?albumId=3 using provided album number. Will default to album 1 if nothing is provided",
	Run: func(
		cmd *cobra.Command,
		args []string) {
		argsLength := len(args) > 0

		if argsLength {
			albumId := args[0]
			albumIdNum, err := strconv.Atoi(albumId)

			if err != nil {
				fmt.Println(err)
			}

			if albumIdNum > 0 && albumIdNum < 101 {
				makeAlbumRequest(albumId)
			} else {
				fmt.Println("Error: input arg outside of 1-100 range: " + albumId)
			}
		} else {
			fmt.Println("Error: no album Id provided. fetching album 1")
			makeAlbumRequest("1")
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func parseAndPrintAlbumResponse(Data io.ReadCloser) {
	body, err := ioutil.ReadAll(Data)

	var result []Album

	marshalErr := json.Unmarshal(body, &result)

	if marshalErr != nil {
		fmt.Println(err)
	}

	for i := range result {
		album := result[i]
		fmt.Println("[" + fmt.Sprint(album.Id) + "]" + album.Title + " - " + album.Url)
	}
}

func makeAlbumRequest(albumId string) {
	URL := "https://jsonplaceholder.typicode.com/photos?albumId=" + albumId

	fmt.Println("attempting get request for album: " + albumId)
	res, err := http.Get(URL)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	parseAndPrintAlbumResponse(res.Body)
}
