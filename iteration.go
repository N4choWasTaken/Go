package main

func Repeat(x string) string {

	var repeated string

	for i := 0; i < 4; i++ {
		repeated += x
	}
	return repeated
}
