package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/values"
	"github.com/alphagov/migrator/credentials"
)

func BulkSet(credentials *credentials.Credentials, credHubClient CredHub, observer BulkSetObserver) error {
	observer.BeginBulkSet(
		len(credentials.Passwords),
		len(credentials.Certificates),
		len(credentials.RsaKeys),
		len(credentials.SshKeys),
	)
	for _, pass := range credentials.Passwords {
		if _, err := credHubClient.SetPassword(pass.Name, pass.Value); err != nil {
			observer.FailPasswordSet(pass.Name, err)
		}
	}
	observer.EndPasswordsSet()

	for _, cert := range credentials.Certificates {
		if _, err := credHubClient.SetCertificate(cert.Name, cert.Value); err != nil {
			observer.FailCertificateSet(cert.Name, err)
		}
	}
	observer.EndCertificatesSet()

	for _, rsa := range credentials.RsaKeys {
		if _, err := credHubClient.SetRSA(rsa.Name, rsa.Value); err != nil {
			observer.FailRsaKeySet(rsa.Name, err)
		}
	}
	observer.EndRsaKeysSet()

	for _, ssh := range credentials.SshKeys {
		_, err := credHubClient.SetSSH(
			ssh.Name,
			values.SSH{
				PublicKey:  ssh.Value.PublicKey,
				PrivateKey: ssh.Value.PrivateKey,
			})
		if err != nil {
			observer.FailSshKeySet(ssh.Name, err)
		}
	}
	observer.EndSshKeysSet()

	return observer.EndBulkSet()
}
