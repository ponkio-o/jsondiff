package jsondiff

import (
	"fmt"
	"os"

	"github.com/itchyny/gojq"
)

var (
	lhs = map[string]interface{}{"a": 1, "b": 2, "c": 3, "d": 4}
	rhs = map[string]interface{}{"a": 1, "b": 1, "c": 2, "d": 3}
)

func ExampleDiffFromFiles() {
	query, err := gojq.Parse(".d")
	if err != nil {
		panic(err)
	}
	from, err := os.Open("./testdata/from.json")
	if err != nil {
		panic(err)
	}
	to, err := os.Open("./testdata/to.json")
	if err != nil {
		panic(err)
	}
	diff, err := DiffFromFiles(from, to, Only(query))
	if err != nil {
		panic(err)
	}
	fmt.Print(diff)
	// Output:
	// --- from.json
	// +++ to.json
	// @@ -1,2 +1,2 @@
	// -4
	// +3
}

func ExampleDiffFromObjects_only() {
	query, err := gojq.Parse(".d")
	if err != nil {
		panic(err)
	}
	diff, err := DiffFromObjects(lhs, rhs, Only(query))
	if err != nil {
		panic(err)
	}
	fmt.Println(diff)
	// Output:
	// --- from
	// +++ to
	// @@ -1,2 +1,2 @@
	// -4
	// +3
}

func ExampleDiffFromObjects_ignore() {
	query, err := gojq.Parse(".b, .c")
	if err != nil {
		panic(err)
	}
	diff, err := DiffFromObjects(lhs, rhs, Ignore(query))
	if err != nil {
		panic(err)
	}
	fmt.Println(diff)
	// Output:
	// --- from
	// +++ to
	// @@ -1,5 +1,5 @@
	//  {
	//    "a": 1,
	// -  "d": 4
	// +  "d": 3
	//  }
}
