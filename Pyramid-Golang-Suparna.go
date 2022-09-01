package main
import "fmt"

func main()
{
var number int;
fmt.Println("Enter any number:");
fmt.Scanln(&number);

if number % 2 == 0
{
fmt.Println("\nThe number entered is even");
}

else 
{
fmt.Println("\nThe number entered is odd");
}


}
