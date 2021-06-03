package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type MarkdownContent struct {
	Content string `json:"content"`
}

type MarkdownMessage struct {
	Type     string          `json:"msgtype"`
	Markdown MarkdownContent `json:"markdown"`
}

func NewMarkdownMessage() *MarkdownMessage {
	return &MarkdownMessage{
		Type:     "markdown",
		Markdown: MarkdownContent{},
	}
}

func (m *MarkdownMessage) Json() []byte {
	data, _ := json.Marshal(m)
	return data
}

func (m *MarkdownMessage) SetString(md string) {
	m.Markdown.Content = md
}

func (m *MarkdownMessage) SetFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	m.Markdown.Content = string(data)
}
