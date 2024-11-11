package convert

// MapSlice применяет функцию fn к каждому элементу inputSlice, возвращает слайс без указателей
func MapSlice[in, out any](inputSlice []*in, fn func(*in) *out) []out {
	if inputSlice == nil {
		return nil
	}

	outputSlice := make([]out, 0, len(inputSlice))

	for _, item := range inputSlice {
		outputSlice = append(outputSlice, *fn(item))
	}

	return outputSlice
}
