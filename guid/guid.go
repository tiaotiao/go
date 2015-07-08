/*
	A global unique id generator

	Use the algorithm named Snowflake by Twitter
	Implements by Tiaotiao tiaotiaoyly@gmail.com
*/

package guid

import (
	"crypto/rand"
	"hash/crc32"
	"sync/atomic"
	"time"
)

const (
	toMillisecond = int64(time.Millisecond / time.Nanosecond)

	offsetTimestamp = uint(22)
	offsetMachineId = uint(14)
	offsetCounter   = uint(0)

	maskTimestamp = int64(0x7FFFFFFFFFC00000) // top 41bits
	maskMachineId = int64(0x00000000003FC000) // mid 8bits
	maskCounter   = int64(0x0000000000003FFF) // tail 14bits
)

type GUID struct {
	machine uint32
	counter uint32
}

func NewGUID() *GUID {
	var machine [32]byte
	_, err := rand.Read(machine[:32])
	if err != nil {
		panic(err.Error())
	}
	return NewGUIDMachine(string(machine[:]))
}

func NewGUIDMachine(machine string) *GUID {
	g := new(GUID)
	h := crc32.NewIEEE()
	h.Write([]byte(machine))
	g.machine = h.Sum32()
	g.counter = 0
	return g
}

func (g *GUID) Next() int64 {
	var id int64 = 0
	now := time.Now().UnixNano() / toMillisecond
	count := atomic.AddUint32(&g.counter, 1)

	id |= int64(now<<offsetTimestamp) & maskTimestamp
	id |= int64(g.machine<<offsetMachineId) & maskMachineId
	id |= int64(count<<offsetCounter) & maskCounter

	if id < 0 { // will NOT occur within 24 years
		id = -id
	}
	return id
}

var defaultGUID *GUID

func Next() int64 {
	return defaultGUID.Next()
}

func init() {
	defaultGUID = NewGUID()
}
