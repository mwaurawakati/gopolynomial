package gopolynomial

import (
		"os"
		"fmt"
		"math"
		)
const(
	version = "v 0.1.0"
	author	= "Mwaura Wakati"
	licence	= ""
	about	= "gopolynomial is an open source Golang library for polynomial\n"+
				"for more information visit https://github.com/mwaurawakati/gopolinomial"
	)
//Given a slice of complex variables, cUniqueCombinations will 
//return all the complex unique combinations

func cUniqueCombinations(s interface{}, k int) [][]complex128{
	switch s.(type){

		case []complex128:
			x, _ := s.([]complex128)
			return complexUniqueCombinations(x,k)
		case []complex64:
			x, _ := s.([]complex64)
			return complexUniqueCombinations(x,k)
		
	}
	r:=make([][]complex128,0)
	return r
}

//intcombinations gives gives unique combinations once given n, the number of variables, and k the length of
//combinations 

func intcombinations(n, k int) [][]float64 {
	combins := binomial(n, k)
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

//slice combinations is similar to int combinations only that you are specifing the array from
//which the combination variables are chosen from

func slicecombinations(s []float64, k int) ([][]float64){
	combins:=binomial(len(s),k)
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

//FindArraySum return the sum of array variables

func findArraySum(arr []float64) float64{
   var res float64
   res=0
   for i:=0; i<len(arr); i++ {
      res += arr[i]
   }
   return res
}


func complexArraySum(arr []complex128) complex128{
	var res complex128
	res = 0 +0i
	for i:=0; i<len(arr); i++{
		res +=arr[i]
	}
	return res
}
func findArrayProduct(arr []float64) float64{
   var res float64
   res=1
   for i:=0; i<len(arr); i++ {
      res *= arr[i]
   }
   return res
}
func complexArrayProduct(arr []complex128) complex128{
	  var res complex128
    res=1
    for i:=0; i<len(arr); i++ {
      res *= arr[i]
    }
    return res
}
func product2DArray(arr [][]float64) float64{
	var res float64
	res=1
	for i:=0; i<len(arr); i++ {
      fmt.Print(arr[i])
	}
	return res
}

func binomial(n, k int) int {


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

func uniqueCombinations(s interface{}, k int) [][]float64{
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



func complexUniqueCombinations(s interface{}, k int) ([][]complex128){
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
	
	combins:=binomial(len(s1),k)
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

func normalize1DCoeffs(arr []float64) []float64{
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
/*When creating a compaion matrix e.g:
		[0 1 0 0]
		[0 0 1 0]
		[0 0 0 1]
		[-c0 -c1 -c2 -c3]
  CombinationOneSeries helps to come up with posion of ones in the companion matrix
*/
		
func companionOneSeries(n int) []int{
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

func polymul(array1, array2 []float64) []float64{
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

func vversion(){
	fmt.Print(version)
	}
func aabout(){
	fmt.Println("\nThis is gopolinomial version: ", version)
	fmt.Println("Author                      : ", author)
	fmt.Println("Licence                     : ", licence)
	fmt.Println(about,"\n")
}


//Linspace is a utility function similar to python's numpy linspace
//Given the start, stop, the number of values to be retuned and whether
//the endpoint should be included or not
func Linspace(start,stop float64,num int,endpoint bool) []float64{
	var total float64
	if start<0{
		total=math.Abs(start)+stop
	}else{
		total=stop-start
	}
	nums:=make([]float64,num)
	if endpoint==false{
		step:=total/float64(num)
		for i,_ :=range nums{
			nums[i]=start+step*float64(i)
		}
		nums[0]=start
		return nums

	}else{
		step:=total/float64(num-1)
		for i,_ :=range nums{
			nums[i]=start+step*float64(i)
		}
		nums[0]=start
		return nums
	}
}
