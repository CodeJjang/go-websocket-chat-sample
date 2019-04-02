package models

type Message struct {
	Sender string `json:sender,omitempty`
	Receiver string `json:receiver,omitempty`
	Content string `json:content,omitempty`
}