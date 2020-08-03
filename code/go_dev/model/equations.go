package model

//二元线性方程组
type BinaryLinearEquations struct{
	A11 float64
	A12 float64
	A21 float64
	A22 float64
	X1 float64
	X2 float64
	Y1 float64
	Y2 float64
}

//二元线性方程组的工厂方法
func NewBinaryLinear(a11 float64, a12 float64,a21 float64,a22 float64, y1 float64, y2 float64) *BinaryLinearEquations {
	return &BinaryLinearEquations{
		A11:a11,
		A12:a12,
		A21:a21,
		A22:a22,
		Y1:y1,
		Y2:y2,
    }
}

//检查二元线性方程组是否有唯一解
func (this BinaryLinearEquations) TestResult() bool{
	var flag = this.A11*this.A22 - this.A12*this.A21
	//不等于0时有唯一解
	if(flag != 0){
        return true
	}else{
		return false
	}
}

//求解二元线性方程组,未知数将被赋值并返回
func (this BinaryLinearEquations) GetResult() *BinaryLinearEquations{
	var D = this.A11*this.A22 - this.A12*this.A21
	if(D == 0){
		return nil
	}
	var D1 = this.Y1*this.A22 - this.A12*this.Y2
	var D2 = this.A11*this.Y2 - this.Y1*this.A21
	this.X1 = D1/D
	this.X2 = D2/D
	return &this
}



