package transport

import (
	"encoding/json"
	"errors"
	"github.com/qPyth/green-api-task/internal/types"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)

type SettingsProvider interface {
	Settings(idInstance, apiTokenInstance string) (types.Settings, error)
}

type StateInstanceProvider interface {
	StateInstance(idInstance, apiTokenInstance string) (types.StateInstance, error)
}

type MessageSender interface {
	Send(idInstance, apiTokenInstance string, m types.Message) (types.MessageSendResp, error)
}

type FileSender interface {
	SendFile(idInstance, apiTokenInstance string, file types.File) (types.MessageSendResp, error)
}

type Handler struct {
	log           *slog.Logger
	sp            SettingsProvider
	stateProvider StateInstanceProvider
	ms            MessageSender
	fs            FileSender
}

func NewHandler(sp SettingsProvider, StateProvider StateInstanceProvider, ms MessageSender, fs FileSender, l *slog.Logger) *Handler {
	return &Handler{sp: sp, stateProvider: StateProvider, ms: ms, fs: fs, log: l}
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.log.Warn("method not allowed", "method", r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		h.log.Warn("invalid path", "path", r.URL.Path)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("ui/templates/index.html")
	if err != nil {
		h.log.Error("template parsing error", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		h.log.Error("template execute error", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h Handler) Settings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	var i Instances
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		h.log.Error("decode error", "err", err)
	}
	settings, err := h.sp.Settings(strconv.Itoa(i.ID), i.APIToken)
	if err != nil {
		var customErr *types.GreenAPIError
		if errors.As(err, &customErr) {
			newJsonResponse(w, ErrorData{customErr.Error()}, customErr.GetErrorCode())
			return
		}
		newJsonResponse(w, ErrorData{Text: http.StatusText(500)}, 500)
		return
	}

	newJsonResponse(w, settings, http.StatusOK)
}

func (h Handler) StateInstance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var i Instances
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		h.log.Error("decode error", "err", err)
	}
	settings, err := h.stateProvider.StateInstance(strconv.Itoa(i.ID), i.APIToken)
	if err != nil {
		var customErr *types.GreenAPIError
		if errors.As(err, &customErr) {
			newJsonResponse(w, ErrorData{customErr.Error()}, customErr.GetErrorCode())
			return
		}
		newJsonResponse(w, ErrorData{Text: http.StatusText(500)}, 500)
		return
	}

	newJsonResponse(w, settings, http.StatusOK)
}

func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var m MessageSendData

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp, err := h.ms.Send(strconv.Itoa(m.ID), m.APIToken, types.NewMessage(m.PhoneNumber, m.Message))
	if err != nil {
		var customErr *types.GreenAPIError
		if errors.As(err, &customErr) {
			newJsonResponse(w, ErrorData{customErr.Error()}, customErr.GetErrorCode())
			return
		}
		newJsonResponse(w, ErrorData{Text: http.StatusText(500)}, 500)
		return
	}

	newJsonResponse(w, resp, http.StatusOK)
}

func (h Handler) SendFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var f FileSendData

	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp, err := h.fs.SendFile(strconv.Itoa(f.ID), f.APIToken, types.NewFile(f.PhoneNumber, f.FileUrl))
	if err != nil {
		var customErr *types.GreenAPIError
		if errors.As(err, &customErr) {
			newJsonResponse(w, ErrorData{customErr.Error()}, customErr.GetErrorCode())
			return
		}
		newJsonResponse(w, ErrorData{Text: http.StatusText(500)}, 500)
		return
	}

	newJsonResponse(w, resp, http.StatusOK)
}

func newJsonResponse(w http.ResponseWriter, data any, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}
