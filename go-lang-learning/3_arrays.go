package main
 
import "fmt"
 
func main() {
    // define array named domains as string type
    var domains [2]string
    fmt.Println("current values for array:", domains)
 
    // add value and print it
    domains[0] = "cyberciti.biz"
    fmt.Println("Set value : ", domains)
    fmt.Println("Get value for 0 element : ", domains[0])
    // get array length 
    fmt.Println("Array length : ", len(domains))
    // add one more value 
    domains[1] = "nixcraft.com"
    // use for loop to print our array
    for i := 0; i < len(domains); i++ {
	    fmt.Println("Get value for element ", i, " is ", domains[i])
    }
}