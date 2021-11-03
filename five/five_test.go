package five_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/samcolson4/advent2015/five"
)

var _ = Describe("Five pt.1", func() {
	It("Is nice because it has three vowels", func() {
		input := "ugknbfddgicrmopn"

		assert := five.CheckVowels(input)

		Expect(assert).To(Equal(true))
	})

	It("Is nasty because it does not have three vowels", func() {
		input := "yhjplkqm"

		assert := five.CheckVowels(input)

		Expect(assert).To(Equal(false))
	})

	It("Is nice because it has a double letter", func() {
		input := "ugknbfddgicrmopn"

		assert := five.CheckDoubles(input)

		Expect(assert).To(Equal(true))
	})

	It("Is nasty because it does not have a double letter", func() {
		input := "yhjplkqm"

		assert := five.CheckDoubles(input)

		Expect(assert).To(Equal(false))
	})

	It("Is nice because it does not contain a banned string", func() {
		input := "ugknbfddgicrmopn"

		assert := five.CheckBanned(input)

		Expect(assert).To(Equal(true))
	})

	It("Is nasty because it contains banned strings", func() {
		input := "adcdpqxy"

		assert := five.CheckBanned(input)

		Expect(assert).To(Equal(false))
	})

	It("Checks for all issues", func() {
		input := "jchzalrnumimnmhp"

		assert := five.CheckString(input)

		Expect(assert).To(Equal(false))

	})
})

var _ = Describe("Five pt.2", func() {
	It("Is nice because it has a pair repeating twice", func() {
		input := "xxyxx"

		assert := five.TwoCheckPairs(input)

		Expect(assert).To(Equal(true))
	})

	It("Is naughty because it has an overlapping pair", func() {
		input := "aaa"

		assert := five.TwoCheckPairs(input)

		Expect(assert).To(Equal(false))
	})

	It("Is naughty because it has a pair (tg) but no repeat with a single letter between them", func() {
		input := "uurcxstgmygtbstg"

		assert := five.TwoCheckPairs(input)

		Expect(assert).To(Equal(false))
	})
})
