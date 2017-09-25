package adding

import(
	"fmt"

)

func init(){
	fmt.Println("Adding goodly")
}

func AddWrong(x, y int) int{
	return x + y + 1
}
