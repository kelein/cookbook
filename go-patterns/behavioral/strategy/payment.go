package strategy

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// PayBehavior of abstract
type PayBehavior interface {
	OrderPay(ctx *PayContext)
}

// WechatPay pay by Wechat
type WechatPay struct{}

// OrderPay of Wechat
func (w *WechatPay) OrderPay(ctx *PayContext) {
	log.Printf("WechatPay deal with requst: %v", ctx.params)
	log.Printf("WechatPay payment done")
}

// ThirdPartyPay pay by ThirdParty
type ThirdPartyPay struct{}

// OrderPay of ThirdPartyPay
func (t *ThirdPartyPay) OrderPay(ctx *PayContext) {
	log.Printf("ThirdPartyPay deal with requst: %v", ctx.params)
	log.Printf("ThirdPartyPay payment done")
}

// PayContext context of Pay
type PayContext struct {
	behavior PayBehavior
	params   map[string]interface{}
}

func (ctx *PayContext) setBehavior(behavior PayBehavior) {
	ctx.behavior = behavior
}

// Pay for order with different behavior
func (ctx *PayContext) Pay() {
	ctx.behavior.OrderPay(ctx)
}

// NewPayContext create a PayContext instance
func NewPayContext(behavior PayBehavior) *PayContext {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &PayContext{
		behavior: behavior,
		params: map[string]interface{}{
			"AppID": fmt.Sprintf("%03d", random.Intn(999)),
			"Order": fmt.Sprintf("%06d", random.Intn(999999)),
		},
	}
}
