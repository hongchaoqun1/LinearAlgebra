package main

import(
	"go_dev/model"
	"fmt"
)

func main()  {
	binaryLinearEquations := new (model.BinaryLinearEquations)
	binaryLinearEquations = model.NewBinaryLinear(2,1,3,-2,4,-1)
	fmt.Println(binaryLinearEquations.TestResult())
	binaryLinearEquations = binaryLinearEquations.GetResult()
	fmt.Println(binaryLinearEquations.X1, binaryLinearEquations.X2)
}