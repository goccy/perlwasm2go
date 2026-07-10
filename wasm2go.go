package wasm2go

import (
	base "github.com/goccy/perlwasm2go/base"
	"fmt"
	"unsafe"
	_ "github.com/goccy/perlwasm2go/p1"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	__memcap := reserveBytes
	if __memcap < 10485760 {
		__memcap = 10485760
	}
	m.Memory = make([]byte, 10485760, __memcap)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 1379)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
	InitElemSeg_1_5(m)
	initData_0(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, 13107200)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env)
}
func initData_0(m *base.Module) {
	copy(m.Memory[8388608:], wasm2goData_data_bin[0:1896251])
	copy(m.Memory[10286096:], wasm2goData_data_bin[1896251:1899386])
	copy(m.Memory[10290444:], wasm2goData_data_bin[1899386:1915325])
	copy(m.Memory[10307516:], wasm2goData_data_bin[1915325:1919284])
	copy(m.Memory[10312752:], wasm2goData_data_bin[1919284:1937507])
	copy(m.Memory[10332048:], wasm2goData_data_bin[1937507:1947498])
	copy(m.Memory[10343208:], wasm2goData_data_bin[1947498:1949565])
	copy(m.Memory[10347008:], wasm2goData_data_bin[1949565:1955020])
	copy(m.Memory[10353596:], wasm2goData_data_bin[1955020:1958979])
	copy(m.Memory[10358832:], wasm2goData_data_bin[1958979:2004346])
	copy(m.Memory[10405380:], wasm2goData_data_bin[2004346:2006349])
	copy(m.Memory[10408409:], wasm2goData_data_bin[2006349:2007759])
	copy(m.Memory[10410924:], wasm2goData_data_bin[2007759:2011738])
	copy(m.Memory[10416156:], wasm2goData_data_bin[2011738:2017313])
}
func Initialize(m *base.Module) {
	Fn128(m)
}
func WasmAlloc(m *base.Module, l0 int32) int32 {
	return Fn129(m, l0)
}
func WasmFree(m *base.Module, l0 int32) {
	Fn130(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int32, l1 int32) int64 {
	return Fn147(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn149(m)
}
func WasmShutdown(m *base.Module) {
	Fn150(m)
}
func Inv_0_0(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn131(m, l0, l1)
	return
}
func Inv_0_1(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn138(m, l0, l1)
	return
}
func Inv_0_2(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn145(m, l0, l1)
	return
}
func Inv_0_3(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn146(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
