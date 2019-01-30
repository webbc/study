package observer

import "fmt"

type WeatherObserver struct {
	name string
	do   string
}

func (wo *WeatherObserver) SetName(name string) {
	wo.name = name
}

func (wo *WeatherObserver) SetDo(do string) {
	wo.do = do
}

func (wo *WeatherObserver) Update(subject ISubject, args ...interface{}) {
	fmt.Println(wo.name + "收到天气信息：" + subject.(*WeatherSubject).content + "," + wo.do)
	fmt.Println(wo.name + "收到天气信息：" + args[0].(string) + "," + wo.do)
}
