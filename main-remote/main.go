package main

import (
	"fmt"
	"strconv"

	"github.com/hdhauk/tv/remote"
	"github.com/hdhauk/tv/tv"
)

func main() {
	for {
		fmt.Print("Choose action: [0:Up, 1:Down, 2:Power, 3:Volume Up, 4: Volume Down] ")

		// Read from user aka. `os.stdin`
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Printf("Failed to read input %v\n", err)
			continue
		}

		// Attempt to convert string to an integer.
		cmd, err := strconv.Atoi(input) // Atoi stands for ASCII-to-Integer if the name seem strange.
		if err != nil {
			fmt.Printf("Failed to parse input: %v\n", err)
			continue
		}

		// Send the command to the server.
		if err := remote.Send(tv.Command{ID: cmd}); err != nil {
			fmt.Printf("Failed to send command: %v\n", err)
		}
	}
}
