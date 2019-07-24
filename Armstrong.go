package main
import (
	"fmt"
	"math"
)
func main(){
	var upper,lower,n1,n2,n3 int
	var power=0
	var sum=0
	
	fmt.Println("Enter the lower limit")
	fmt.Scan(&lower)
	fmt.Println("Enter the upper limit")
	fmt.Scan(&upper)
	for i:=lower;i<=upper;i++{
		n1=i
		n2=n1
		n3=n1
		for{
			if n1==0{
				break
			}else{
				power++
				n1=n1/10
			}
		}
		for{
			if n2==0{
				break
			}else
			{
				calc:=n2%10
				sum+= int(math.Pow(float64(calc),float64(power)))
				n2=n2/10
			}
		}
		if(n3==sum){
			fmt.Println("Armstrong Number: ")
			fmt.Println(n3)
		}
		sum=0;
		power=0;

	}
}