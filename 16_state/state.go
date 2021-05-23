package state

import "fmt"

//Week 接口,定义了后续每个对象都需要实现的方法
type Week interface {
	Today()
	Next(*DayContext)
}

//DayContext 定义了当前状态
type DayContext struct {
	today Week
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

//NewDayContext 入口
func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

//Sunday 定义每个状态
type Sunday struct{} // 这里为什么不继承DayContext?

//Today 当前状态
func (*Sunday) Today() {
	fmt.Printf("Sunday\n")
}

//Next 下个状态
func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

//下面都是重复定义每个阶段的状态以及对应的转换
type Monday struct{}

func (*Monday) Today() {
	fmt.Printf("Monday\n")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct{}

func (*Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

func (*Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct{}

func (*Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct{}

func (*Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct{}

func (*Friday) Today() {
	fmt.Printf("Friday\n")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct{}

func (*Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}
