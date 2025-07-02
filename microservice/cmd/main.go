package main

// func main() {
// 	file, err := os.Open("go.mod")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close()

// 	reader := io.NewFullReader(file, 4096)
// 	// reader.Close()

// 	// file.Read()
// 	// file.Write()

// 	file2, err := os.Open("go.mod")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// defer file2.Close()

// }

// type DatabaseAccessor struct {
// }

// func NewdatabseAccessor() (accessor *DatabaseAccessor, cleanupFunc func(), error error) {
// 	return &DatabaseAccessor{}, func() {
// 		// clean task
// 	}, nil
// }

// func main() {
// 	if err := recover(); err != nil {
// 		// handle exception
// 	}

// 	accessor, cleanupFunc, err := NewdatabseAccessor()

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cleanupFunc()
// }
