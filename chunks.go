package chunks

// Chunk will split a slice up into chunks of a given size
func Chunk[T any](s []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(s); i += size {
		end := i + size

		if end > len(s) {
			end = len(s)
		}

		chunks = append(chunks, s[i:end])
	}

	return chunks
}
