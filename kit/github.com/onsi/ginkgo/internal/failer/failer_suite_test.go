package failer_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"

	"testing"
)

func TestFailer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failer Suite")
}
