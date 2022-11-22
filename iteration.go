// i followed this tutorial on github - https://github.com/quii/learn-go-with-tests/blob/main/iteration.md
package main

func Repeat(x string) string {

	var repeated string

	for i := 0; i < 4; i++ {
		repeated += x
	}
	return repeated
}
