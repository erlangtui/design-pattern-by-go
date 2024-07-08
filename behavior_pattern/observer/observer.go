// 观察者模式
package main

import (
	"fmt"
)

// Subject 接口，它相当于是发布者的定义
type Subject interface {
	Subscribe(observer Observer)
	Notify(msg string)
}

// Observer 观察者接口
type Observer interface {
	Update(msg string)
}

// Subject 实现
type SubjectImpl struct {
	observers []Observer
}

// Subscribe 添加观察者（订阅者）
func (sub *SubjectImpl) Subscribe(observer Observer) {
	sub.observers = append(sub.observers, observer)
}

// Notify 发布通知
func (sub *SubjectImpl) Notify(msg string) {
	// 事件发布给订阅者的过程，其实就是遍历一下已经注册的事件订阅者，逐个去调用订阅者实现的观察者接口方法，
	// 比如叫 handleEvent 之类的方法，这个方法的参数一般就是当前的事件对象。
	for _, o := range sub.observers {
		o.Update(msg) // 同步调用，逐个等待订阅着执行完成
	}
}

// Observer1 Observer1
type Observer1 struct{}

// Update 实现观察者接口
func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

// Observer2 Observer2
type Observer2 struct{}

// Update 实现观察者接口
func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}

func main() {
	// 与 sync.Cond 类似
	sub := &SubjectImpl{}
	sub.Subscribe(&Observer1{})
	sub.Subscribe(&Observer2{})
	sub.Notify("Hello")
}
