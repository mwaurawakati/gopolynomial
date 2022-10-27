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
	fmt.Printf("The roots of the polynomial:\n%v\n", p.Coeffs)
	fmt.Printf("The polynomial is of degree:\n%d\n", p.Degree)
	a := mat.NewDense(p.Degree, p.Degree, CompanionMatrix(p.Coeffs))
	fmt.Printf("The companion matrix of the polynomial is:\n = %v\n\n", mat.Formatted(a, mat.Prefix("    ")))
	}