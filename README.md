# Ring Buffer

This is a library for ring buffer operations in Go. A ring buffer, also known as
a circular buffer, is a fixed-size data structure that overwrites the oldest
data when it becomes full. You can learn more about ring buffers on
[Wikipedia](https://en.wikipedia.org/wiki/Circular_buffer). The primary goal of
this library is to provide a buffer that favors more recent data and prevents
back pressure on the buffer.

## Channel

This implementation uses Go channels as the underlying buffer. You can define
the size and type of the buffer.

Example:

```go
buffer := ringbuffer.NewChannel[int](1)
buffer.Write(1)
buffer.Write(2)

// in a goroutine far away
value := buffer.Read()
// value == 2
```

## Limitations

This library does not currently support:

- Emitting metrics when data is dropped
- Timeout/context for a buffer
