package cor

import "fmt"

// 销售类
type SalesHandler struct {
	BaseHandler
}

func (h *SalesHandler) PriceHandle(price float64) {
	if price < 0.1 {
		fmt.Printf("SalesHandler 通过 , price :%.2f\n", price)
	} else {
		h.successor.PriceHandle(price)
	}
}
