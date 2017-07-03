package grants

import (
	"strings"
	"testing"

	"code.monax.io/platform/hoard/core/reference"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/openpgp"
)

func TestOpenPGPGrant(t *testing.T) {
	testRef := testReference()
	// Pull in recipient and sender test keypairs
	to, err := openpgp.ReadArmoredKeyRing(strings.NewReader(keyPairA))
	assert.NoError(t, err)
	from, err := openpgp.ReadArmoredKeyRing(strings.NewReader(keyPairB))
	assert.NoError(t, err)

	// Create grant
	grant, err := OpenPGPGrant(testRef, to, from[0])
	assert.NoError(t, err)

	// Try to read reference from grant
	ref, err := OpenPGPGrantReference(grant, to)
	assert.NoError(t, err)
	assert.EqualValues(t, testRef, ref)
}

const (
	keyPairA = `-----BEGIN PGP PRIVATE KEY BLOCK-----

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
=+yuc
-----END PGP PRIVATE KEY BLOCK-----
`
	keyPairB = `-----BEGIN PGP PRIVATE KEY BLOCK-----

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
=Kdad
-----END PGP PRIVATE KEY BLOCK-----
`
)

func testReference() *reference.Ref {
	address := []byte{
		1, 2, 3, 4, 5, 6, 7, 1,
		1, 2, 3, 4, 5, 6, 7, 1,
		1, 2, 3, 4, 5, 6, 7, 1,
		1, 2, 3, 4, 5, 6, 7, 1,
	}
	secretKey := []byte{
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	return reference.New(address, secretKey, nil)
}