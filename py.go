package main
import "fmt"
func findsplpytrip(n int)(int,int,int){
	for a:=1;a<n/3;a++{
		for b:=a;b<(n-a)/2;b++{
			c:=n-a-b
			if a*a+b*b==c*c{
				return a,b,c
			}
		}
		
	}
	return -1,-1,-1
}
func main(){
	var num int
	fmt.Print("enter the sum")
	fmt.Scan(&num)
	a,b,c:=findsplpytrip(num)
	if a!=-1&&b!=-1&&c!=-1{
		fmt.Printf("The py triplet of %d is(%d,%d,%d)",num,a,b,c)
	}else{
		fmt.Print("Not found")
	}

	}
