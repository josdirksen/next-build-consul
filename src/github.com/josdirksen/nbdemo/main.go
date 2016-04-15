package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"text/template"
	"flag"
	"io/ioutil"
)

var count = 0
var serverAddress *string

// a very basic handler, which just shows the environment variables
// and some other information to identify this server.
func backendHandler(w http.ResponseWriter, r *http.Request) {

	// we load a simple external HTML template to format the output and return that
	var Template2, _ = template.New("backend-template.html").ParseFiles("./resources/backend-template.html")

	// output HTML
	w.Header().Set("Content-Type", "text/html")

	// fill the input values for our HTML template
	data := struct {
		Envs       []string
		Network    []string
		ServerName string
	}{
		Envs:       envVariables(),
		Network:    networkAddresses(),
		ServerName: getProviderServerName(),
	}

	Template2.Execute(w, data)
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {

	if (strings.HasSuffix(r.RequestURI, "/api")) {
		var address = fmt.Sprintf("http://%s/", *serverAddress)
		resp, err := http.Get(address)
		if err != nil {
			// handle error
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			w.Header().Set("Content-Type",resp.Header.Get("Content-Type"))
			w.Write(body)
		}
	} else {
		// we load a simple external HTML template to format the output and return that
		var Template2, _ = template.New("frontend-template.html").ParseFiles("./resources/frontend-template.html")
		w.Header().Set("Content-Type", "text/html")

		// fill the input values for our HTML template
		data := struct {
			Envs       []string
			Network    []string
			ServerName string
		}{
			Envs:       envVariables(),
			Network:    networkAddresses(),
			ServerName: getProviderServerName(),
		}

		Template2.Execute(w, data)
	}
}


func backendApiHandler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.RequestURI, "ico") {
	} else {
		count = count + 1
		w.Header().Set("Content-Type", "application/json")
		var response = `
	{"result" : {
	  "servername" : "%s",
	  "querycount" : %d
	  }
	}
	`
		fmt.Fprintf(w, response, getProviderServerName(), count)
	}
}


func networkAddresses() []string {
	var result []string

	interfaces, _ := net.Interfaces()

	for _, interf := range interfaces {
		var buffer bytes.Buffer

		buffer.WriteString(interf.Name + ": ")
		intAddresses, _ := interf.Addrs()
		for _, interfAddress := range intAddresses {
			buffer.WriteString("\"" + interfAddress.String() + "\"  ")
		}
		result = append(result, string(buffer.Bytes()))
	}

	return result
}

func getProviderServerName() string {
	name, found := os.LookupEnv("SERVERNAME")
	if found {
		return name
	} else {
		return "unspecified"
	}
}

func envVariables() []string {
	return os.Environ()
}

// run a webserver on port 8080
func main() {

	var t = flag.String("type", "backend", "Either a backend or frontend service")
	serverAddress = flag.String("serverAddress", "backend-service:8081", "the address of the backend server")
	flag.Parse()

	if (*t == "backend") {
		backend := &http.Server{
			Addr:    ":8080",
			Handler: http.HandlerFunc(backendHandler),
		}
		api := &http.Server{
			Addr:    ":8081",
			Handler: http.HandlerFunc(backendApiHandler),
		}

		go api.ListenAndServe()
		backend.ListenAndServe()

	} else {
		frontend := &http.Server{
			Addr:    ":8090",
			Handler: http.HandlerFunc(frontendHandler),
		}
		frontend.ListenAndServe()
	}
}
