package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var i, n, j, k, elemnHilang int

	fmt.Scan(&n)
	for i = 0; i<n; i++{
		fmt.Scan(&A[i])
		j = A[i]
	}
	for i = 0; i < n; i++{
		if j < A[i]{
			j = A[i]
		}
	}

	for i = 1; i <= j; i++{
		k = 0
		for k = 0; k < n; k++{
			if i == A[k]{
				break
			}
		}
		if k == n{
			elemnHilang = i
			fmt.Println(elemnHilang)
		}
	}
}