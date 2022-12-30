package aes_test

import (
	"github.com/WebXense/sql/aes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test 'aes'", func() {

	var encryptKey = ")H@McQfTjWmZq4t7w!z%C*F-JaNdRgUk" // 256-bit key

	It("should be able to encrypt and decrypt", func() {
		plainText := "Hello World!"

		encrypted := aes.Encrypt(plainText, encryptKey)
		Expect(encrypted).NotTo(BeNil())
		Expect(encrypted).NotTo(BeEmpty())
		Expect(encrypted).NotTo(Equal(plainText))

		decrypted := aes.Decrypt(encrypted, encryptKey)
		Expect(decrypted).NotTo(BeNil())
		Expect(decrypted).NotTo(BeEmpty())
		Expect(decrypted).To(Equal(plainText))
	})

	It("should not be able to encrypt with empty key", func() {
		plainText := "Hello World!"
		Expect(func() { aes.Encrypt(plainText, "") }).To(Panic())
	})
})
