package tv

import "fmt"

const (
	Up int = iota
	Down
	Power
	VolumeUp
	VolumeDown
)

type Command struct {
	ID int `json:"ID"`
}

// TV represent a television set with state and the ability to receive commands from a remote.
type TV struct {
	// No need for these to be public if they aren't accessed directly.
	volume  int
	channel int
	on      bool
	Receive chan Command
}

// NewTV returns a TV with some default values. `NewXXXX` is the idiomatic format for constructor functions like this.
func NewTV() *TV {
	return &TV{
		volume:  25,
		channel: 1,
		// Boolean values default to `false` so we usually don't bother doing this explicitly.
		Receive: make(chan Command),
	}
}

func (t *TV) power() {
	t.on = !t.on
}

func (t *TV) up() {
	t.channel++
}

func (t *TV) down() {
	t.channel--
}

func (t *TV) volumeUp() {
	t.volume++
}

func (t *TV) volumeDown() {
	t.volume--
}

// PrintInfo prints the state of the TV to stdout.
func (t TV) PrintInfo() {
	fmt.Printf("Volume: %d\t Channel: %d\t Powered on: %t\n", t.volume, t.channel, t.on)
}

// ListenAndLog reads incoming commands on the receive channel and handles them,
// before printing out the new state of the television.
func (t *TV) ListenAndLog() {
	for cmd := range t.Receive {
		if !t.on && cmd.ID != Power {
			continue
		}

		switch cmd.ID {
		case Up:
			t.up()
		case Down:
			t.down()
		case Power:
			t.power()
		case VolumeDown:
			t.volumeDown()
		case VolumeUp:
			t.volumeUp()
		}
		t.PrintInfo()
	}
}
