package util

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func printObject(o Object, t *testing.T) {
	t.Log(o)
	o.DoSomething("printObject")
}
func TestAverage(t *testing.T) {

	fmt.Println("started")

	fmt.Printf("test")

	t.Log("Started")
	nums := []int{1, 2, 3, 4, 5, 6}
	for _, num := range nums {
		n := num
		go func(s string) {
			p := &Pod{"oleg" + strconv.Itoa(n)}
			t.Log("inside go routine")
			printObject(p, t)

			t.Log(s)
			t.Log(n)

		}("test")
	}
	time.Sleep(15 * time.Second)

}
