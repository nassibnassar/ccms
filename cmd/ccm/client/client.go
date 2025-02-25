package client

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

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/api"
	"github.com/indexdata/ccms/internal/eout"
)

func Send(cmd string) error {
	var rq = &api.CommandRequest{Commands: []string{cmd}}
	// send the request
	var httprs *http.Response
	var err error
	if httprs, err = SendRequest("POST", "/cmd", rq); err != nil {
		return err
	}
	// check for error response
	if httprs.StatusCode != http.StatusCreated {
		var m string
		if m, err = ReadResponseMessage(httprs); err != nil {
			return err
		}
		fmt.Println(m)
		return errors.New(m)
	}

	var cmdr api.CommandResponse
	if err = ReadResponse(httprs, &cmdr); err != nil {
		return err
	}
	for _, a := range cmdr.Fields {
		fmt.Printf("%s\t", a.Name)
	}
	fmt.Printf("\n")
	for _, a := range cmdr.Data {
		for _, b := range a.Values {
			fmt.Printf("%s\t", b)
		}
		fmt.Printf("\n")
	}
	// fmt.Printf("%#v\n", cmdr)
	// print confirmation
	// eout.Info("enabled: %s", rq.Command)
	return nil
}

func SendRequest(method, url string, requestStruct interface{}) (*http.Response, error) {
	// WarnNoTLS(opt.NoTLS)
	host := "127.0.0.1"
	var rqj []byte
	var err error
	if rqj, err = json.Marshal(requestStruct); err != nil {
		return nil, err
	}

	var conn *tls.Conn
	var transport *http.Transport
	if host == "" {
		transport = &http.Transport{}
	} else {
		var tlsConfig = http.DefaultTransport.(*http.Transport).TLSClientConfig
		var tlsClientConfig *tls.Config
		// if opt.TLSSkipVerify {
		// 	tlsClientConfig = &tls.Config{InsecureSkipVerify: true}
		// }
		transport = &http.Transport{
			TLSClientConfig: tlsClientConfig,
			// func(ctx context.Context, network, addr string) (net.Conn, error)
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err = tls.Dial(network, addr, tlsConfig)
				return conn, err
			},
		}
	}
	var client = &http.Client{Transport: transport}
	var remote string
	adminPort := ccms.DefaultPort
	if host == "" {
		remote = "http://127.0.0.1:" + adminPort
	} else {
		// if opt.NoTLS {
		remote = "http://" + host + ":" + adminPort
		// } else {
		// 	remote = "https://" + host + ":" + adminPort
		// }
	}
	var httprq *http.Request
	if httprq, err = http.NewRequest(method, remote+url, bytes.NewBuffer(rqj)); err != nil {
		return nil, err
	}
	httprq.SetBasicAuth("admin", "admin")
	httprq.Header.Set("Content-Type", "application/json")
	var hrs *http.Response
	if hrs, err = client.Do(httprq); err != nil {
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

func ReadResponse(httpResponse *http.Response, responseStruct interface{}) error {
	var body []byte
	var err error
	if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return err
	}
	fmt.Printf("%s\n", body)
	if err = json.Unmarshal(body, responseStruct); err != nil {
		return err
	}
	return nil
}

func ReadResponseMessage(httpResponse *http.Response) (string, error) {
	var m map[string]interface{}
	var err error
	if m, err = ReadResponseMap(httpResponse); err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", m["message"]), nil
}

func ReadResponseMap(httpResponse *http.Response) (map[string]interface{}, error) {
	var m map[string]interface{}
	var err error
	if err = json.NewDecoder(httpResponse.Body).Decode(&m); err != nil {
		return nil, fmt.Errorf("decoding server response: %s", err)
	}
	return m, nil
}

func WarnNoTLS(noTLS bool) {
	if noTLS {
		eout.Warning("TLS disabled in connection to server")
	}
}
