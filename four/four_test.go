package four_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/samcolson4/advent2015/four"
)

var _ = Describe("Four", func() {
	It("Calculates the hash number", func() {
		input := "abcdef"

		hashNumber, err := four.CalulcateHash(input)

		Expect(err).NotTo(HaveOccurred())
		Expect(hashNumber).To(Equal(609043))
	})

	It("Calculates another hash number", func() {
		input := "pqrstuv"

		hashNumber, err := four.CalulcateHash(input)

		Expect(err).NotTo(HaveOccurred())
		Expect(hashNumber).To(Equal(1048970))
	})
})
