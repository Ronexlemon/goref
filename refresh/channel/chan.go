package channel



type Request struct {
	Num []int64
}

// Counter reads requests and sends sum of each request
func Counter(receiving <-chan Request, send chan<- int64) {
	for rec := range receiving {
		var sum int64

		for _, v := range rec.Num {
			sum += v
		}

		send <- sum
	}

	close(send)
}