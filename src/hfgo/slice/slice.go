package main

import "fmt"

func main() {
	s0 := []string{}
	s1 := append(s0, "s1", "s1")
	//s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	s5 := append(s4, "s5", "s5")
	fmt.Println(s1, cap(s1),
		s2, cap(s2),
		s3, cap(s3),
		s4, cap(s4),
		s5, cap(s5),
	)
	s4[0] = "XX"
	fmt.Println(s1, cap(s1),
		s2, cap(s2),
		s3, cap(s3),
		s4, cap(s4),
		s5, cap(s5),
	)
	s2[1] = "YY"
	fmt.Println(s1, cap(s1),
		s2, cap(s2),
		s3, cap(s3),
		s4, cap(s4),
		s5, cap(s5),
	)
}
