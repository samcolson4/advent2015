package noMath_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Two Suite")
}
