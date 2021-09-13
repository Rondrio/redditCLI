package commands

import (
	"context"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"log"
	"strconv"
	"strings"
)

type prevPostsStruct struct {
}


func (nextPosts prevPostsStruct) getKeyword() string{
	return "prev"
}

func (nextPosts prevPostsStruct) execute(client *redditClient, command []string) ([]*reddit.Post,*reddit.Response,error) {
	var err error

	for i,word := range command {
		if strings.HasPrefix(word,"-") {
			switch word {
			case "-c":
				count,err = strconv.Atoi(command[i+1])
				if err != nil {
					count = 1
				}
			}
		}
	}

	log.Printf("getting top %d posts of %s from %s",count,time,subreddit)
	return client.Subreddit.TopPosts(context.Background(),subreddit,&reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: count,
			Before: client.response.After,
		},
		Time: time,
	})
}