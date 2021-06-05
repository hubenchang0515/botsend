package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	botsend "github.com/hubenchang0515/botsend"
)

func Key() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	keyFile := path.Join(home, ".config", "botsend.key")
	data, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func SetKey(key string) {
	home, _ := os.UserHomeDir()
	keyFile := path.Join(home, ".config", "botsend.key")
	ioutil.WriteFile(keyFile, []byte(key), 0644)
}

func main() {
	key := flag.String("key", "", "Webhook key, will be saved in ~/.config/botsend")
	messageType := flag.String("type", "text", "Message type: text picture markdown url file")
	at := flag.String("at", "", "Mention someone, splited by comma e.g. '-at=137xxxxxxx2,153xxxxxxx2' or '-at=@all'")
	title := flag.String("title", "", "Title of URL message")
	description := flag.String("description", "", "Description of URL message")
	cover := flag.String("cover", "", "Cover image of URL message")
	url := flag.String("url", "", "URL of URL message")
	markdown := flag.String("markdown", "", "markdown content")
	flag.Parse()

	if len(*key) != 0 {
		SetKey(*key)
	} else {
		*key = Key()
	}

	if len(*messageType) == 0 {
		return
	}

	switch *messageType {
	case "text":
		msg := botsend.NewTextMessgae()
		msg.SetContent(flag.Arg(0))
		if *at != "" {
			msg.SetMentionString(*at)
		}
		botsend.Request(*key, msg.Json())
	case "picture":
		msg := botsend.NewPictureMessage()
		msg.SetPicture(flag.Arg(0))
		botsend.Request(*key, msg.Json())
	case "markdown":
		msg := botsend.NewMarkdownMessage()
		if *markdown == "" {
			msg.SetFile(flag.Arg(0))
		} else {
			msg.SetString(*markdown)
		}
		botsend.Request(*key, msg.Json())
	case "url":
		msg := botsend.NewUrlMessage()
		msg.SetUrl(*url)
		msg.SetTitle(*title)
		msg.SetDescription(*description)
		msg.SetCover(*cover)
		botsend.Request(*key, msg.Json())
	case "file":
		msg := botsend.NewFileMessage()
		msg.Upload(*key, flag.Arg(0))
		botsend.Request(*key, msg.Json())
	default:
		log.Fatalf("-type '%s' invalid", *messageType)
	}
}
