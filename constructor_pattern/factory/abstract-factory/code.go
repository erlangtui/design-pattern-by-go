// 抽象工厂模式
package main

import "fmt"

// 示例源码运行Demo
// 给公众号「网管叨bi叨」发生 go-factory 即可领取
type AbstractFactory interface {
	CreateTV() TV
	CreateAirCleaner() AirCleaner
}

type HuaWeiFactory struct{}

func (hf *HuaWeiFactory) CreateTV() TV {
	return &HuaWeiTV{}
}
func (hf *HuaWeiFactory) CreateAirCleaner() AirCleaner {
	return &HuaWeiAirCleaner{}
}

type XiaoMiFactory struct{}

func (mf *XiaoMiFactory) CreateTV() TV {
	return &XiaoMiTV{}
}
func (mf *XiaoMiFactory) CreateAirCleaner() AirCleaner {
	return &XiaoMiAirCleaner{}
}

type TV interface {
	Watch()
}

type HuaWeiTV struct{}

func (ht *HuaWeiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

type XiaoMiTV struct{}

func (mt *XiaoMiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

type AirCleaner interface {
	SetTemperature(int)
}

type HuaWeiAirCleaner struct{}

func (ha *HuaWeiAirCleaner) SetTemperature(temp int) {
	fmt.Printf("HuaWei AirConditioner set temperature to %d ℃\n", temp)
}

type XiaoMiAirCleaner struct{}

func (ma *XiaoMiAirCleaner) SetTemperature(temp int) {
	fmt.Printf("Mi AirConditioner set temperature to %d ℃\n", temp)
}

func main() {
	var factory AbstractFactory
	var tv TV
	var air AirCleaner

	factory = &HuaWeiFactory{}
	tv = factory.CreateTV()
	air = factory.CreateAirCleaner()
	tv.Watch()
	air.SetTemperature(25)

	factory = &XiaoMiFactory{}
	tv = factory.CreateTV()
	air = factory.CreateAirCleaner()
	tv.Watch()
	air.SetTemperature(26)
}
