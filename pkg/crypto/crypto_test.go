package crypto

import (
	"io/ioutil"
	"math/big"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testVerificationWithParamsFile(t *testing.T, paramsFile string) {
	// Setup.
	params, err := Load(path.Join("../../data", paramsFile))
	require.NoError(t, err, "No error after loading params file")

	x := big.NewInt(123456789)

	y1, y2, err := ComputeY1AndY2(params, x)
	require.NoError(t, err, "No error after computing y1 and y2")

	c, err := GenerateC()
	require.NoError(t, err, "No error after generating c")

	r1, r2, k, err := ComputeR1AndR2(params)
	require.NoError(t, err, "No error after computing r1 and r2")

	s, err := ComputeS(params, x, k, c)
	require.NoError(t, err, "No error after computing s")

	// Assert.
	verified := VerifyR1AndR2(params, r1, r2, s, c, y1, y2)
	assert.True(t, verified, "r1 and r2 are verified")
}

func TestVerificationAgainstAllFixtures(t *testing.T) {
	paramsFiles, err := ioutil.ReadDir("../../data")
	require.NoError(t, err, "No error listing files in data dir")

	for _, f := range paramsFiles {
		testVerificationWithParamsFile(t, f.Name())
	}
}

func TestWrongS(t *testing.T) {
	// Setup.
	params, err := Load(path.Join("../../", DefaultParamsFile))
	require.NoError(t, err, "No error after loading params file")

	x := big.NewInt(123456789)

	y1, y2, err := ComputeY1AndY2(params, x)
	require.NoError(t, err, "No error after computing y1 and y2")

	c, err := GenerateC()
	require.NoError(t, err, "No error after generating c")

	r1, r2, k, err := ComputeR1AndR2(params)
	require.NoError(t, err, "No error after computing r1 and r2")

	wrongX := big.NewInt(987654321)
	s, err := ComputeS(params, wrongX, k, c)
	require.NoError(t, err, "No error after computing s")

	// Assert.
	verified := VerifyR1AndR2(params, r1, r2, s, c, y1, y2)
	assert.False(t, verified, "r1 and r2 are not verified")
}
