package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	content, err := ioutil.ReadFile("/home/pradhyushrestha/.bashrc")
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	text := string(content)
	fmt.Fprintf(w, text)
	fmt.Fprint(w, exampleCommand("find", "/home/pradhyushrestha/"))
}

func exampleCommand(command string, params ...string) string {
	cmd := exec.Command(command, strings.Join(params, " "))
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	return out.String()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
