package goreplicate

import (
	"time"
)

// Client is the main struct for interacting with the API
type Client struct {
	Authorization string
	Model         *Model
	Request       *Request
	Response      *Response
}

// Request is the prediction request that is sent to the API
type Request struct {
	Version string `json:"version"`
	Input   any    `json:"input"`
}

// Response is the prediction response that is returned from the API
type Response struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Urls    struct {
		Get    string `json:"get"`
		Cancel string `json:"cancel"`
	} `json:"urls"`
	CreatedAt   time.Time   `json:"created_at"`
	CompletedAt time.Time   `json:"completed_at"`
	Status      interface{} `json:"status"`
	Input       any         `json:"input"`
	Output      any         `json:"output"`
	Error       string      `json:"error"`
	Logs        string      `json:"logs"`
	Metrics     struct {
		PredictTime float64 `json:"predict_time"`
	} `json:"metrics"`
}

// Model is the model that represents the model to be used for prediction
type Model struct {
	Owner   string                 `json:"owner"`
	Name    string                 `json:"name"`
	Version string                 `json:"version"`
	Input   map[string]interface{} `json:"input"`
}

// CreateRequest creates a new request to send to the API
func (m *Model) CreateRequest() *Request {
	return &Request{
		Version: m.Version,
		Input:   m.Input,
	}
}

// NewModel creates a new model to be used for prediction
func NewModel(owner, name, version string) (model *Model) {
	model = &Model{
		Owner:   owner,
		Name:    name,
		Version: version,
	}
	// TODO: think of a better way to return models
	if owner == "stability-ai" && name == "stable-diffusion" {
		model.Input = map[string]interface{}{
			"prompt":              "",
			"width":               512,
			"height":              512,
			"prompt_strength":     0.8,
			"num_outputs":         1,
			"num_inference_steps": 100,
			"guidance_scale":      7.5,
		}
		return
	}

	if owner == "openai" && name == "whisper" {
		model.Input = map[string]interface{}{
			"audio":           "",
			"model":           "base",
			"transcription":   "plain text",
			"translate":       false,
			"temperature":     0.5,
			"patience":        1.0,
			"suppress_tokens": -1,
			"initial_prompt":  "",
		}
		return
	}
	model.Input = map[string]interface{}{}
	return
}
