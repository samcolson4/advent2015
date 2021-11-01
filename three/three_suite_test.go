package three_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestThree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Three Suite")
}
