package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/cmd/ccd/option"
	"github.com/indexdata/ccms/cmd/ccd/osutil"
	"github.com/indexdata/ccms/cmd/ccd/parser"
	"github.com/indexdata/ccms/cmd/ccd/process"
	"github.com/indexdata/ccms/internal/eout"
	"github.com/indexdata/ccms/internal/global"
	"github.com/indexdata/ccms/internal/protocol"
)

type server struct {
	Opt *option.Server
}

func Start(opt *option.Server) error {
	var err error
	// Require datadir specified
	if opt.Datadir == "" {
		return fmt.Errorf("data directory not specified")
	}
	// Require datadir exists
	var exists bool
	if exists, err = osutil.FileExists(opt.Datadir); err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("data directory not found: %s", opt.Datadir)
	}

	// Check if server is already running.
	running, pid, err := process.IsServerRunning(opt.Datadir)
	if err != nil {
		return err
	}
	if running {
		eout.Error("lock file %q already exists and server (PID %d) appears to be running", osutil.SystemPIDFileName(opt.Datadir), pid)
		return fmt.Errorf("could not start server")
	}
	// Write lock file for new server instance.
	if err = process.WritePIDFile(opt.Datadir); err != nil {
		return err
	}
	defer process.RemovePIDFile(opt.Datadir)

	if err = startServer(opt); err != nil {
		return fmt.Errorf("server stopped: %s", err)
	}
	return nil
}

func startServer(opt *option.Server) error {
	var sigc = make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM)
	go func() {
		<-sigc
		process.SetStop()
	}()

	go serve(opt)
	//go harvest.Harvest()

	for {
		if process.Stop() {
			break
		}
		time.Sleep(5 * time.Second)
	}
	log.Info("server stopped")
	return nil
}

func serve(opt *option.Server) {
	var listen string
	if opt.Listen == "" {
		listen = "127.0.0.1"
	} else {
		listen = opt.Listen
	}
	addr := net.JoinHostPort(listen, opt.Port)
	httpsvr := http.Server{
		Addr:    addr,
		Handler: setupHandlers(&server{Opt: opt}),
	}
	log.Info("CCMS %s, listening on %s", global.Version, addr)
	if opt.NoTLS && opt.Listen != "" {
		log.Warning("disabling TLS (insecure)")
	}
	var err error
	if opt.Listen == "" || opt.NoTLS {
		err = httpsvr.ListenAndServe()
	} else {
		err = httpsvr.ListenAndServeTLS(opt.TLSCert, opt.TLSKey)
	}
	if err != nil {
		m := fmt.Sprintf("error starting server: %v", err)
		log.Fatal("%s", m)
		eout.Error("%s", m)
		os.Exit(1)
	}
}

func setupHandlers(svr *server) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/cmd", svr.handleCommand)
	return mux
}

func (svr *server) handleCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// log.Info("request: %s", requestString(r))
		svr.handleCommandPost(w, r)
		return
	}
	// var m = unsupportedMethod("/config", r)
	// log.Info(m)
	// http.Error(w, m, http.StatusMethodNotAllowed)
}

func (svr *server) handleCommandPost(w http.ResponseWriter, r *http.Request) {
	// read request
	var rq protocol.CommandRequest
	var ok bool
	if ok = ReadRequest(w, r, &rq); !ok {
		return
	}
	node, err, pass := parser.Parse(rq.Command)
	//fmt.Printf("### %#v --- %v\n", node, err)
	if err != nil {
		returnError(w, err.Error(), http.StatusOK /* http.StatusBadRequest */)
		return
	}
	//if node == nil {
	//        returnError(w, errors.New("syntax error"), http.StatusBadRequest)
	//        return
	//}
	//log.Info("parsed: %#v", node)
	_ = pass
	var cmdr protocol.CommandResponse
	switch n := node.(type) {
	case *ast.HelpStmt:
		cmdr = protocol.CommandResponse{
			Status: "help",
			Fields: []protocol.FieldDescription{
				{
					Name: "command",
				},
				{
					Name: "description",
				},
			},
			Data: []protocol.DataRow{
				{
					Values: []string{"show filters", "list existing filters"},
				},
				{
					Values: []string{"show sets", "list existing sets"},
				},
			},
			//Message: "create set <set_name>\tdefine a new set\n" +
			//        "show sets\t\tlist existing sets",
		}
	case *ast.PingStmt:
		cmdr = protocol.CommandResponse{Status: "ping"}
	case *ast.ShowStmt:
		switch n.Name {
		case "filters":
			cmdr = protocol.CommandResponse{
				Status: "show",
				Fields: []protocol.FieldDescription{
					{
						Name: "filter",
					},
				},
				Data: []protocol.DataRow{},
			}
		case "sets":
			cmdr = protocol.CommandResponse{
				Status: "show",
				Fields: []protocol.FieldDescription{
					{
						Name: "set",
					},
				},
				Data: []protocol.DataRow{
					{
						Values: []string{"reserve"},
					},
				},
			}
		default:
			cmdr = protocol.CommandResponse{
				Status:  "error",
				Message: fmt.Sprintf("unknown type %q", n.Name),
			}
		}
	default:
		cmdr = protocol.CommandResponse{
			Status: "error",
			Message: fmt.Sprintf("syntax error near %q\n%s\n^",
				strings.Fields(rq.Command)[0],
				rq.Command),
		}
	}
	// success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(cmdr); err != nil {
		// TODO error handling
		_ = err
	}
}

func returnError(w http.ResponseWriter, errString string, statusCode int) {
	HTTPError(w, errString, statusCode)
}

func ReadRequest(w http.ResponseWriter, r *http.Request, requestStruct interface{}) bool {
	// Authenticate user.
	var user string
	var ok bool
	if user, ok = HandleBasicAuth(w, r); !ok {
		return false
	}
	_ = user
	// Read the json request.
	var body []byte
	var err error
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return false
	}
	if err = json.Unmarshal(body, requestStruct); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return false
	}
	log.Trace("request %s %v\n", r.RemoteAddr, requestStruct)
	return true
}

func HandleBasicAuth(w http.ResponseWriter, r *http.Request) (string, bool) {
	host := osutil.AddrHost(r.RemoteAddr)
	var user, password string
	var ok bool
	user, password, ok = r.BasicAuth()
	if !ok {
		e := "invalid HTTP basic authentication"
		log.Info("%s from %s", e, host)
		//http.Error(w, e, http.StatusForbidden)
		returnError(w, e, http.StatusForbidden)
		return user, false
	}
	if user != "nemo" || password != "testpass" {
		e := fmt.Sprintf("authentication failed")
		//
		// TODO include user name as follows, once we support
		//      multiple users:
		// e := fmt.Sprintf("authentication failed for user %q", user)
		//
		log.Info("%s from %s", e, host)
		//http.Error(w, e, http.StatusForbidden)
		returnError(w, e, http.StatusForbidden)
		return user, false
	}
	//var match bool
	//var err error
	//match, err = srv.storage.Authenticate(user, password)
	//if err != nil {
	//        var m = "Unauthorized (user '" + user + "')"
	//        log.Println(m + ": " + err.Error())
	//        //w.Header().Set("WWW-Authenticate", "Basic")
	//        http.Error(w, m, http.StatusForbidden)
	//        return user, false
	//}
	/*	if !match {
			var m = "Unauthorized (user '" + user + "'): " + "Unable to authenticate username/password"
			log.Info(m)
			//w.Header().Set("WWW-Authenticate", "Basic")
			http.Error(w, m, http.StatusForbidden)
			return user, false
		}
	*/
	return user, true
}

/*
func ReadRequest(w http.ResponseWriter, r *http.Request, requestStruct interface{}) bool {
	var body []byte
	var err error
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return false
	}
	//log.Info("received: %s\n", body)
	if err = json.Unmarshal(body, requestStruct); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return false
	}
	log.Trace("request %s %v\n", r.RemoteAddr, requestStruct)
	return true
}
*/

func HandleError(w http.ResponseWriter, err error, statusCode int) {
	log.Error("%s", err)
	HTTPError(w, err.Error(), statusCode)
}

func HTTPError(w http.ResponseWriter, errString string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	m := &protocol.CommandResponse{
		Status:  "error",
		Message: errString,
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		// TODO error handling
		panic(err)
	}
}

func requestString(r *http.Request) string {
	var remoteHost, remotePort string
	remoteHost, remotePort, _ = net.SplitHostPort(r.RemoteAddr)
	return fmt.Sprintf("host=%s port=%s method=%s uri=%s", remoteHost, remotePort, r.Method, r.URL)
}
