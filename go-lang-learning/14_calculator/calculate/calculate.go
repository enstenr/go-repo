package calculate

func Calculate(num1 , num2 int , action string) int64{
	returnval:=0
	 
	switch action {
	case "ADD":
		returnval= num1 + num2
	case "SUB":
		returnval= num1 - num2

	}
	return int64(returnval)
}
