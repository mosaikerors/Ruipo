
package main
 
import "fmt"
import "math/rand"
import "time"
 
func Generate_Randnum() int{
    rand.Seed(time.Now().Unix())
    rnd := rand.Intn(100)
 
    fmt.Printf("rand is %v\n", rnd)
 
    return rnd
} 
 
func main(){
	rand.Seed(time.Now().Unix())
	for i:=0;i<10;i++{
		fmt.Printf("rand: %s\n", generateShortLink(6))
	}
    
}