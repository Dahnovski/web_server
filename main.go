import (
	"fmt"
	"log"
	"net/http"
)

func greetingHandler(w http.ResponseWriter, r *http) {
	/// -> TODO
}

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
