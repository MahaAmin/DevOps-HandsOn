package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	backendHost = flag.String("backend-host", "localhost", "Host for backend service")
	backendPort = flag.Int("backend-port", 80, "Port for backend service")
	listenAddr  = flag.String("listen-addr", ":8080", "Listen address")
)

var responseTemplate = `
<p>
Backend response code: {{ .Code }} </br>
Backend response: {{ .Body }} </br>
</p>
`

var compiledResponseTemplate *template.Template

func main() {
	var err error
	flag.Parse()
	http.Handle("/callBackend", &backendHandler{})
	http.Handle("/", &healthHandler{})
	compiledResponseTemplate, err = template.New("response").Parse(responseTemplate)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(*listenAddr, nil)
}

type backendHandler struct{}
type healthHandler struct{}
type templateValuesContainer struct {
	Code int
	Body string
}

func (b backendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/log", *backendHost, *backendPort), nil)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("ECODE1"))
		log.Print(err)
		return
	}

	requestID := r.Header.Get("X-Request-Id")
	if requestID == "" {
		log.Print("Request ID not present")
	} else {
		req.Header.Add("X-Request-Id", requestID)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("ECODE2"))
		log.Print(err)
		return
	}

	templateValues := templateValuesContainer{}

	templateValues.Code = resp.StatusCode
	if resp.StatusCode == 200 {
		bts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("ECODE3"))
			log.Print(err)
			return
		}
		templateValues.Body = string(bts)
	}

	compiledResponseTemplate.Execute(w, templateValues)
}

func (b healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
