package concurrent

import (
	"context"
	"fmt"
	"testing"
)

type orderID int
type traceId int

// https://github.com/cch123/golang-notes/blob/master/context.md
// valueCtx 主要就是用来携带贯穿整个逻辑流程的数据的，
// 在分布式系统中最常见的就是 trace_id，在一些业务系统中，
// 一些业务数据项也需要贯穿整个请求的生命周期，如 order_id，payment_id 等。
func TestValueContext(t *testing.T) {
	var x = context.TODO()
	x = context.WithValue(x, orderID(1), "1234")
	x = context.WithValue(x, orderID(2), "2345")

	y := context.WithValue(x, orderID(3), "4567")
	x = context.WithValue(x, orderID(3), "3456")

	fmt.Println(x.Value(orderID(1)))
	fmt.Println(y.Value(orderID(3)))
}
