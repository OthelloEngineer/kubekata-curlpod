package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

const commandsDescriptions = `
/curl?url=<url> - curl the url
/nslookup?url=<url> - nslookup the url
/ping?url=<url> - ping the url
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(commandsDescriptions))
	})

	http.HandleFunc("/curl", curl)
	http.HandleFunc("/nslookup", nslookup)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/ls"", ls)
	http.HandleFunc("/cat", cat)
	http.ListenAndServe(":8080", nil)
}

func curl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	cmd := exec.Command("curl", " -Lk ", url)
	out, err := cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte(err.Error()))
	}
	fmt.Print(string(out))
	w.Write(out)
}

func nslookup(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	cmd := exec.Command("nslookup", url)
	out, err := cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte(err.Error()))
	}
	fmt.Print(string(out))
	w.Write(out)
}

func ping(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	cmd := exec.Command("ping", url, "-c", "3")
	out, err := cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte(err.Error()))
	}
	fmt.Print(string(out))
	w.Write(out)
}

func ls(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ls")
	out, err := cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte(err.Error()))
	}
	fmt.Print(string(out))
	w.Write(out)
}

func cat(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("cat")
	out, err := cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte(err.Error()))
	}
	fmt.Print(string(out))
	w.Write(out)
}
