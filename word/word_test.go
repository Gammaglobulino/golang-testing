package word

import (
	"fmt"
	"quotes"
	"testing"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("Expected", 3, "got", n)
	}
}
func TestUseCount(t *testing.T) {
	m := UseCount("one two three")
	m1 := make(map[string]int)
	if m["one"] != 1 && m["two"] != 1 && m["three"] != 1 {
		t.Error("Expected", m1, "got", m)
	}
}
func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quotes.Quote1)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quotes.Quote1)
	}
}
