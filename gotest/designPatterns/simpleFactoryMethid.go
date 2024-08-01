package designpatterns

import "errors"

type IFactory interface {
	setName(name string)
	setAI(tp string)
	getName() string
	getAI() string
}

type CameraDev struct {
	name   string
	AIType string
}

func (dev *CameraDev) setName(name string) {
	dev.name = name
}

func (dev *CameraDev) setAI(tp string) {
	dev.AIType = tp
}

func (dev *CameraDev) getName() string {
	return dev.name
}

func (dev *CameraDev) getAI() string {
	return dev.AIType
}

type GB281 struct {
	CameraDev
}

func newGB281(name, tp string) IFactory {
	return &GB281{
		CameraDev{
			name:   name,
			AIType: tp,
		},
	}
}

type AI360Pro struct {
	CameraDev
}

func newAI360Pro(name, tp string) IFactory {
	return &AI360Pro{
		CameraDev{
			name:   name,
			AIType: tp,
		},
	}
}

func getcamerDev(name, AIType string) (IFactory, error) {
	if name == "GB281" {
		return newGB281(name, AIType), nil
	} else if name == "AI360Pro" {
		return newAI360Pro(name, AIType), nil
	}
	return nil, errors.New("NO such camera!!!")
}
