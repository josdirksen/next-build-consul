package main

import (
	"net/http"
	"os"
	"text/template"
	"net"
	"bytes"
)

// a very basic handler, which just shows the environment variables
// and some other information to identify this server.
func handler(w http.ResponseWriter, r *http.Request) {

	// we load a simple external HTML template to format the output and return that
	var Template2, _ = template.New("backend-template.html").ParseFiles("./resources/backend-template.html")

	// output HTML
	w.Header().Set("Content-Type", "text/html")

	// fill the input values for our HTML template
	data := struct {
		Envs []string
		Network[] string
		ServerName string
	}{
		Envs: envVariables(),
		Network: networkAddresses(),
		ServerName: getProviderServerName(),
	}

	Template2.Execute(w, data)
}

func networkAddresses() ([]string) {
	var result []string


	interfaces,_ := net.Interfaces()

	for _,interf := range interfaces {
		var buffer bytes.Buffer

		buffer.WriteString(interf.Name + ": ")
		intAddresses,_ := interf.Addrs()
		for _,interfAddress := range intAddresses {
			buffer.WriteString("\"" +interfAddress.String() + "\"  ")
		}
		result = append(result, string(buffer.Bytes()))
	}

	return result
}

func getProviderServerName() (string) {
	name,found:= os.LookupEnv("SERVERNAME")
	if (found) {
		return name
	} else {
		return "unspecified"
	}
}

func envVariables() ([]string) {
	return os.Environ()
}

// run a webserver on port 8080
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
