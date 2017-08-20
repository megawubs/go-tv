package tv

var television *Television = MakeTelevision()

type Command struct {
	Action Action
}

type Action int

const (
	Up Action = iota
	Down
	Power
	VolumeUp
	VolumeDown
)

func Receive(command Command)  (*Television, bool)  {
	defer television.Info()
	if television.On == false && command.Action != Power {
		return nil, true
	}
	switch command.Action {
	case Up:
		television.Up()
		return television, false
	case Down:
		television.Down()
		return television, false
	case Power:
		television.Power()
		return television, false
	case VolumeDown:
		television.VolumeDown()
		return television, false
	case VolumeUp:
		television.VolumeUp()
		return television, false
	}

	return nil, true
}
