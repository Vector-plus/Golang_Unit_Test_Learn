package designpatterns

import "fmt"

// 观察者模式，类似于订阅发布模式
type Production interface {
	register(id string, consum Consumer)
	deregister(id string)
	notify(name string)
}

type Consumer interface {
	update(id string, name string)
}

type Observe struct {
	consumers map[string]Consumer
	name      string
}

func (o *Observe) ExistProducts() {
	o.notify(o.name)
}

func (o *Observe) register(id string, consumer Consumer) {
	if _, ok := o.consumers[id]; !ok {
		o.consumers[id] = consumer
	}
}

func (o *Observe) deregister(id string) {
	if _, ok := o.consumers[id]; ok {
		delete(o.consumers, id)
	}
}

func (o *Observe) notify(name string) {
	for id, consumer := range o.consumers {
		consumer.update(id, name)
	}
}

type Subscribe struct {
	id string
}

func (sub *Subscribe) update(id string, name string) {
	fmt.Println(name + "----> 已通知：" + id)
}
