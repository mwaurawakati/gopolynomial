package polynomial

import (
		
		"math"
		"fmt"
		//"math/rand"
		"os"
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




type Polinomial interface{
	Coefficients() []float64
	Roots() []float64
	Evaluate() float64
	}
	
	

type Poly struct{
	root []float64
	Coeffs []float64
	Variable string
	}
	

	
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


func FindArraySum(arr []float64) float64{
   var res float64
   res=0
   for i:=0; i<len(arr); i++ {
      res += arr[i]
   }
   return res
}

func ComplexArraySum(arr []complex128) complex128{
	var res complex128
	res = 0 +0i
	for i:=0; i<len(arr); i++{
		res +=arr[i]
	}
	return res
}
func FindArrayProduct(arr []float64) float64{
   var res float64
   res=1
   for i:=0; i<len(arr); i++ {
      res *= arr[i]
   }
   return res
}
func ComplexArrayProduct(arr []complex128) complex128{
	  var res complex128
    res=1
    for i:=0; i<len(arr); i++ {
      res *= arr[i]
    }
    return res
}
func Product2DArray(arr [][]float64) float64{
	var res float64
	res=1
	for i:=0; i<len(arr); i++ {
      fmt.Print(arr[i])
	}
	return res
}

func Binomial(n, k int) int {


	/* Binomial returns the binomial coefficient of (n,k), also commonly referred to
		as "n choose k".

		The binomial coefficient, C(n,k), is the number of unordered combinations of
		k elements in a set that is n elements big, and is defined as

		C(n,k) = n!/((n-k)!k!)

		n and k must be non-negative with n >= k, otherwise Binomial will panic.
		No check is made for overflow.
	*/
	
	if n < 0 || k < 0 {
		panic(errNegInput)
	}
	if n < k {
		panic(badSetSize)
	}
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

func UniqueCombinations(s interface{}, k int) [][]float64{
	switch s.(type) {
		case int:
			x, _ := s.(int)
			return intcombinations(x,k)
		case float64:
			x, _ := s.(float64)
			return intcombinations(int(x), k)
		
		case string:
			fmt.Printf("Input can only be a slice or an integer")
		case []float64:
			x, _ := s.([]float64)
			return slicecombinations(x,k)
		}
	r:=make([][]float64,0)
	return r
}

func cUniqueCombinations(s interface{}, k int) [][]complex128{
	switch s.(type){

		case []complex128:
			x, _ := s.([]complex128)
			return ComplexUniqueCombinations(x,k)
		case []complex64:
			x, _ := s.([]complex64)
			return ComplexUniqueCombinations(x,k)
		
	}
	r:=make([][]complex128,0)
	return r
}

func intcombinations(n, k int) [][]float64 {
	combins := Binomial(n, k)
	data := make([][]float64, combins)
	if len(data) == 0 {
		return data
	}
	data[0] = make([]float64, k)
	for i := range data[0] {
		data[0][i] = float64(i)
	}
	for i := 1; i < combins; i++ {
		next := make([]float64, k)
		copy(next, data[i-1])
		nextCombination(next, n, k)
		data[i] = next
	}
	return data
}

// nextCombination generates the combination after s, overwriting the input value.
func nextCombination(s []float64, n, k int) {
	for j := k - 1; j >= 0; j-- {
		if float64(s[j]) == float64(n)+float64(j)-float64(k) {
			continue
		}
		s[j]++
		for l := j + 1; l < k; l++ {
			s[l] = float64(s[j]) + float64(l) - float64(j)
		}
		break
	}
}
func slicecombinations(s []float64, k int) ([][]float64){
	combins:=Binomial(len(s),k)
	data:=make([][]float64,combins)
	if len(data) == 0 {
		return data
	}
	combs:=intcombinations(len(s),k)
	
	arr:=make([]float64,len(s))
	for i:=0;i<len(s);i++{
		arr[i]=float64(i)
	}
	for i:=0;i<len(combs);i++{
		c:=combs[i]
		
		var intSlice intSlice
		intSlice=arr
		index:=[]int{}
		for j:=0;j<len(c);j++{
			ind:=intSlice.pos(c[j])
			if ind !=-1{
				index=append(index,ind)
			}
		}
		
		com:=make([]float64,len(index))
		for j:=0;j<len(index);j++{
			p:=index[j]
			com[j]=s[p]
		}
		
		data[i] = com
	}
	return data
}

func ComplexUniqueCombinations(s interface{}, k int) ([][]complex128){
	s1 := []complex128{}
	switch s.(type){
		case []complex128:
			x, _ :=s.([]complex128)
			s1=x
		case []complex64:
			//x, _ :=s.([]complex64)
			//for i,val :=range x{
			//	r:=floareal(val)
			//	im:=imag(val)
			//	num:=complex(r,im)
			//	s1[i]=num
			//}
			panic(errNegInput)
			os.Exit(0)
	}
	
	combins:=Binomial(len(s1),k)
	data:=make([][]complex128,combins)
	if len(data) == 0 {
		return data
	}
	combs:=intcombinations(len(s1),k)
	
	arr:=make([]float64,len(s1))
	for i:=0;i<len(s1);i++{
		arr[i]=float64(i)
	}
	for i:=0;i<len(combs);i++{
		c:=combs[i]
		
		var intSlice intSlice
		intSlice=arr
		index:=[]int{}
		for j:=0;j<len(c);j++{
			ind:=intSlice.pos(c[j])
			if ind !=-1{
				index=append(index,ind)
			}
		}
		
		com:=make([]complex128,len(index))
		for j:=0;j<len(index);j++{
			p:=index[j]
			com[j]=s1[p]
		}
		
		data[i] = com
	}
	return data
}	

func Normalize1DCoeffs(arr []float64) []float64{
	if arr[0] != 1{
		num:=arr[0]
		for i,val := range arr{
			arr[i]=val/num
		}
	}
	return arr
}
type intSlice []float64

func (slice intSlice) pos(value float64) int {
    for p, v := range slice {
        if (v == value) {
            return p
        }
    }
    return -1
}


func PolyMul(array1, array2 []float64, r bool) []float64{
	/*
		This function multiplies two polinomials;A and Binomial
		params:
		array1: a 1D array of the first polinomial's roots or coefficients
		array2: a 1D array of the 2nd polinomial's roots of coefficients
		r:specifies whether array1 and array2 are roots. they are roots if r==true
		returns:
		The resulting product
	*/
	if r==true{
		array1=PolyCoefficients(array1)
		array2=PolyCoefficients(array2)
	}else{
		array1=array1
		array2=array2
		}
	m := len(array1)
	n := len(array2)
	prod :=make([]float64, m+n-1)
	for i := 0; i< m+n-1; i++{
		prod[i]=float64(0)
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			prod[i+j] +=array1[i]*array2[j]
		}
	}
	return prod
}

func PolyDiv(array1,array2 []float64, r bool) []float64{
	return array1
}

func guesssolution(array []float64, x float64) float64{
	
	n := len(array)
	p :=float64(1)//array[n]
	for i :=n;i>=1;i--{
		p=array[i-1]+(x*p)
	}
	return p

}

func CompanionMatrix(coeffs []float64) []float64{
	if len(coeffs)<=2{
		fmt.Println("A companion matrix can not be created."+
		"This function only works with quadratic equations or polinomials of higher degree")
		panic(errWrongPolinomial)
	}
	p:=len(coeffs)-1
	comp := make([]float64, int(math.Pow(float64(p),float64(2))))
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
	return comp
}
	
	
func CompanionOneSeries(n int) []int{
	if n<2{
		panic(errWrongPolinomial)
	}
	index := make([]int, n-1)
	
	//index[0]=1
	k:=1
	for i:=0;i<len(index);i++{
	
		index[i]=k
		k=k+n+1
	}
	return index
}
	
func Product(a []float64, r int) func() []float64 {
	/*This the cartesian product of input iterables. Its python equivalent is
	  itertools.product() function(https://docs.python.org/3/library/itertools.html)*/
	p := make([]float64, r)
	x := make([]int, len(p))
	return func() []float64 {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

	