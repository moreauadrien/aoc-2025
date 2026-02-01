package main

import (
	"flag"
	"fmt"
	"github.com/moreauadrien/aoc-2025/days"
	"os"
	"path"
	"reflect"
)

func main() {
	examplePtr := flag.Bool("E", false, "run the day with example_dayXX.txt")
	dayPtr := flag.Uint("day", 1, "run the day")
	dirPtr := flag.String("dir", "./inputs", "directory of the inputs")

	flag.Parse()

	fn := fmt.Sprintf("day%02d.txt", *dayPtr)
	if *examplePtr {
		fn = fmt.Sprintf("example_%v", fn)
	}

	fp := path.Join(*dirPtr, fn)

	b, err := os.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	fc := string(b)

	mn := fmt.Sprintf("Day%02d", *dayPtr)

	meth := reflect.ValueOf(days.Days{}).MethodByName(mn)
	if meth.IsValid() == false {
		panic(fmt.Errorf("%v is not a valid method of 'Days'", mn))
	}

	args := []reflect.Value{reflect.ValueOf(fc)}
	ret := meth.Call(args)

	if len(ret) != 2 {
		panic(fmt.Errorf("%v should have two string as return values", mn))
	}

	sol1, ok1 := ret[0].Interface().(string)
	sol2, ok2 := ret[1].Interface().(string)

	if (ok1 == false) || (ok2 == false) {
		panic(fmt.Errorf("%v should have two string as return values", mn))
	}

	exampleStr := ""

	if *examplePtr {
		exampleStr = "(Example)"
	}

	fmt.Printf("--- DAY%02d%v ---\nPart1: %v\nPart2: %v\n", *dayPtr, exampleStr, sol1, sol2)
}
