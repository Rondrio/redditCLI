package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"redditCLI/GUI"
	"redditCLI/commands"
	"strings"
)

func main() {
	color.HiYellow("Hello")
	color.Yellow("There")

	client := commands.GetInstance()
	go client.ReadConsole()
	acceptPosts(client.RespChan)
}

func acceptPosts(acceptChan chan *commands.Response) {
	for {
		select {
		case resp := <-acceptChan:
			for _,post := range resp.Posts {
				log.Println(post.Title)
				log.Println(post.URL)
				meme,err := getMeme(post.URL)
				if err != nil {
					log.Println(err)
				}
				if meme != nil {
					GUI.PaintMeme(meme)
				}
			}
		}
	}
}

func getMeme(url string) (image.Image,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("failed getting meme %d", resp.StatusCode))
	}

	switch {
	case strings.HasSuffix(url, "jpg"):
		return jpeg.Decode(resp.Body)
	case strings.HasSuffix(url, "png"):
		return png.Decode(resp.Body)
	case strings.HasSuffix(url, "gif"):
		return gif.Decode(resp.Body)

	default:
		log.Println("image type not supported")
		return nil, nil //errors.New("file type not supported")

	}
}