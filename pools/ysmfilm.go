package pools

import (
	"fmt"
	"time"

	"github.com/vertcoin-project/one-click-miner-vnext/util"
)

var _ Pool = &Ysmfilm{}

type Ysmfilm struct {
	Address           string
	LastFetchedPayout time.Time
	LastPayout        uint64
}

func NewYsmfilm(addr string) *Ysmfilm {
	return &Ysmfilm{Address: addr}
}

func (p *Ysmfilm) GetPendingPayout() uint64 {
	jsonPayload := map[string]interface{}{}
	err := util.GetJson(fmt.Sprintf("https://pool.ysmfilm.net/api/pools/vert1/miners/%s", p.Address), &jsonPayload)
	if err != nil {
		return 0
	}
	vtc, ok := jsonPayload["pendingBalance"].(float64)
	if !ok {
		return 0
	}
	return uint64(vtc)
}

func (p *Ysmfilm) GetStratumUrl() string {
	return "stratum+tcp://vtc.pool.ysmfilm.net:3057"
}

func (p *Ysmfilm) GetUsername() string {
	return p.Address
}

func (p *Ysmfilm) GetPassword() string {
	return "x"
}

func (p *Ysmfilm) GetID() int {
	return 7
}

func (p *Ysmfilm) GetName() string {
	return "pool.ysmfilm.net"
}

func (p *Ysmfilm) GetFee() float64 {
	return 1
}

func (p *Ysmfilm) OpenBrowserPayoutInfo(addr string) {
	util.OpenBrowser(fmt.Sprintf("https://pool.ysmfilm.net/?#vert1/dashboard?address=%s", addr))
}
