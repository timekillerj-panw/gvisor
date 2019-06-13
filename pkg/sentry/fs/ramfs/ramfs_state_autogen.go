// automatically generated by stateify.

package ramfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *Dir) beforeSave() {}
func (x *Dir) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Save("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Save("children", &x.children)
	m.Save("dentryMap", &x.dentryMap)
}

func (x *Dir) afterLoad() {}
func (x *Dir) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Load("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Load("children", &x.children)
	m.Load("dentryMap", &x.dentryMap)
}

func (x *dirFileOperations) beforeSave() {}
func (x *dirFileOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("dirCursor", &x.dirCursor)
	m.Save("dir", &x.dir)
}

func (x *dirFileOperations) afterLoad() {}
func (x *dirFileOperations) load(m state.Map) {
	m.Load("dirCursor", &x.dirCursor)
	m.Load("dir", &x.dir)
}

func (x *Socket) beforeSave() {}
func (x *Socket) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Save("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Save("ep", &x.ep)
}

func (x *Socket) afterLoad() {}
func (x *Socket) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Load("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Load("ep", &x.ep)
}

func (x *socketFileOperations) beforeSave() {}
func (x *socketFileOperations) save(m state.Map) {
	x.beforeSave()
}

func (x *socketFileOperations) afterLoad() {}
func (x *socketFileOperations) load(m state.Map) {
}

func (x *Symlink) beforeSave() {}
func (x *Symlink) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Save("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Save("Target", &x.Target)
}

func (x *Symlink) afterLoad() {}
func (x *Symlink) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Load("InodeSimpleExtendedAttributes", &x.InodeSimpleExtendedAttributes)
	m.Load("Target", &x.Target)
}

func (x *symlinkFileOperations) beforeSave() {}
func (x *symlinkFileOperations) save(m state.Map) {
	x.beforeSave()
}

func (x *symlinkFileOperations) afterLoad() {}
func (x *symlinkFileOperations) load(m state.Map) {
}

func init() {
	state.Register("ramfs.Dir", (*Dir)(nil), state.Fns{Save: (*Dir).save, Load: (*Dir).load})
	state.Register("ramfs.dirFileOperations", (*dirFileOperations)(nil), state.Fns{Save: (*dirFileOperations).save, Load: (*dirFileOperations).load})
	state.Register("ramfs.Socket", (*Socket)(nil), state.Fns{Save: (*Socket).save, Load: (*Socket).load})
	state.Register("ramfs.socketFileOperations", (*socketFileOperations)(nil), state.Fns{Save: (*socketFileOperations).save, Load: (*socketFileOperations).load})
	state.Register("ramfs.Symlink", (*Symlink)(nil), state.Fns{Save: (*Symlink).save, Load: (*Symlink).load})
	state.Register("ramfs.symlinkFileOperations", (*symlinkFileOperations)(nil), state.Fns{Save: (*symlinkFileOperations).save, Load: (*symlinkFileOperations).load})
}
