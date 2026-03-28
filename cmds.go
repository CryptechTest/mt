//go:generate ./mkpktinfos.sh
//go:generate ./cmd.sh

package mt

import (
	"github.com/CryptechTest/mt/rudp"
)

type Cmd interface {
	DefaultPktInfo() rudp.PktInfo
	cmd()
}
