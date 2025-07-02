package main

// duck typing
// build hermetic
// go mod vender
type CodeError struct {
}

func (e *CodeError) Error() string {
	return ""
}

func main() {
	var _ error = &CodeError{}
}
