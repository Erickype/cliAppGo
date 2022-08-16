package expenses

func Average(expens ...float32) float32 {
	var sum float32

	for _, expen := range expens {
		sum += expen
	}

	return sum / float32(len(expens))
}
