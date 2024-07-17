package core

import "sync/atomic"

type Pero struct {
	sID    atomic.Uint64
	iID    atomic.Uint64
	status bool
}

var pero Pero

func init() {

}
