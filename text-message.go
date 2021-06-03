package main

import (
	"encoding/json"
	"strings"
)

type TextContent struct {
	Content string   `json:"content"`
	Mention []string `json:"mentioned_mobile_list"`
}

type TextMessage struct {
	Type string      `json:"msgtype"`
	Text TextContent `json:"text"`
}

func NewTextMessgae() *TextMessage {
	return &TextMessage{
		Type: "text",
		Text: TextContent{},
	}
}

func (m *TextMessage) Json() []byte {
	data, _ := json.Marshal(m)
	return data
}

func (m *TextMessage) SetContent(content string) {
	m.Text.Content = content
}

func (m *TextMessage) SetMention(mention []string) {
	m.Text.Mention = mention
}

func (m *TextMessage) SetMentionString(mention string) {
	m.SetMention(strings.Split(mention, ","))
}
