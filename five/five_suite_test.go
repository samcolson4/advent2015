package five_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFive(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Five Suite")
}
