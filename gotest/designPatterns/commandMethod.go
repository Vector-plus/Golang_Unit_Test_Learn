package designpatterns

import "fmt"

type Command interface {
	Excute()
}

type Device interface {
	On()
	Off()
}

type OnCommand struct {
	Device
}

func (on *OnCommand) Excute() {
	on.Device.On()
}

type OffCommand struct {
	Device
}

func (off *OffCommand) Excute() {
	off.Device.Off()
}

type TV struct{}

func (tv *TV) On() {
	fmt.Println("TV is turn on")
}

func (tv *TV) Off() {
	fmt.Println("TV is turn off")
}
