package crypto

import "math/big"

const (
	DefaultParamsFile = "data/params1.json"
)

// Params represents the public set of parameters known to the client/server in
// order to successfully perform the authentication protocol. There's a current
// limitation which uses 1 set of parameters for all registered users as opposed
// to using 1 set of parameters per user.
type Params struct {
	P64 int64 `json:"p"`
	G64 int64 `json:"g"`
	H64 int64 `json:"h"`
	Q64 int64 `json:"q"`
}

func (p *Params) P() *big.Int {
	return big.NewInt(p.P64)
}

func (p *Params) G() *big.Int {
	return big.NewInt(p.G64)
}

func (p *Params) H() *big.Int {
	return big.NewInt(p.H64)
}

func (p *Params) Q() *big.Int {
	return big.NewInt(p.Q64)
}
