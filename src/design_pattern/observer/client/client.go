package main

import "design_pattern/observer"

func main() {

	// 初始化目标
	subject := &observer.WeatherSubject{}

	// 初始化观察者
	a_observer := &observer.WeatherObserver{}
	a_observer.SetName("A")
	a_observer.SetDo("出去打球")
	b_observer := &observer.WeatherObserver{}
	b_observer.SetName("B")
	b_observer.SetDo("出去逛街")

	// 注册
	subject.Attach(a_observer)
	subject.Attach(b_observer)

	// 通知
	subject.SetContent("天气晴朗，室温28°C")

	// 解绑
	subject.Detach(a_observer)
	subject.Detach(b_observer)

	// 再次通知
	subject.SetContent("天气晴朗，室温10°C")
}
