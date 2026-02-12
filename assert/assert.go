package assert

func Assert(truth bool, msg string) {
	if !truth {
		panic(msg)
	}
}

