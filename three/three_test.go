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

	It("Accepts multiple symbols at a time", func() {
		coords, err := three.SetCoords("^>v<")

		Expect(err).NotTo(HaveOccurred())
		Expect(coords[0]).To(Equal([]int{0, 0}))
		Expect(coords[1]).To(Equal([]int{0, 1}))
		Expect(coords[2]).To(Equal([]int{1, 1}))
		Expect(coords[3]).To(Equal([]int{1, 0}))
		Expect(coords[4]).To(Equal([]int{0, 0}))
	})

	It("Calculates how many time a house has been visited", func() {
		coords, err := three.SetCoords("^>v<")
		Expect(err).NotTo(HaveOccurred())

		coordStructs := three.MakeCoords(coords)

		housesVisited, err := three.CalculateHouses(coordStructs)

		Expect(err).NotTo(HaveOccurred())
		Expect(housesVisited).To(Equal(4))
	})

	It("Calculates how many time a house has been visited", func() {
		coords, err := three.SetCoords("^v^v^v^v^v")
		Expect(err).NotTo(HaveOccurred())

		coordStructs := three.MakeCoords(coords)

		housesVisited, err := three.CalculateHouses(coordStructs)

		Expect(err).NotTo(HaveOccurred())
		Expect(housesVisited).To(Equal(2))
	})
})

var _ = Describe("RoboThree", func() {
	It("Creates coords for Santa & roboSanta", func() {
		SantaCoords, RoboCoords, err := three.SplitInput("^v")
		Expect(err).NotTo(HaveOccurred())
		Expect(SantaCoords).To(Equal("^"))
		Expect(RoboCoords).To(Equal("v"))
	})
})
