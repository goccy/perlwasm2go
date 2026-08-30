package wasm2go

import (
	base "github.com/goccy/perlwasm2go/base"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/goccy/perlwasm2go/p1"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	__memcap := reserveBytes
	if __memcap < 10485760 {
		__memcap = 10485760
	}
	m.Memory = make([]byte, 10485760, __memcap)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(10485760)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 1407)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
	InitElemSeg_1_5(m)
	m.DataEnd = 10428003
	initData_0(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, wasmify, 13107200)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env, wasmify)
}

const InitialMemoryBytes = 10485760

func NewWithMemory(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, memory []byte, memSize uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 1407)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
	InitElemSeg_1_5(m)
	m.DataEnd = 10428003
	return m
}
func NewFromSnapshot(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, memory []byte, memSize uint64, globals []uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 1407)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
	InitElemSeg_1_5(m)
	m.DataEnd = 10428003
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[8388608:], wasm2goData_data_bin[0:1902523])
	copy(m.Memory[10292368:], wasm2goData_data_bin[1902523:1905658])
	copy(m.Memory[10296716:], wasm2goData_data_bin[1905658:1921597])
	copy(m.Memory[10313788:], wasm2goData_data_bin[1921597:1925556])
	copy(m.Memory[10319024:], wasm2goData_data_bin[1925556:1943779])
	copy(m.Memory[10338320:], wasm2goData_data_bin[1943779:1953770])
	copy(m.Memory[10349480:], wasm2goData_data_bin[1953770:1955837])
	copy(m.Memory[10353280:], wasm2goData_data_bin[1955837:1961292])
	copy(m.Memory[10359868:], wasm2goData_data_bin[1961292:1965251])
	copy(m.Memory[10365104:], wasm2goData_data_bin[1965251:2010618])
	copy(m.Memory[10411652:], wasm2goData_data_bin[2010618:2012621])
	copy(m.Memory[10414680:], wasm2goData_data_bin[2012621:2014032])
	copy(m.Memory[10417196:], wasm2goData_data_bin[2014032:2018011])
	copy(m.Memory[10422428:], wasm2goData_data_bin[2018011:2023586])
}
func Initialize(m *base.Module) {
	Fn130(m)
}
func WasmAlloc(m *base.Module, l0 int32) int32 {
	return Fn131(m, l0)
}
func WasmFree(m *base.Module, l0 int32) {
	Fn132(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int32, l1 int32) int64 {
	return Fn156(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn158(m)
}
func WasmShutdown(m *base.Module) {
	Fn159(m)
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
	packed = Fn133(m, l0, l1)
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
	packed = Fn148(m, l0, l1)
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
	packed = Fn149(m, l0, l1)
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
	packed = Fn150(m, l0, l1)
	return
}
func Inv_0_4(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn151(m, l0, l1)
	return
}
func Inv_0_5(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn153(m, l0, l1)
	return
}
func Inv_0_6(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn154(m, l0, l1)
	return
}
func Inv_0_7(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn155(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
