package ghttp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"

	"testing"
)

func TestGHTTP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GHTTP Suite")
}
