package cor

import "fmt"

// 经理类
type ManagerHandler struct {
	BaseHandler
}

func (h *ManagerHandler) PriceHandle(price float64) {
	if price < 0.3 {
		fmt.Printf("ManagerHandler 通过 , price:%.2f\n", price)
	} else {
		h.successor.PriceHandle(price)
	}
}
