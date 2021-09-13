package commands

import (
	"bufio"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"log"
	"os"
	"strings"
)

var (
	subreddit = "ich_iel"
	count = 5
	time = "all"
)

type iCommand interface {
	getKeyword() string
	execute(client *redditClient, command []string) ([]*reddit.Post,*reddit.Response,error)
}

type InputReader struct {
	commands     []iCommand
	RedditClient *redditClient
	RespChan     chan *Response
}

type Response struct {
	Posts []*reddit.Post
	Response *reddit.Response
}

type redditClient struct {
	*reddit.Client
	response *reddit.Response
}

var inputReader *InputReader

func newInputReader() *InputReader{
	respChan := make(chan *Response)
	return &InputReader{
		commands: []iCommand{
			getPostsStruct{},
			nextPostsStruct{},
			prevPostsStruct{},
		},
		RespChan:     respChan,
		RedditClient: &redditClient{Client:reddit.DefaultClient()},
	}
}

func GetInstance() *InputReader{
	if inputReader == nil {
		return newInputReader()
	}
	return inputReader
}

func (i *InputReader) ReadConsole() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		input,err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed reading input:",err)
			return err
		}
		var resp Response
		resp.Posts,resp.Response,err = i.parseCommand(input)
		if err != nil {
			return err
		}

		i.RespChan <- &resp
	}
}

func (i *InputReader) parseCommand(input string) ([]*reddit.Post,*reddit.Response,error) {
	parts := strings.Fields(input)
	for _,command := range i.commands {
		if strings.ToLower(parts[0]) == command.getKeyword() {
			return command.execute(i.RedditClient,parts)
		}
	}
	return nil,nil,nil
}

func (i *InputReader) addCommand(cmd iCommand) {
	i.commands = append(i.commands,cmd)
}