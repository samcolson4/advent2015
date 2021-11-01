package three_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	three "github.com/samcolson4/advent2015/three"
)

var _ = Describe("Three", func() {

	It("Works out one output based on coordinate: '>'", func() {
		coords, err := three.SetCoords(">")

		Expect(err).NotTo(HaveOccurred())
		Expect(coords[0]).To(Equal([]int{0, 0}))
		Expect(coords[1]).To(Equal([]int{1, 0}))
	})

	It("Works out one output based on coordinate: '^'", func() {
		coords, err := three.SetCoords("^")

		Expect(err).NotTo(HaveOccurred())
		Expect(coords[0]).To(Equal([]int{0, 0}))
		Expect(coords[1]).To(Equal([]int{0, 1}))
	})

	It("Works out one output based on coordinate: 'v'", func() {
		coords, err := three.SetCoords("v")

		Expect(err).NotTo(HaveOccurred())
		Expect(coords[0]).To(Equal([]int{0, 0}))
		Expect(coords[1]).To(Equal([]int{0, -1}))
	})

	It("Works out one output based on coordinate: '<'", func() {
		coords, err := three.SetCoords("<")

		Expect(err).NotTo(HaveOccurred())
		Expect(coords[0]).To(Equal([]int{0, 0}))
		Expect(coords[1]).To(Equal([]int{-1, 0}))
	})
})
