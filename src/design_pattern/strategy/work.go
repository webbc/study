package strategy

type Work struct {
	Tool IStrategy // 交通工具
}

// 上班
func (w *Work) GoToWork() {
	w.Tool.Do()
}
