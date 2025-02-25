package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/api"
	"github.com/indexdata/ccms/cmd/ccms/log"
	"github.com/indexdata/ccms/cmd/ccms/option"
	"github.com/indexdata/ccms/cmd/ccms/osutil"
	"github.com/indexdata/ccms/cmd/ccms/parser"
	"github.com/indexdata/ccms/cmd/ccms/process"
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
		log.Fatal("lock file %q already exists and server (PID %d) appears to be running", osutil.SystemPIDFileName(opt.Datadir), pid)
		return fmt.Errorf("could not start server")
	}
	// Write lock file for new server instance.
	if err = process.WritePIDFile(opt.Datadir); err != nil {
		return err
	}
	defer process.RemovePIDFile(opt.Datadir)

	log.Info("starting ccms %s", ccms.Version)
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
		log.Debug("received shutdown request")
		log.Info("shutting down")
		process.SetStop()
	}()

	log.Info("listening on address \"%s\", port %s", opt.Listen, opt.Port)

	go serve(opt)

	for {
		if process.Stop() {
			break
		}
		time.Sleep(5 * time.Second)
	}

	return nil
}

func serve(opt *option.Server) {
	httpsvr := http.Server{
		Addr:    net.JoinHostPort(opt.Listen, opt.Port),
		Handler: setupHandlers(&server{Opt: opt}),
	}
	err := httpsvr.ListenAndServe()
	if err != nil {
		log.Error("%v", err)
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
	var rq api.CommandRequest
	var ok bool
	if ok = ReadRequest(w, r, &rq); !ok {
		return
	}
	node, err, pass := parser.Parse(rq.Commands[0])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", node)
	_ = pass
	// TODO run command
	cmdr := api.CommandResponse{
		Fields: []api.FieldDescription{
			{
				Name: "one",
				// DataType: 0,
			},
			{
				Name: "two",
				// DataType: 0,
			},
		},
		Data: []api.DataRow{
			{
				Values: []string{"a", "b"},
			},
			{
				Values: []string{"c", "d"},
			},
		},
	}
	// var err error
	// if err = sysdb.EnableConnector(&rq); err != nil {
	// 	util.HandleError(w, err, http.StatusBadRequest)
	// 	return
	// }
	// success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(cmdr); err != nil {
		// TODO error handling
		_ = err
	}
}

func ReadRequest(w http.ResponseWriter, r *http.Request, requestStruct interface{}) bool {
	var body []byte
	var err error
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return false
	}
	fmt.Printf("%s\n", body)
	if err = json.Unmarshal(body, requestStruct); err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return false
	}
	log.Trace("request %s %v\n", r.RemoteAddr, requestStruct)
	return true
}

func HandleError(w http.ResponseWriter, err error, statusCode int) {
	log.Error("%s", err)
	HTTPError(w, err, statusCode)
}

func HTTPError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	var m = map[string]interface{}{
		"status": "error",
		//"message": fmt.Sprintf("%s: %s", http.StatusText(code), err),
		"message": err.Error(),
		"code":    code,
		//"data":    "",
	}
	//json.NewEncoder(w).Encode(err)
	if err = json.NewEncoder(w).Encode(m); err != nil {
		// TODO error handling
		_ = err
	}
}

func requestString(r *http.Request) string {
	var remoteHost, remotePort string
	remoteHost, remotePort, _ = net.SplitHostPort(r.RemoteAddr)
	return fmt.Sprintf("host=%s port=%s method=%s uri=%s", remoteHost, remotePort, r.Method, r.URL)
}
