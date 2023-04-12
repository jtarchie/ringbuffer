package ringbuffer

type Channel[T any] struct {
	input  chan T
	output chan T
}

func NewChannel[T any](size int) *Channel[T] {
	b := &Channel[T]{
		input:  make(chan T),
		output: make(chan T, size),
	}
	go b.run()

	return b
}

func (b *Channel[T]) run() {
	for v := range b.input {
	retry:
		select {
		case b.output <- v:
		default:
			<-b.output

			goto retry
		}
	}

	close(b.output)
}

func (b *Channel[T]) Write(value T) {
	b.input <- value
}

func (b *Channel[T]) Read() T {
	return <-b.output
}

func (b *Channel[T]) Close() {
	close(b.input)
}
