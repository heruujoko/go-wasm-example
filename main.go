package main

import (
	"math"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	println("init go wasm engine")
	registerCallbacks()
	<-c
}

// exposed func(s)

func add(this js.Value, i []js.Value) interface{} {
	num1 := i[0].Int()
	num2 := i[1].Int()

	result := num1 + num2
	println(result)
	return js.ValueOf(result)
}

func intensiveTask(this js.Value, i []js.Value) interface{} {
	var result float64 = 0.0
	var baseNumber = i[0].Float()
	for i := math.Pow(baseNumber, 7); i >= 0; i-- {
		result += math.Atan(i) * math.Tan(i)
	}
	return js.ValueOf(result)
}

// expose function to JS runtime
func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("intensiveTask", js.FuncOf(intensiveTask))
}
