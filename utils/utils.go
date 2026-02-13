package utils

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Assert(b bool, msg string) {
	if b == false {
		panic(msg)
	}
}
