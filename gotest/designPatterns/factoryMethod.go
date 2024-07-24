package designpatterns

import "errors"

//抽象工厂，用一个固定的一套模板，生成不同类别的对象

//工厂-->电脑，cpu、内存、显卡、固态、主板

type ComputerFactory interface {
	makeCpu() MCpu
	makeCache() MCache
	makeGpu() MGpu
	//...
}

func getComputer(brand string) (ComputerFactory, error) {
	var computer ComputerFactory
	if brand == "INVIDA" {
		computer = &INVIDA{}
	} else if brand == "AMD" {
		computer = &AMD{}
	}
	if computer == nil {
		return computer, errors.New("该类型不存在")
	} else {
		return computer, nil
	}
}

type MCpu interface {
	setName(name string)
}

type Cpu struct {
	Name string
}

// 利用匿名参数的伪继承属性
type Intel struct {
	Cpu
}

type AmdC struct {
	Cpu
}

func (cpu *Cpu) setName(name string) {
	cpu.Name = name
}

type MCache interface {
	setSize(size float64)
}

type Kinston struct {
	cache
}

type cache struct {
	size float64
}

func (cache *cache) setSize(size float64) {
	cache.size = size
}

type MGpu interface {
	setName(name string)
}

type gpu struct {
	name string
}

func (gpu *gpu) setName(name string) {
	gpu.name = name
}

func (gpu *gpu) getName() string {
	return gpu.name
}

type Invida struct {
	gpu
}

type AmdG struct {
	gpu
}

// INVIDA电脑配置
type INVIDA struct {
}

func (com *INVIDA) makeCpu() MCpu {
	return &Intel{Cpu: Cpu{
		Name: "invida",
	}}
}

func (com *INVIDA) makeCache() MCache {
	return &Kinston{}
}

func (com *INVIDA) makeGpu() MGpu {
	return &Invida{}
}

// AMD电脑配置
type AMD struct {
}

func (com *AMD) makeCpu() MCpu {
	return &Intel{}
}

func (com *AMD) makeCache() MCache {
	return &Kinston{}
}

func (com *AMD) makeGpu() MGpu {
	return &Invida{}
}
