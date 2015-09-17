package main
/*
func main() {

var number int64
fmt.Println("Please enter a number (n>0) to find factorial")

fmt.Scanf("%d", &number)


factorial := Fibo(number)

fmt.Println("Factorial Calculated is :",factorial)
}*/

func Fibo(input int64) (ret int64) {

if(input<0){
	return -1 //FAIL
}

if input == 0 {
return 0
}else if input == 1 {
return 1
}else {
return Fibo(input-1)+Fibo(input-2)
}


}
