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

// response from CCMS server
type Response struct {
	Results []*Result `json:"results"` // result for each command
}

// result of a command
type Result struct {
	Status  string              `json:"status"`            // status of command, or "error"
	Message string              `json:"message,omitempty"` // error message
	Fields  []*FieldDescription `json:"fields,omitempty"`  // attribute metadata for query result
	Data    []*DataRow          `json:"data,omitempty"`    // query result data
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

// send one or more commands to the server and return the response
func (c *Client) Send(cmd string) (*Response, error) {
	var rq = &protocol.Request{Commands: cmd}
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

	var resp Response
	if err = readResponse(httprs, &resp); err != nil {
		return nil, err
	}
	return &resp, nil

	/*
		results := make([]Result, 0)
		for j := range cmdr.Results {
			r := cmdr.Results[j]
			//if r.Status == "error" {
			//        resp := &Result{Status: r.Status, Message: r.Message}
			//        return resp, nil
			//}

			//if r.Status == "ping" {
			//        return &Result{Status: r.Status}, nil
			//}

			fields := make([]FieldDescription, 0)
			for i := range r.Fields {
				fd := FieldDescription{Name: r.Fields[i].Name, Type: r.Fields[i].Type}
				fields = append(fields, fd)
			}
			data := make([]DataRow, 0)
			for i := range r.Data {
				values := make([]any, 0)
				for j := range r.Data[i].Values {
					values = append(values, r.Data[i].Values[j])
				}
				dr := DataRow{Values: values}
				data = append(data, dr)
			}
			results = append(results, Result{
				Status:  r.Status,
				Fields:  fields,
				Data:    data,
				Message: r.Message,
			})
		}
		resp := &Response{Results: results}
		// fmt.Printf("%#v\n", cmdr)
		// print confirmation
		// eout.Info("enabled: %s", rq.Command)
		return resp, nil
	*/
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

func readResponse(httpResponse *http.Response, responseStruct any) error {
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
