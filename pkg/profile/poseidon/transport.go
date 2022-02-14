package poseidon

import (
	common "ArchiMoebius/mythic_c2_websocket/pkg/profile/common"
	"encoding/json"
	"fmt"
)

type Transport struct {
	common.BaseTransportConfig
	BindAddress  string `json:"bindaddress"`
	UseSSL       bool   `json:"usessl"`
	SSLKey       string `json:"sslkey"`
	SSLCert      string `json:"sslcert"`
	WebSocketURI string `json:"websocketuri"`
	DefaultPage  string `json:"defaultpage"`
	LogFile      string `json:"logfile"`
	Debug        bool   `json:"debug"`
}

func (d *Transport) UseMTLS() bool {
	return false
}

func (d *Transport) GetLogPath(string) string {
	return fmt.Sprintf("/Mythic/cs_code/%s", d.LogFile)
}

func (d *Transport) GetHTTPFilename() string {
	return d.DefaultPage
}

func (d *Transport) GetWebSocketFilename() string {
	return d.WebSocketURI
}

func (d *Transport) GetServerAddress() string {
	return d.BindAddress
}

func (d *Transport) ParseClientMessage(blob []byte) ([]byte, error) {
	// TODO: Unsure if custom parsing is required...
	return blob, nil
}

func (d *Transport) Load() error {
	return nil
}

// MarshalString hack to return data as a json string
func (d *Transport) MarshalString() (string, error) {
	data, err := json.Marshal(d)

	if err != nil {
		return "", err
	}

	return string(data[:]), nil
}
