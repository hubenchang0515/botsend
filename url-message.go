package botsend

import "encoding/json"

type UrlContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Cover       string `json:"picurl"`
}

type UrlNews struct {
	Articles []UrlContent `json:"articles"`
}

type UrlMessage struct {
	Type string  `json:"msgtype"`
	News UrlNews `json:"news"`
}

func NewUrlMessage() *UrlMessage {
	return &UrlMessage{
		Type: "news",
		News: UrlNews{
			Articles: make([]UrlContent, 1),
		},
	}
}

func (m *UrlMessage) Json() []byte {
	data, _ := json.Marshal(m)
	return data
}

func (m *UrlMessage) SetUrl(url string) {
	m.News.Articles[0].Url = url
}

func (m *UrlMessage) SetTitle(title string) {
	m.News.Articles[0].Title = title
}

func (m *UrlMessage) SetDescription(description string) {
	m.News.Articles[0].Description = description
}

func (m *UrlMessage) SetCover(url string) {
	m.News.Articles[0].Cover = url
}
