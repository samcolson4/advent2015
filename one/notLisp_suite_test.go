package notLisp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrackercsv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "notLisp Suite")
}
