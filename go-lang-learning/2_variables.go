// Hello Word in Go by Vivek Gite
package main
 
// Import OS and fmt packages
import ( 
	"fmt" 

)
 
func main() {
var name string 
_=name
var name1 = 20
//shorthand syntax 
name2 := "S Rajesh"
 
// define i
var i int
// set value for i
i = 10
// we can also set value as follows
var y = 5
var msg = "Remote host found."
var foo, bar int = 100, 200
// shorthand syntax
vehicle := "Mercedes"
age := 52
// Bool true or false
var is_job_failed = false
// print it
fmt.Printf("%d %d %s\n", i,y,msg)
fmt.Println(foo)
fmt.Println(age)
fmt.Println(vehicle)
fmt.Println(name1,name2,bar,is_job_failed)

m := 1
for m <= 5 {
	fmt.Printf("Welcome %d times.\n",m)
	m = m + 1
}
// classic for loop example 
for i := 6; i <= 10; i++ {
	fmt.Printf("Welcome %d times.\n",i)
}
} 