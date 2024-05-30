package utils

import (
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/stretchr/testify/assert"
)

/*
const keyTemplate = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GopenPGP

xsBNBFmEoNEBCADNwbStWJTOetG3Ts4IzB2XbHnLsN84QpPB+/1ffXgQPVffr32g
CAkUweCWikcS/vH8f/XtAPTZ7XnsEccXXqkMkzq8jdoheW9U8Lwq8BbPoae2RS4e
Y+jO1O1sDi8Ofn6m4YswxD5rxYcHbVggMNnqG5p5Xp2CW3+Ahvnf9IV9Ph6JcOe3
Gz26E01MMS1lO/6ASmnEDK4R4T7Xk8lDBimUcfTm/s+xkKOGx2A9HBU+7B44s8tL
Q0E8fy/z0wrUQu2rTzV6gMxk8/R3NU4SzJ61ChXgNkuQ+NjUzyH1KsoU3fw6ylgO
kq1ScFtXGFxMwGq6aRFXQTrGrSdpwQ9lUsZ1ABEBAAHNG2pvaG5kb2VAZXhhbXBs
ZS5jb20gPHRlc3RAZXhhbXBsZS5jb20+wsBcBBABCAApBQJZhKDRBgsJBwgDAgQV
CAoCBBYCAwECHgECF4AACgkQFrF9QaNxZNBHEQgAmk1GmMrdA5uc/BbJ6xM6DzwH
GC5TPw8DbOimUIWJSkA8tU5VZshp8pVDJ39mErz2PLd6D3vdsSK8uE1ZzNTpVGMs
J3duyrmvQ1Zw6Jf+O+KePKyM6pK8RBvwI4EDs3RuyBZVXPu5KFL0ZwvRFUJ1Ar64
clz2Z0Qk8FgDtO2mOp3Vprx91Be6Nsbvj5K5mYNm9XRRl8wivAE3mrYQ02hZDSqL
7r1UeA24ELXYPtYrJ8g1Jff7PPRMv/cLuF/mZZYZj/bZsKO3S3FZrJ0V2bqcd/nO
AIlO5h07sxC4yNglDJ40TF0/yOscozXoS5Zse4SbF9dFjcvsApjG7sxchvK0U7gI
p6S/Nw==
=H8zF
-----END PGP PUBLIC KEY BLOCK-----`
*/
func TestCheckKeyExpiration(t *testing.T) {
	// Generate a test key with no expiration (or a very long expiration date)
	rsaKey, _ := crypto.GenerateKey("", "hi@example.com", "rsa", 2048)
	publicKey, err := rsaKey.GetArmoredPublicKey()

	// Test with a non-expired key
	assert.False(t, CheckKeyExpiration(publicKey))
	assert.NoError(t, err)
	/*
		assert.False(t, CheckKeyExpiration(keyTemplate))
		assert.NoError(t, err)

		//TODO add some old key :)
	*/
}
