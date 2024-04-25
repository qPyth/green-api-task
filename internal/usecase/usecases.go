package usecases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/qPyth/green-api-task/internal/types"
	"io"
	"net/http"
	"net/url"
	"time"
)

var hostUrl = "https://api.green-api.com"

var (
	settingsMethod    = "getSettings"
	stateInstance     = "getStateInstance"
	sendMessageMethod = "sendMessage"
	sendFile          = "sendFileByUrl"
)

type GreenApiUC struct {
}

func NewGreenApiUC() *GreenApiUC {
	return new(GreenApiUC)
}

func (g GreenApiUC) Settings(idInstance, apiTokenInstance string) (types.Settings, error) {
	var s types.Settings

	uri, err := url.JoinPath(hostUrl, fmt.Sprintf("waInstance%s", idInstance), settingsMethod, apiTokenInstance)
	if err != nil {
		return s, fmt.Errorf("url join error")
	}

	data, err := request(uri, http.MethodGet, nil)
	if err != nil {
		return s, fmt.Errorf("fetching error: %w", err)
	}
	defer data.Close()

	err = json.NewDecoder(data).Decode(&s)
	if err != nil {
		return s, fmt.Errorf("decode error: %w", err)
	}

	return s, nil
}

func (g GreenApiUC) StateInstance(idInstance, apiTokenInstance string) (types.StateInstance, error) {
	var s types.StateInstance

	uri, err := url.JoinPath(hostUrl, fmt.Sprintf("waInstance%s", idInstance), stateInstance, apiTokenInstance)
	if err != nil {
		return s, fmt.Errorf("url join error")
	}

	data, err := request(uri, http.MethodGet, nil)
	if err != nil {
		return s, fmt.Errorf("fetching error: %w", err)
	}
	defer data.Close()

	err = json.NewDecoder(data).Decode(&s)
	if err != nil {
		return s, fmt.Errorf("decode error: %w", err)
	}

	return s, nil
}

func (g GreenApiUC) Send(idInstance, apiTokenInstance string, msg types.Message) (types.MessageSendResp, error) {
	var m types.MessageSendResp

	uri, err := url.JoinPath(hostUrl, fmt.Sprintf("waInstance%s", idInstance), sendMessageMethod, apiTokenInstance)
	if err != nil {
		return m, fmt.Errorf("url join error")
	}

	dataInByte, err := json.Marshal(&msg)

	data, err := request(uri, http.MethodPost, bytes.NewReader(dataInByte))
	if err != nil {
		return m, fmt.Errorf("data fetching error: %w", err)
	}
	defer data.Close()

	err = json.NewDecoder(data).Decode(&m)
	if err != nil {
		return m, fmt.Errorf("json decode error: %w", err)
	}

	return m, nil
}

func (g GreenApiUC) SendFile(idInstance, apiTokenInstance string, file types.File) (types.MessageSendResp, error) {
	var m types.MessageSendResp

	uri, err := url.JoinPath(hostUrl, fmt.Sprintf("waInstance%s", idInstance), sendFile, apiTokenInstance)
	if err != nil {
		return m, fmt.Errorf("url join error")
	}

	dataInByte, err := json.Marshal(&file)

	data, err := request(uri, http.MethodPost, bytes.NewReader(dataInByte))
	if err != nil {
		return m, fmt.Errorf("data fetching error: %w", err)
	}
	defer data.Close()

	err = json.NewDecoder(data).Decode(&m)
	if err != nil {
		return m, fmt.Errorf("json decode error: %w", err)
	}

	return m, nil
}

func request(uri string, method string, r io.Reader) (io.ReadCloser, error) {

	req, err := http.NewRequest(method, uri, r)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//data, _ := io.ReadAll(resp.Body)
		//fmt.Println(string(data))
		err = types.NewGreenAPIError(fmt.Sprintf("non-200 status code from API: %d (%s)", resp.StatusCode, resp.Status), resp.StatusCode)
		return nil, err
	}

	return resp.Body, nil
}
