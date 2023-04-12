package ringbuffer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRingBuffer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RingBuffer Suite")
}
