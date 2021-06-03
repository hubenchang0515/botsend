package main

import "encoding/json"

type FileContent struct {
	Id string `json:"media_id"`
}

type FileMessage struct {
	Type string      `json:"msgtype"`
	File FileContent `json:"file"`
}

func NewFileMessage() *FileMessage {
	return &FileMessage{
		Type: "file",
		File: FileContent{},
	}
}

func (m *FileMessage) Json() []byte {
	data, _ := json.Marshal(m)
	return data
}

func (m *FileMessage) SetFile(path string) {
	id := Upload(Key(), path)
	m.File.Id = id
}
