package gopolynomial

import (
		"fmt"
		"gonum.org/v1/gonum/mat"
		)

/*
Poly1D is a type for storing information of a ID polynomial
It is also use for manipulating a 1D polynomial
*/

type Poly1D struct{
	Coeffs []float64
	Degree int
	Roots []complex128
	}

/* 
NewPoly1D is used to create Poly1D type
Poly1D can be created from roots of the polynomial or from the coefficient
The coefficient is a slice/array of type float64
The roots can be an array/slice of type float64 or complex128
inputs:
array:This is an array or roots or coefficients
r: This is a boolean variable. If true, the first input is taken as an array of roots
*/

func NewPoly1D(array interface{}, r bool) (Poly1D){
	var roots []complex128
	var coeffs []float64
	var p Poly1D
	if r==false{
		arr:=array.([]float64)
		roots=PolyRoots(arr)
		coeffs=arr
		degree:=len(arr)-1
		p=Poly1D{coeffs,degree,roots}
	}else{
		switch array.(type){
			case []complex128:
				arr:=array.([]complex128)
				roots=arr
				coeffs=PolyCoefficients(arr)
				degree:=len(coeffs)-1
				p=Poly1D{coeffs,degree,roots}
			case []float64:
				arr:=array.([]float64)
				for i,val := range arr{
					roots[i]=complex(val,0)
				}
				coeffs=PolyCoefficients(arr)
				degree:=len(coeffs)-1
				p=Poly1D{coeffs,degree,roots}
		}
	}
	return p
}
	
func (p Poly1D) Evaluate(x float64) (float64){
	l:=len(p.Coeffs)
	
	var sum float64
	sum=0
	for i:=0;i<l;i++{
		sum=sum+(p.Coeffs[i]*(float64(p.Degree-i)))
		
		}
	return sum
	}
func (p Poly1D) ViewPolynomial(){
	fmt.Printf("The roots of the polynomial:\n%v\n", p.Roots)
	fmt.Printf("The coefficients of the polynomial:\n%v\n", p.Coeffs)
	fmt.Printf("The polynomial is of degree:\n%d\n", p.Degree)
	
	if p.Degree<=2{
		fmt.Println("There is no companion matrix for degree 2 or below")
	}else{
		a := mat.NewDense(p.Degree, p.Degree, CompanionMatrix(p.Coeffs))
		fmt.Printf("The companion matrix of the polynomial is:\n = %v\n\n", mat.Formatted(a, mat.Prefix("    ")))
	}
}
func (p *Poly1D) PolyDifferentiate() *Poly1D{
	//This function differentaites the 1D polynomial
	coef:=make([]float64,len(p.Coeffs)-1)
	for i:=0;i<len(coef);i++{
		coef[i]=float64(p.Degree-i)*p.Coeffs[i]
	}
	p.Coeffs=coef
	p.Roots=PolyRoots(coef)
	p.Degree=len(coef)-1
	return p
}

func (p *Poly1D) NormalizeCoeff() *Poly1D{
	p.Coeffs=Normalize1DCoeffs(p.Coeffs)
	return p
}

// Expand returns the string representation of the polynomial x.
func (p Poly1D) Expand() string {
	if p.Degree == 0 {
		return "0"
	}

	s := ""
	index:=0
	for i := p.Degree; i > 1; i-- {
		if p.Degree+1&(1<<uint(i)) > 0 {
			if p.Coeffs[index]<0{
				s += fmt.Sprintf("%fx^%d",p.Coeffs[index], i)
			}else{
				s += fmt.Sprintf("+%fx^%d",p.Coeffs[index], i)
			}
			index++
		}
	}

	if p.Degree+1&2 > 0 {
		if p.Coeffs[index]<0{
			s += fmt.Sprintf("%fx",p.Coeffs[index])
		}else{
			s += fmt.Sprintf("+%fx",p.Coeffs[index])
		}
		index++
	}

	if p.Degree+1&1 > 0 {
		if p.Coeffs[index]<0{
			s += fmt.Sprintf("%f",p.Coeffs[index])
		}else{
			s += fmt.Sprintf("+%f",p.Coeffs[index])
		}
			
	}

	return s[1:]
}
/*
PolyExapand does (x-5)^5
*/

func (p *Poly1D) PolyExpand(n int) *Poly1D{
	if n<0{
		panic("Negative Values not handles")
	}else if n == 0{
		arr:=[]float64{1}
		p1:=NewPoly1D(arr,false)
		p.Coeffs=p1.Coeffs
		p.Degree=p1.Degree
		p.Roots=p1.Roots
	}else if n == 1{
		p=p
	}else if n == 2{
		p1:=PolyMul(p.Coeffs,p.Coeffs,"c")
		p.Coeffs=p1.Coeffs
		p.Degree=p1.Degree
		p.Roots=p1.Roots
	}else{
		p1:=PolyMul(p.Coeffs,p.Coeffs,"c")
		for i:=1;i<=n-2;i++{
			p1=PolyMul(p1.Coeffs,p.Coeffs,"c")
		}
		p.Coeffs=p1.Coeffs
		p.Degree=p1.Degree
		p.Roots=p1.Roots
	}
	return p

}
/*

func Random1DPoly(degree int) Poly1D{

}
*/
func Binary1DPoly(degree int) Poly1D{
	coeffs:=make([]float64, degree+1)
	for i:=0;i<=degree;i++{
		coeffs[i]=1
	}
	return NewPoly1D(coeffs, false)

}
/*
func Uniform1DPoly(n, degree int) Poly1D{

}

*/