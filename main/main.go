package main

import (
	"github.com/megawubs/tv"
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

func decodeCommand(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	decoder := json.NewDecoder(request.Body)
	var command tv.Command
	err := decoder.Decode(&command)
	if err != nil {
		panic(err)
	}

	instance,_ := tv.Receive(command)
	io.WriteString(writer, instance.Info())
	fmt.Printf("%s\n", instance.Info())
}

func main()  {
	http.HandleFunc("/", decodeCommand)
	http.ListenAndServe(":8000", nil)
}

