package demo

import (
	"github.com/bitwormhole/gss/application"
	"github.com/bitwormhole/gss/lang"
)

// Driver class
type Driver struct {
	car  *Car
	name string
}

// Car class
type Car struct {
	id      string
	context application.Context

	driver *Driver

	wheelFrontLeft  *Wheel
	wheelFrontRight *Wheel
	wheelBackLeft   *Wheel
	wheelBackRight  *Wheel
}

// Wheel class
type Wheel struct {
	car   *Car
	id    string
	phase float32
}

// impl Car

func (inst *Car) start() error {
	return nil
}

func (inst *Car) stop() error {
	return nil
}

// impl Driver

func (inst *Driver) xxx() {}

// impl Wheel

func (inst *Wheel) xxx() {}

// component of Driver

// component of Wheel

// component of Car

type carCom struct {
	application.ComponentRegistration
	application.Component `id:"car1"  class:"car"  initMethod:"start"  destroyMethod:"stop"  `

	myCar *Car `id:"#car1"  driver:".driver"  door:".door"  `
}

func (inst *carCom) Inject(ctx application.Context) error {

	getter := ctx.NewGetter()
	ok := false

	inst.myCar.context = ctx
	inst.myCar.id = getter.GetProperty("car.id")

	inst.myCar.driver, ok = getter.GetComponentByClass(".driver").(*Driver)
	getter.Feedback(ok, "")

	inst.myCar.wheelBackLeft, ok = getter.GetComponent("#wheel-front-left").(*Wheel)
	getter.Feedback(ok, "")
	inst.myCar.wheelBackRight, ok = getter.GetComponent("#wheel-front-right").(*Wheel)
	getter.Feedback(ok, "")
	inst.myCar.wheelFrontLeft, ok = getter.GetComponent("#wheel-back-left").(*Wheel)
	getter.Feedback(ok, "")
	inst.myCar.wheelFrontRight, ok = getter.GetComponent("#wheel-back-right").(*Wheel)
	getter.Feedback(ok, "")

	return getter.Result()
}

// 在正式版本中， 以下这些函数由自动化工具生成

// auto-gen
func (inst *carCom) GetComponent() lang.Object {
	return inst.myCar
}

// auto-gen
func (inst *carCom) GetInfo() application.ComponentInfo {
	info := &application.ComponentInfo{
		Name:  "",
		Class: "",
		Scope: 0,
	}
	return *info
}

// auto-gen
func (inst *carCom) NewInstance() application.ComponentInstance {
	mycar := &Car{}
	return &carCom{
		myCar: mycar,
	}
}

// auto-gen
func (inst *carCom) Init() error {
	return inst.myCar.start()
}

// auto-gen
func (inst *carCom) Destroy() error {
	return inst.myCar.stop()
}

// RegisterComponents 注册本宝宝的组件 (auto-gen)
func RegisterComponents(cfg *application.Configuration) {

	cfg.Component(&carCom{})

}
