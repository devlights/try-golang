package main

import "fmt"

func main() {
    func01()

    var ret = func02(10, 20)
    fmt.Println(ret)
}
func func02(x int, y int) int {
    return x + y
}

func func01() {
    fmt.Println("func01")
}
