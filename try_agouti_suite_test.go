package try_agouti_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
)

func TestTryAgouti(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TryAgouti Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	// Choose a WebDriver:

	// agoutiDriver = agouti.PhantomJS()
	agoutiDriver = agouti.Selenium()
	// agoutiDriver = agouti.ChromeDriver()

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
