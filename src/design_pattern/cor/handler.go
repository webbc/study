package cor

// 责任链模式

// 处理接口
type Handler interface {
	PriceHandle(float64) // 处理价格的方法
}

// 基础的处理类
type BaseHandler struct {
	successor Handler
}
