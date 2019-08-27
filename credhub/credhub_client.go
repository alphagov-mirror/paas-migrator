package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub/credentials"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/values"
)

type CredHub interface {
	SetPassword(name string, value values.Password) (credentials.Password, error)
	SetCertificate(name string, value values.Certificate) (credentials.Certificate, error)
	SetRSA(name string, value values.RSA) (credentials.RSA, error)
	SetSSH(name string, value values.SSH) (credentials.SSH, error)
}
