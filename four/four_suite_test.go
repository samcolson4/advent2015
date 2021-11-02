package four_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFour(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Four Suite")
}
