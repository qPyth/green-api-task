package types

import "path/filepath"

type Settings struct {
	Wid                               string `json:"wid"`
	CountryInstance                   string `json:"countryInstance"`
	TypeAccount                       string `json:"typeAccount"`
	WebhookUrl                        string `json:"webhookUrl"`
	WebhookUrlToken                   string `json:"webhookUrlToken"`
	DelaySendMessagesMilliseconds     int    `json:"delaySendMessagesMilliseconds"`
	MarkIncomingMessagesReaded        string `json:"markIncomingMessagesReaded"`
	MarkIncomingMessagesReadedOnReply string `json:"markIncomingMessagesReadedOnReply"`
	SharedSession                     string `json:"sharedSession"`
	OutgoingWebhook                   string `json:"outgoingWebhook"`
	OutgoingMessageWebhook            string `json:"outgoingMessageWebhook"`
	OutgoingAPIMessageWebhook         string `json:"outgoingAPIMessageWebhook"`
	IncomingWebhook                   string `json:"incomingWebhook"`
	DeviceWebhook                     string `json:"deviceWebhook"`
	StatusInstanceWebhook             string `json:"statusInstanceWebhook"`
	StateWebhook                      string `json:"stateWebhook"`
	EnableMessagesHistory             string `json:"enableMessagesHistory"`
	KeepOnlineStatus                  string `json:"keepOnlineStatus"`
	PollMessageWebhook                string `json:"pollMessageWebhook"`
	IncomingBlockWebhook              string `json:"incomingBlockWebhook"`
	IncomingCallWebhook               string `json:"incomingCallWebhook"`
}

type StateInstance struct {
	StateInstance string `json:"stateInstance"`
}

type MessageSendResp struct {
	IdMessage string `json:"idMessage"`
}

type Message struct {
	ChatId  string `json:"chatId"`
	Message string `json:"message"`
}

func NewMessage(chatID, message string) Message {
	return Message{
		ChatId:  chatID + "@c.us",
		Message: message,
	}
}

type File struct {
	ChatId   string `json:"chatId"`
	UrlFile  string `json:"urlFile"`
	FileName string `json:"fileName"`
}

func NewFile(chatID string, urlFile string) File {
	return File{
		ChatId:   chatID + "@c.us",
		UrlFile:  urlFile,
		FileName: filepath.Base(urlFile),
	}
}
