//Coded by Suparna

// Program to enter any number and find the table in GOLANG

// Thanks for watching

package main

import "fmt"

func main() {
    
    var i int = 1
    
    fmt.Scanln(&i)
    
    fmt.Println("Table of " , i , "\n")
    
    for j:= 1; j<=12; j++ {
        var sum int = i * j
        
        fmt.Println(i, "×" ,j , " ="  , sum)
    }
     
}
