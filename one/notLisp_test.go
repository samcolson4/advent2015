package notLisp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/samcolson4/advent2015/one/notLisp"
)

var _ = Describe("notLisp", func() {
	It("returns 0 when passed '(())'", func() {
		input := "(())"
		output, err := notLisp.Calculate(input)

		Expect(err).NotTo(HaveOccurred())
		Expect(output).To(Equal(0))
	})
})
