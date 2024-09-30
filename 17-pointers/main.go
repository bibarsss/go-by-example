package main

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func updateName(s *string) {
	*s = "biba"
}

func main() {
	//s := "ute"
	//
	//fmt.Println("Memory address of s is: ", &s)
	//
	////m := &s
	//updateName(&s)
	//
	//fmt.Println(s)

	//i := 1
	//fmt.Println("initial:", i)
	//zeroval(i)
	//fmt.Println("zeroval:", i)
	//
	//zeroptr(&i)
	//fmt.Println("zeroptr:", i)
	//fmt.Println("pointer:", &i)
}
