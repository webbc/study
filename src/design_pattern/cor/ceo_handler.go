package cor

import "fmt"

// 老板类
type CeoHandler struct {
	BaseHandler
}

func (h *CeoHandler) PriceHandle(price float64) {
	if price < 0.5 {
		fmt.Printf("CeoHandler 通过 , price :%.2f\n", price)
	} else {
		fmt.Printf("CeoHandler 不通过 , price :%.2f\n", price)
	}
}
