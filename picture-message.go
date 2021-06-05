package botsend

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type PictureContent struct {
	Data string `json:"base64"`
	Md5  string `json:"md5"`
}

type PictureMessage struct {
	Type    string         `json:"msgtype"`
	Picture PictureContent `json:"image"`
}

func NewPictureMessage() *PictureMessage {
	return &PictureMessage{
		Type:    "image",
		Picture: PictureContent{},
	}
}

func (m *PictureMessage) Json() []byte {
	data, _ := json.Marshal(m)
	return data
}

func (m *PictureMessage) SetPicture(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	base64Data := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(base64Data, data)
	m.Picture.Data = string(base64Data)
	m.Picture.Md5 = fmt.Sprintf("%x", md5.Sum(data))
}
