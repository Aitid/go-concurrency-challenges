package fanout

func RoundRobinSplit[T any](inCh <-chan T, n int) []chan T {
	outChs := make([]chan T, n)

	if n == 0 {
		panic("RoundRobinSplit: n must be > 0")
	}

	for i := range n {
		outChs[i] = make(chan T)
	}

	go func() {
		idx := 0
		for in := range inCh {
			// deadlock if consumers are slow or absent
			outChs[idx] <- in

			// No guard for n <= 0
			idx = (idx + 1) % n // divide by zero
		}

		for _, outCh := range outChs {
			close(outCh)
		}
	}()

	return outChs
}
