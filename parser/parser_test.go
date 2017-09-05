package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/ishustava/migrator/test_fixtures"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/values"
	"github.com/ishustava/migrator/credentials"
	"github.com/ishustava/migrator/parser"
)

var _ = Describe("Parser", func() {
	Describe("#FindCredentials", func() {
		Context("Passwords", func() {
			It("finds and returns password credentials", func() {
				password1 := credentials.NewPassword("path1", "password1")
				password2 := credentials.NewPassword("path2", "password2")

				varsStore := map[string]interface{}{
					"path1": "password1",
					"path2": "password2",
				}

				creds, err := parser.FindCredentials(varsStore)

				Expect(err).ToNot(HaveOccurred())
				Expect(creds.Passwords).To(ConsistOf(password1, password2))
			})
		})

		Context("Certificates", func() {
			It("finds and returns certificate credentials", func() {
				cert1 := credentials.NewCertificate("path3", values.Certificate{Ca: CA1, Certificate: CERT1, PrivateKey: PRIV1})
				cert2 := credentials.NewCertificate("path4", values.Certificate{Certificate: CERT2, PrivateKey: PRIV2})

				varsStore := map[string]interface{}{
					"path3": map[interface{}]interface{}{
						"ca": CA1,
						"certificate": CERT1,
						"private_key": PRIV1,
					},
					"path4": map[interface{}]interface{}{
						"certificate": CERT2,
						"private_key": PRIV2,
					},
				}

				creds, err := parser.FindCredentials(varsStore)

				Expect(err).ToNot(HaveOccurred())
				Expect(creds.Certificates).To(ConsistOf(cert1, cert2))
			})
		})

		Context("SSH", func() {
			It("finds and returns ssh credentials", func() {
				ssh := credentials.NewSsh("path5", values.SSH{PublicKey: SSH_PUB, PrivateKey: SSH_PRIV})

				varsStore := map[string]interface{}{
					"path5": map[interface{}]interface{}{
						"public_key": SSH_PUB,
						"private_key": SSH_PRIV,
						"public_key_fingerprint": "fingerprint",
					},
				}

				creds, err := parser.FindCredentials(varsStore)

				Expect(err).ToNot(HaveOccurred())
				Expect(creds.SshKeys).To(ConsistOf(ssh))
			})
		})

		Context("RSA", func() {
			It("finds and returns rsa credentials", func() {
				rsa := credentials.NewRsa("path6", values.RSA{PublicKey: RSA_PUB, PrivateKey: RSA_PRIV})

				varsStore := map[string]interface{}{
					"path6": map[interface{}]interface{}{
						"public_key": RSA_PUB,
						"private_key": RSA_PRIV,
					},
				}

				creds, err := parser.FindCredentials(varsStore)

				Expect(err).ToNot(HaveOccurred())
				Expect(creds.RsaKeys).To(ConsistOf(rsa))
			})
		})
	})
})
