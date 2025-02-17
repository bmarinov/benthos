package input

import (
	ihttpdocs "github.com/benthosdev/benthos/v4/internal/http/docs"
)

// StreamConfig contains fields for specifying consumption behaviour when the
// body of a request is a constant stream of bytes.
type StreamConfig struct {
	Enabled   bool   `json:"enabled" yaml:"enabled"`
	Reconnect bool   `json:"reconnect" yaml:"reconnect"`
	Codec     string `json:"codec" yaml:"codec"`
	MaxBuffer int    `json:"max_buffer" yaml:"max_buffer"`
}

// HTTPClientConfig contains configuration for the HTTPClient output type.
type HTTPClientConfig struct {
	ihttpdocs.Config `json:",inline" yaml:",inline"`
	Payload          string       `json:"payload" yaml:"payload"`
	DropEmptyBodies  bool         `json:"drop_empty_bodies" yaml:"drop_empty_bodies"`
	Stream           StreamConfig `json:"stream" yaml:"stream"`
}

// NewHTTPClientConfig creates a new HTTPClientConfig with default values.
func NewHTTPClientConfig() HTTPClientConfig {
	cConf := ihttpdocs.NewConfig()
	cConf.Verb = "GET"
	cConf.URL = ""
	return HTTPClientConfig{
		Config:          cConf,
		Payload:         "",
		DropEmptyBodies: true,
		Stream: StreamConfig{
			Enabled:   false,
			Reconnect: true,
			Codec:     "lines",
			MaxBuffer: 1000000,
		},
	}
}
