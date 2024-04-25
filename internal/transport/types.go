package transport

type Instances struct {
	ID       int    `json:"idInstance"`
	APIToken string `json:"apiTokenInstance"`
}

type ErrorData struct {
	Text string `json:"error"`
}

type MessageSendData struct {
	Instances
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
}

type FileSendData struct {
	Instances
	PhoneNumber string `json:"phoneNumber"`
	FileUrl     string `json:"fileUrl"`
}
