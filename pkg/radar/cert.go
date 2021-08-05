package radar

import "fmt"

type Cert struct {
	Issuer    string
	ExpiresAt string
	IsValid   bool
}

func (c *Cert) String() string {
	return fmt.Sprintf(`Issuer: %s
Expries At: %s
Valid: %t
`, c.Issuer, c.ExpiresAt, c.IsValid)
}
