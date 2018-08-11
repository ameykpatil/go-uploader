package s3

import (
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"testing"
)

func Test(t *testing.T) {
	id := "123"

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("SignedURLForPut", func() {

		g.It("Should return putObject signed url for given id successfully", func() {
			signedURL, err := SignedURLForPut(id)
			Expect(err).To(BeNil())
			Expect(signedURL).ToNot(BeEmpty())
			Expect(signedURL).To(ContainSubstring(id))
		})
	})

	g.Describe("SignedURLForGet", func() {
		existingID := "1533912247399"

		g.It("Should return getObject signed url for existing id", func() {
			signedURL, err := SignedURLForGet(existingID, 60)
			Expect(err).To(BeNil())
			Expect(signedURL).ToNot(BeEmpty())
			Expect(signedURL).To(ContainSubstring(existingID))
		})

		g.It("Should return getObject signed url for given timeout", func() {
			signedURL, err := SignedURLForGet(existingID, 100)
			Expect(err).To(BeNil())
			Expect(signedURL).ToNot(BeEmpty())
			Expect(signedURL).To(ContainSubstring(existingID))
			Expect(signedURL).To(ContainSubstring("X-Amz-Expires=100"))
		})
	})
}
