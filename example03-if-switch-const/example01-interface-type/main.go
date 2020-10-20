package main

var (
	i interface{}
)

func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is an integer", t)
	case string:
		println("i is an string", t)
	case float64:
		println("i is an float64", t)
	default:
		println("type not found")
	}
}

func main() {
	i = 100
	convert(i)
	i = float64(45.55)
	convert(i)
	i = "foo"
	convert(i)
	convert(float32(10.0))
}
