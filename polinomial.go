package gopolynomial

import (
		
		"math"
		"fmt"
		"strings"
		"gonum.org/v1/gonum/mat"
		)

const (
	errNegInput             = "combin: negative input"
	badSetSize              = "combin: n < k"
	badInput                = "combin: wrong input slice length"
	errNonpositiveDimension = "combin: non-positive dimension"
	errWrongPolinomial		= "companion: works with order two and above"
	errZero					= "Polynomial: The array can not be empty"
)

	
func PolyCoefficients(roots interface{}) ([]float64){

	/*This function returns the coefficients given the roots of the polinomial
	inputs:
	roots:An array of roots eg[1,2,3] which repreents(x-1)(x-2)(x-3)
			The input might also be an array made up of complex128 variables eg.{1+1i, 2+2i}
	outputs
	an array of the polynomial's coefficients
	*/
	s :=roots
	c := []float64{}
	switch s.(type) {
		case []float64:
			x, _ := s.([]float64)
		
			l:=len(x)
			c =make([]float64, l+1)
			c[0]=1
			var coef float64
	
			for i:=1;i<=l;i++{
				coef =float64(0)
				p:=UniqueCombinations(x, i)
				for j:=0;j<len(p);j++{
					arr:=p[j]
					prod:=FindArrayProduct(arr)
					coef += prod
				}
			
		
	
	
				c[i]=math.Pow(float64(-1),float64(i))*coef
		
			}
			
		case []complex128:
			x, _ := s.([]complex128)
		
			l:=len(x)
			c =make([]float64, l+1)
			c[0]=1
			var coef complex128
	
			for i:=1;i<=l;i++{
				coef =0
				p:=cUniqueCombinations(x, i)
				for j:=0;j<len(p);j++{
					arr:=p[j]
					prod:=ComplexArrayProduct(arr)
					coef += prod
				}
			
		
	
	
				c[i]=math.Pow(float64(-1),float64(i))*real(coef)
		
			}
			
		case []complex64:
			x, _ := s.([]complex64)
		
			l:=len(x)
			c =make([]float64, l+1)
			c[0]=1
			var coef complex128
	
			for i:=1;i<=l;i++{
				coef =0
				p:=cUniqueCombinations(x, i)
				for j:=0;j<len(p);j++{
					arr:=p[j]
					prod:=ComplexArrayProduct(arr)
					coef += prod
				}
			
		
	
	
				c[i]=math.Pow(float64(-1),float64(i))*real(coef)
		
			}	
	}
	return c
}
func PolyRoots(coeffs []float64) []complex128{
	/*
	This function returns the roots given the coefficients
	*/

	if len(coeffs)==0{
		return []complex128{0+0i}
	}
	if len(coeffs)==1{
		return []complex128{complex((-1)*coeffs[0],0)}
	}
	
	
	c:=CompanionMatrix(coeffs)
	a := mat.NewDense(len(coeffs)-1,len(coeffs)-1,c)
	var eig mat.Eigen
	eig.Factorize(a, mat.EigenLeft)
	return eig.Values(nil)
}




func PolyMul(array1, array2 interface{}, t string) Poly1D{
	/*
		This function multiplies two polinomials;A and B
		params:
		array1: a 1D array of the first polinomial's roots or coefficients, or Poly1D
		array2: a 1D array of the 2nd polinomial's roots of coefficients, or PolY1D
		r:specifies the type of input, i.e p/poly1D/poly or r/roots or c/coef/coefficients
		returns:
		The resulting poly1D
	*/
	var poly Poly1D
	t=strings.ToLower(t)
	if t=="c"||t=="coef"||t=="coeffs"||t=="coefficients"||t=="coefficient"{
		a1:=array1.([]float64)
		a2:=array2.([]float64)
		poly=NewPoly1D(polymul(a1,a2),false)
		
	}else if t=="roots"||t=="r"||t=="root"{
		a1:=PolyCoefficients(array1)
		a2:=PolyCoefficients(array2)
		poly=NewPoly1D(polymul(a1,a2),false)
	}else if t=="p"||t=="poly"||t=="poly1d"{
		a1:=array1.(Poly1D)
		a2:=array2.(Poly1D)
		poly=NewPoly1D(polymul(a1.Coeffs,a2.Coeffs),false)
	}else{
		panic(errWrongPolinomial)
	}
	//poly.Coeffs=Normalize1DCoeffs(poly.Coeffs)
	return poly
	
		
}

func PolyDiv(array1,array2 []float64) ([]float64,[]float64){
	if len(array1)  == 0 {
		panic("division by zero")
	}

	var arr1,arr2 []float64
	if len(array1)<len(array2){
		arr1 = arr1
		arr2 = array2
	}else if len(array1)==len(array2){
		arr1=make([]float64, 1)
		arr2=make([]float64, len(array2))
		if array1[0] != array2[0]{
			if array1[0]>array2[0]{
				arr1[0]=array1[0]
			}else{
				arr1[0]=float64(1)/array2[0]
			}
		}else{
			arr1[0]=1
		}
		for k:=0;k<len(array2);k++{
			array2[k]=array2[k]*arr1[0]
			arr2[k]=array1[k]-array2[k]
		}
	}else{
		diff:=len(array1)-len(array2)
		arr1=make([]float64, (diff+1))
		index:=0
		ar1:=array1
		for i:=diff;i>=0;i--{
			
			ar2:=make([]float64, len(ar1))
			fac:=ar1[0]/array2[0]
			for k:=0;k<len(array2);k++{
				ar2[k]=fac*array2[k]
			}
			for k:=len(array2);k<len(ar2);k++{
				ar2[k]=0
			}
			
			ar11:=ar1
			for k,_ := range ar1{
				ar11[k]=ar1[k]-ar2[k]
			}
			arr1[index]=fac
			index++
			ar1=make([]float64,len(ar11)-1)
			
			for k:=1;k<len(ar11);k++{
				ar1[k-1]=ar11[k]
			}
			
		}
		arr2=ar1
		
	}	
		
	
	return arr1, arr2
}



func CompanionMatrix(coeffs []float64) []float64{
	var comp []float64
	if len(coeffs)<=2{
		fmt.Println("A companion matrix can not be created."+
		"This function only works with quadratic equations or polinomials of higher degree")
		//panic(errWrongPolinomial)
		comp=comp
	/*}
	if len(coeffs)==2{
		c:=make([]float64, 4)
		c[0]=0
		c[1]=0
		c[2]=coeffs[1]
		c[3]=coeffs[2]
		comp=c*/
	}else{
	p:=len(coeffs)-1
	comp = make([]float64, int(math.Pow(float64(p),float64(2))))
	for i:=0;i<len(comp);i++{
		comp[i] =0 
		}
	i :=len(comp)-len(coeffs)+1
	for j:=len(coeffs)-1;j>0;j--{
		comp[i]=float64(-1)*coeffs[j]
		i=i+1
	}
	index := CompanionOneSeries(len(coeffs)-1)
	for _, val := range index{
		comp[val]=float64(1)
	}
	}
	return comp
}
	
	



	