package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
)

type PictureContent struct {
	Data []byte   `json:"base64"`
	Md5  [16]byte `json:"md5"`
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
	m.Picture.Data = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(m.Picture.Data, data)
	m.Picture.Md5 = md5.Sum(data)
}
