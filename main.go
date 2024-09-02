import (
	"fmt"
	"log"
	"net/http"
)

func greetingHandler(w http.ResponseWriter, r *http) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound) {
			return
		}
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, user!")
}

func formHandler(w http.ResponseWriter, r *http) {
	/// -> TODO
}

/// -> TODO

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer) 
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/greeting", greetingHandler)

	fmt.Print("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
