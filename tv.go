package tv

import (
	"encoding/json"
)

type Television struct {
	Volume int
	Channel int
	On bool
}

func MakeTelevision() *Television  {
	return &Television{Volume:25, Channel:1, On:false}
}

func (t *Television) Power(){
	t.On = ! t.On
}

func (t *Television) Up()  {
	t.Channel++
}

func (t *Television) Down(){
	t.Channel--
}

func (t *Television) VolumeUp() {
	t.Volume++
}

func (t *Television) VolumeDown() {
	t.Volume--
}

func (t *Television) Info() string {
	j,_ := json.Marshal(t)
	return string(j)
}
