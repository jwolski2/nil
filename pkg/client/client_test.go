package client

import (
	"io/ioutil"
	"math/big"
	"path"
	"strconv"
	"testing"

	"github.com/jwolski2/nil-extended/pkg/crypto"
	"github.com/jwolski2/nil-extended/pkg/server"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func TestRegister(t *testing.T) {
	// Setup.
	params, err := crypto.Load(path.Join("../../", crypto.DefaultParamsFile))
	require.NoError(t, err, "No error after loading params file")

	client := &Client{
		AuthServer: server.NewAuthServer(params),
		Params:     params,
	}

	// Exercise test subject.
	err = client.Register("wolski", big.NewInt(12345))

	// Assert.
	assert.NoError(t, err, "No error after registering")
}

func TestRegisterTwice(t *testing.T) {
	// Setup.
	params, err := crypto.Load(path.Join("../../", crypto.DefaultParamsFile))
	require.NoError(t, err, "No error after loading params file")

	client := &Client{
		AuthServer: server.NewAuthServer(params),
		Params:     params,
	}

	// Register once.
	err = client.Register("wolski", big.NewInt(12345))
	assert.NoError(t, err, "No error after registering the first time")

	// Register twice.
	err = client.Register("wolski", big.NewInt(12345))
	assert.Error(t, err, "Error occurred after registering a second time")
}

func TestRegisterAndLogin(t *testing.T) {
	// Setup.
	params, err := crypto.Load(path.Join("../../", crypto.DefaultParamsFile))
	require.NoError(t, err, "No error after loading params file")

	client := &Client{
		AuthServer: server.NewAuthServer(params),
		Params:     params,
	}

	// Register.
	err = client.Register("wolski", big.NewInt(12345))
	assert.NoError(t, err, "No error after registering the first time")

	// Login.
	sessionID, err := client.Login("wolski", big.NewInt(12345))
	assert.NoError(t, err, "No error after logging in")
	assert.NotEmpty(t, sessionID, "Session ID is not empty")
}

func TestLoginWrongPassword(t *testing.T) {
	// Setup.
	params, err := crypto.Load(path.Join("../../", crypto.DefaultParamsFile))
	require.NoError(t, err, "No error after loading params file")

	client := &Client{
		AuthServer: server.NewAuthServer(params),
		Params:     params,
	}

	// Register.
	err = client.Register("wolski", big.NewInt(12345))
	assert.NoError(t, err, "No error after registering the first time")

	// Login with wrong password.
	sessionID, err := client.Login("wolski", big.NewInt(123456))
	assert.Error(t, err, "No error after logging in")
	assert.Empty(t, sessionID, "Session ID is empty")
}

func TestLoginWithoutRegistering(t *testing.T) {
	// Setup.
	params, err := crypto.Load(path.Join("../../", crypto.DefaultParamsFile))
	require.NoError(t, err, "No error after loading params file")

	client := &Client{
		AuthServer: server.NewAuthServer(params),
		Params:     params,
	}

	// Login without registering .
	sessionID, err := client.Login("wolski", big.NewInt(12345))
	assert.Error(t, err, "No error after logging in")
	assert.Empty(t, sessionID, "Session ID is empty")
}

func TestLoginTwice(t *testing.T) {
	// Setup.
	params, err := crypto.Load(path.Join("../../", crypto.DefaultParamsFile))
	require.NoError(t, err, "No error after loading params file")

	client := &Client{
		AuthServer: server.NewAuthServer(params),
		Params:     params,
	}

	// Register.
	err = client.Register("wolski", big.NewInt(12345))
	assert.NoError(t, err, "No error after registering the first time")

	// Login once.
	sessionID, err := client.Login("wolski", big.NewInt(12345))
	assert.NoError(t, err, "No error after logging in once")
	assert.NotEmpty(t, sessionID, "Session ID is not empty after first login")

	// Login twice.
	sessionID, err = client.Login("wolski", big.NewInt(12345))
	assert.Error(t, err, "Error after logging in a second time")
	assert.Empty(t, sessionID, "Session ID is empty after second login")
}

func TestWithAlternativeParamsFiles(t *testing.T) {
	paramsFiles, err := ioutil.ReadDir("../../data")
	require.NoError(t, err, "No error listing files in data dir")

	for i, f := range paramsFiles {
		params, err := crypto.Load(path.Join("../../data", f.Name()))
		require.NoError(t, err, "No error after loading params file")

		client := &Client{
			AuthServer: server.NewAuthServer(params),
			Params:     params,
		}

		user := "wolski" + strconv.Itoa(i)

		// Register.
		err = client.Register(user, big.NewInt(12345))
		assert.NoError(t, err, "No error after registering the first time")

		// Login once.
		sessionID, err := client.Login(user, big.NewInt(12345))
		assert.NoError(t, err, "No error after logging in once")
		assert.NotEmpty(t, sessionID, "Session ID is not empty after first login")
	}
}
