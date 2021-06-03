package main

import (
	"flag"
	"log"
)

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
		msg := NewTextMessgae()
		msg.SetContent(flag.Arg(0))
		if *at != "" {
			msg.SetMentionString(*at)
		}
		// Request(GetHook(*key), msg.Json())
	case "picture":
		msg := NewPictureMessage()
		msg.SetPicture(flag.Arg(0))
		Request(GetHook(*key), msg.Json())
	case "markdown":
		msg := NewMarkdownMessage()
		if *markdown == "" {
			msg.SetFile(flag.Arg(0))
		} else {
			msg.SetString(*markdown)
		}
		Request(GetHook(*key), msg.Json())
	case "url":
		msg := NewUrlMessage()
		msg.SetUrl(*url)
		msg.SetTitle(*title)
		msg.SetDescription(*description)
		msg.SetCover(*cover)
		Request(GetHook(*key), msg.Json())
	case "file":
		msg := NewFileMessage()
		msg.SetFile(flag.Arg(0))
		Request(GetHook(*key), msg.Json())
	default:
		log.Fatalf("-type '%s' invalid", *messageType)
	}
}
