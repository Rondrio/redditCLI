package commands

import (
	"context"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"log"
	"strconv"
	"strings"
)

const (
	DAILY = "daily"
	WEEKLY = "weekly"
	MONTHLY = "monthly"
	YEARLY = "yearly"
	ALL = "all"
)

type getPostsStruct struct {
}

func (getPosts getPostsStruct) getKeyword() string{
 	return "get"
}

func (getPosts getPostsStruct) execute(client *redditClient, command []string) ([]*reddit.Post,*reddit.Response,error) {
	var err error

	for i,word := range command {
		if strings.HasPrefix(word,"-") {
			switch word {
			case "-s":
				subreddit = command[i+1]
			case "-c":
				count,err = strconv.Atoi(command[i+1])
				if err != nil {
					count = 1
				}
			case "-t":
				switch command[i+1] {
				case DAILY:
					time = DAILY
				case WEEKLY:
					time = WEEKLY
				case MONTHLY:
					time = MONTHLY
				case YEARLY:
					time = YEARLY
				case ALL:
					time = ALL
				}
			}
		}
	}

	log.Printf("getting top %d posts of %s from %s",count,time,subreddit)
	posts,resp,err := client.Subreddit.TopPosts(context.Background(),subreddit,&reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: count,
		},
		Time: time,
	})
	if err != nil {
		return nil,nil,err
	}

	client.response = resp
	return posts,resp,nil
}