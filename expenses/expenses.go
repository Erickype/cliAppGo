package expenses

func Average(expens ...float32) float32 {

	return Sum(expens...) / float32(len(expens))
}

func Sum(expens ...float32) float32 {
	var sum float32
	for _, expen := range expens {
		sum += expen
	}
	return sum
}

func Max(expens ...float32) float32 {

	var max float32

	for _, expen := range expens {
		if expen > max {
			max = expen
		}
	}

	return Sum(expens...) / float32(len(expens))
}

func Min(expens ...float32) float32 {

	if len(expens) == 0 {
		return 0
	}

	var min = expens[0]

	for _, expen := range expens {
		if expen < min {
			min = expen
		}
	}

	return Sum(expens...) / float32(len(expens))
}
