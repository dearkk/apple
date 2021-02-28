package src

type HelloObject struct {
}

func (h *HelloObject) Add(a int, b int) int {
	return a + b
}
