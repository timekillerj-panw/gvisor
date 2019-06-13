// automatically generated by stateify.

package usermem

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *AccessType) beforeSave() {}
func (x *AccessType) save(m state.Map) {
	x.beforeSave()
	m.Save("Read", &x.Read)
	m.Save("Write", &x.Write)
	m.Save("Execute", &x.Execute)
}

func (x *AccessType) afterLoad() {}
func (x *AccessType) load(m state.Map) {
	m.Load("Read", &x.Read)
	m.Load("Write", &x.Write)
	m.Load("Execute", &x.Execute)
}

func (x *Addr) save(m state.Map) {
	m.SaveValue("", (uintptr)(*x))
}

func (x *Addr) load(m state.Map) {
	m.LoadValue("", new(uintptr), func(y interface{}) { *x = (Addr)(y.(uintptr)) })
}

func (x *AddrRange) beforeSave() {}
func (x *AddrRange) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
}

func (x *AddrRange) afterLoad() {}
func (x *AddrRange) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
}

func init() {
	state.Register("usermem.AccessType", (*AccessType)(nil), state.Fns{Save: (*AccessType).save, Load: (*AccessType).load})
	state.Register("usermem.Addr", (*Addr)(nil), state.Fns{Save: (*Addr).save, Load: (*Addr).load})
	state.Register("usermem.AddrRange", (*AddrRange)(nil), state.Fns{Save: (*AddrRange).save, Load: (*AddrRange).load})
}
