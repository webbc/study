package observer

type WeatherSubject struct {
	Subject
	content string
}

func (ws *WeatherSubject) SetContent(content string) {
	ws.content = content
	ws.NotifyObservers(ws, ws.content)
}
