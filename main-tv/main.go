package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hdhauk/tv/tv"
)

func main() {
	t := tv.NewTV()
	go t.ListenAndLog()

	http.HandleFunc("/",
		// This is an anonymous function. This lets us use the `t`-variable that live in our main function,
		// instead of having a TV-singleton be instantiated as a package global variable in the tv-package,
		// as one usually should avoid global variables (even those limited to one package).
		func(w http.ResponseWriter, r *http.Request) {
			var cmd tv.Command
			if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
				log.Fatalln(err)
				// log.Fatalln is equivalent to:
				//		log.Println(err)
				//		os.Exit(1)
			}
			t.Receive <- cmd
		})

	// Start listening for http requests.
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}
