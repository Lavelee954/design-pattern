package _interface

import "example.com/design-pattern/Observer/ex_code/subject"

type Subject interface {
	Register(observer subject.Observer)
	Deregister(observer subject.Observer)
	NotifyAll()
}
