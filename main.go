// test project main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Can be loaded from config or from command arguments
const MAX_LINE_PER_REQUEST = 999

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Only POST method allowed!\n"))
		return
	}

	err, lines := Parse(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: " + err.Error() + "\n"))
		return
	}

	err, result := Loader(lines, Load)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: " + err.Error() + "\n"))
		return
	}

	w.Write([]byte(strings.Join(result, "\n")))
}

func main() {
	fmt.Println("Begin")

	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Server error\n")
		os.Exit(1)
	}
}
