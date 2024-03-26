package example

import (
	"fmt"
)
import _ "embed"

//go:embed "/README.md"
var version string

type number interface {
	~int | ~float32 | ~float64
}

type Person[T number] struct {
	name T
	age  T
}

func (app *Person[T]) class(t, s T) T {
	return t / s
}

type apple struct {
	name string
	age  int
}

type appleInterface interface {
	getAge(age int) string
}

func multiple[T number](a, b T) T {
	return a + b
}

func add[T int | float32 | float64](a, b T) T {
	return a + b
}

func decrease[T any](a, b T) {
	fmt.Printf("first is %s, and second is %s", a, b)
}

func (app *apple) getAge(age int) string {
	add(1.2, 2.3)
	go run()
	return "wpc"
}

func suffix(age int) (app *apple) {
	return &apple{name: "wpc", age: 20}
}

func run() {
	app := suffix(20)
	println(app.name)
	fmt.Printf("%v", app.age)
}
