package main

import "fmt"

func UpdateProfile(id string) error {
	return fmt.Errorf(" failed to update profile ")
}
func main() {
	fmt.Println(UpdateProfile("user1"))
}
