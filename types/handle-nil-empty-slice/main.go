package handlenilemptyslice

func handleOperationsWrong(op []float64) {
	// if the slice is empty but not nil, handle will execute and ur kinda fucked
	if op != nil {
		handle(op)
	}
}

func handleOperationsRight(op []float64) {
	// if the slice is nil then len != 0 is false, if slice is only empty then len also = 0, ur good
	if len(op) != 0 {
		handle(op)
	}
}

func handle(op []float64) {}
