package noMath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	noMath "github.com/samcolson4/advent2015/two"
)

var _ = Describe("noMath", func() {
	It("strips text and returns three ints", func() {
		input := "3x11x24"
		measure, err := noMath.StripText(input)

		Expect(err).NotTo(HaveOccurred())
		Expect(measure[0]).To(Equal(3))
		Expect(measure[1]).To(Equal(11))
		Expect(measure[2]).To(Equal(24))
	})
})
