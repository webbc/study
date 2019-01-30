package main

import "design_pattern/adapter"

func main() {

	usber := &adapter.Usber{}
	adaper := adapter.NewAdapter(usber)

	adaper.IsPs2()
}
