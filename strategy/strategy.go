package strategy

import (
	"fmt"
)

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

// Cash 现金策略
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by cash", ctx.Money, ctx.Name)
}

// Bank 银行策略
type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}

// PaymentContext 支付策略的上下文
type PaymentContext struct {
	Name, CardID string
	Money        int
}

type Payment struct {
	ctx      *PaymentContext
	strategy PaymentStrategy
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.ctx)
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		ctx: &PaymentContext{
			Name:   name,
			CardID: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}
