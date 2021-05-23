package strategy

import "fmt"

//Payment 支付对象
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy // 支付策略
}

//PaymentContext 支付需要的基本信息
type PaymentContext struct {
	Name, CardID string
	Money        int
}

//NewPayment 策略选择
func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

//Pay 策略最后的调用
func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

//PaymentStrategy 字符策略接口
type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type Cash struct{}

//Pay 策略1
func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type Bank struct{}

//Pay 策略2
func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}
