package ringbuffer_test

import (
	"github.com/jtarchie/ringbuffer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sourcegraph/conc/iter"
)

var _ = Describe("Buffer", func() {
	It("can handle multiple readers and writers", func() {
		done := make(chan interface{})
		buffer := ringbuffer.NewChannel[int](5)

		go iter.ForEach([]int{1, 2, 3, 4, 5}, func(t *int) { buffer.Write(*t) })
		go func() {
			defer GinkgoRecover()

			results := iter.Map([]int{1, 2, 3, 4, 5}, func(*int) int { return buffer.Read() })
			Expect(results).To(HaveLen(5))
			Expect(results).To(ContainElements([]int{1, 2, 3, 4, 5}))
			close(done)
		}()

		Eventually(done).Should(BeClosed())
		buffer.Close()
	})

	It("favors more recent writes", func() {
		done := make(chan interface{})
		buffer := ringbuffer.NewChannel[int](5)

		for i := 1; i <= 100; i++ {
			buffer.Write(i)
		}

		go func() {
			defer GinkgoRecover()

			results := iter.Map([]int{1, 2, 3, 4, 5}, func(*int) int { return buffer.Read() })
			Expect(results).To(HaveLen(5))
			for _, value := range results {
				Expect(value).To(BeNumerically(">", 5))
			}
			close(done)
		}()

		Eventually(done).Should(BeClosed())
		buffer.Close()
	})
})
