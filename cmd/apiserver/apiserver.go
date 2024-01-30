package main

import "fmt"

// In actually, we can't include much complex logic in here. if you think the code can be imported or can be
// used in others projects. then it should be in the /pkg directory. if the code is not reusable. Or you don't want others
// import or reuse it. then the code should be in the /internal directory
func main() {
	fmt.Printf("apiserver")
}
