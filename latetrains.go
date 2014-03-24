package main

import (
	"fmt"
	"math/cmplx"
	"math"
	"github.com/sivvo/newmath"
)

func add(x, y int) int {
    return x + y
}


var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)


func main() {
    const myconst = "a constant"
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))
    	fmt.Println(add(42, 13))

    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)

    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)

    var x, y int = 3, 4
    var fl float64 = math.Sqrt(float64(3*3 + 4*4))
    var z int = int(fl)
    fmt.Println(x, y, z)

        sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
    fmt.Println("end of line")

}
