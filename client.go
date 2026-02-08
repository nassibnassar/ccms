package ccms

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/indexdata/ccms/internal/eout"
	"github.com/indexdata/ccms/internal/protocol"
)

// client that connects to a CCMS server
type Client struct {
	Host          string // server host name
	Port          string // server port
	User          string // user name for authentication
	Password      string // user password
	NoTLS         bool   // disable TLS (insecure)
	TLSSkipVerify bool   // do not verify server certificate chain and host name (insecure)
}

// response from a CCMS server
type Response struct {
	Status  string             `json:"status"`            // status of command, or "error"
	Message string             `json:"message,omitempty"` // error message
	Fields  []FieldDescription `json:"fields,omitempty"`  // attribute metadata for query result
	Data    []DataRow          `json:"data,omitempty"`    // query result data
}

// metadata for an attribute
type FieldDescription struct {
	Name string `json:"name"` // attribute name
	Type string `json:"type"` // data type
}

// a row of data
type DataRow struct {
	Values []any `json:"values"` // data values
}

// send a command to the server and return the response
func (c *Client) Send(cmd string) (*Response, error) {
	var rq = &protocol.CommandRequest{Command: cmd}
	// send the request
	var httprs *http.Response
	var err error
	if httprs, err = sendRequest(c, "POST", "/cmd", rq); err != nil {
		return nil, err
	}
	// check for error response
	if httprs.StatusCode != http.StatusOK {
		var m string
		if m, err = readResponseMessage(httprs); err != nil {
			return nil, err
		}
		return nil, errors.New(m)
	}

	var cmdr protocol.CommandResponse
	if err = readResponse(httprs, &cmdr); err != nil {
		return nil, err
	}

	if cmdr.Status == "error" {
		resp := &Response{Status: cmdr.Status, Message: cmdr.Message}
		return resp, nil
	}

	if cmdr.Status == "ping" {
		return &Response{Status: cmdr.Status}, nil
	}

	fields := make([]FieldDescription, 0)
	for i := range cmdr.Fields {
		fd := FieldDescription{Name: cmdr.Fields[i].Name, Type: cmdr.Fields[i].Type}
		fields = append(fields, fd)
	}
	data := make([]DataRow, 0)
	for i := range cmdr.Data {
		values := make([]any, 0)
		for j := range cmdr.Data[i].Values {
			values = append(values, cmdr.Data[i].Values[j])
		}
		dr := DataRow{Values: values}
		data = append(data, dr)
	}
	resp := &Response{
		Status:  cmdr.Status,
		Fields:  fields,
		Data:    data,
		Message: cmdr.Message,
	}
	// fmt.Printf("%#v\n", cmdr)
	// print confirmation
	// eout.Info("enabled: %s", rq.Command)
	return resp, nil
}

func sendRequest(client *Client, method, url string, requestStruct interface{}) (*http.Response, error) {
	var rqj []byte
	var err error
	if rqj, err = json.Marshal(requestStruct); err != nil {
		return nil, err
	}

	var conn *tls.Conn
	var transport *http.Transport
	if client.Host == "" {
		transport = &http.Transport{}
	} else {
		var tlsConfig = http.DefaultTransport.(*http.Transport).TLSClientConfig
		var tlsClientConfig *tls.Config
		if client.TLSSkipVerify {
			tlsClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
		transport = &http.Transport{
			TLSClientConfig: tlsClientConfig,
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err = tls.Dial(network, addr, tlsConfig)
				return conn, err
			},
		}
	}
	var httpClient = &http.Client{Transport: transport}
	var remote string
	var s string
	if client.Host != "127.0.0.1" && !client.NoTLS {
		s = "s"
	}
	remote = "http" + s + "://" + client.Host + ":" + client.Port
	var httprq *http.Request
	if httprq, err = http.NewRequest(method, remote+url, bytes.NewBuffer(rqj)); err != nil {
		return nil, err
	}
	httprq.SetBasicAuth(client.User, client.Password)
	httprq.Header.Set("Content-Type", "application/json")
	var hrs *http.Response
	if hrs, err = httpClient.Do(httprq); err != nil {
		return nil, err
	}
	if conn != nil {
		// verbose output
		var v uint16 = conn.ConnectionState().Version
		eout.Trace("protocol version: %d,%d", (v>>8)&255, v&255)
		var s string
		switch v {
		case 0x0300:
			s = "SSL (deprecated)"
		case 0x0301:
			s = "TLS 1.0 (deprecated)"
		case 0x0302:
			s = "TLS 1.1 (deprecated)"
		case 0x0303:
			s = "TLS 1.2"
		case 0x0304:
			s = "TLS 1.3"
		default:
			s = fmt.Sprintf("unknown version: { %d, %d }", (v>>8)&255, v&255)
		}
		eout.Verbose("TLS/SSL protocol: %s", s)
	} else {
		eout.Verbose("no TLS/SSL protocol")
	}
	return hrs, nil
}

func readResponse(httpResponse *http.Response, responseStruct interface{}) error {
	var body []byte
	var err error
	if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(body, responseStruct); err != nil {
		return err
	}
	return nil
}

func readResponseMessage(httpResponse *http.Response) (string, error) {
	var m map[string]interface{}
	var err error
	if m, err = readResponseMap(httpResponse); err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", m["message"]), nil
}

func readResponseMap(httpResponse *http.Response) (map[string]interface{}, error) {
	var m map[string]interface{}
	var err error
	if err = json.NewDecoder(httpResponse.Body).Decode(&m); err != nil {
		return nil, fmt.Errorf("decoding server response: %s", err)
	}
	return m, nil
}
