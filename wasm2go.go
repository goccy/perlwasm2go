package wasm2go

import (
	base "github.com/goccy/perlwasm2go/base"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/goccy/perlwasm2go/p2"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	__memcap := reserveBytes
	if __memcap < 16056320 {
		__memcap = 16056320
	}
	m.Memory = make([]byte, 16056320, __memcap)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(16056320)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 0
	m.T0 = make([]any, 2655)
	m.G0 = int64(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	InitElemSeg_2_8(m)
	InitElemSeg_2_9(m)
	m.DataEnd = 15863771
	initData_0(m)
	m.ForkChild = Fn1129
	m.ForkModule = forkModule
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, wasmify, 20070400)
}

// New constructs a *Module using NewWASI(WASIConfig{}) for the
// wasi_snapshot_preview1 import, which gives the guest nothing it
// is not configured with. Use NewWithWASI to plug in another
// configuration (NewWASI(HostWASIConfig()), ...) or implementation.
func New(env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASI(base.NewWASI(base.WASIConfig{}), env, wasmify)
}

const InitialMemoryBytes = 16056320

func NewWithMemory(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, memory []byte, memSize uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 281474976710656 {
		panic("wasm2go: memory size exceeds the implementation limit (281474976710656 bytes)")
	}
	if memSize > uint64(len(memory)) {
		panic("wasm2go: memory size exceeds the memory handed in")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 2655)
	m.G0 = int64(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	InitElemSeg_2_8(m)
	InitElemSeg_2_9(m)
	m.DataEnd = 15863771
	m.ForkChild = Fn1129
	m.ForkModule = forkModule
	return m
}
func NewFromSnapshot(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, memory []byte, memSize uint64, globals []uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 281474976710656 {
		panic("wasm2go: memory size exceeds the implementation limit (281474976710656 bytes)")
	}
	if memSize > uint64(len(memory)) {
		panic("wasm2go: memory size exceeds the memory handed in")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 2655)
	m.G0 = int64(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	InitElemSeg_2_8(m)
	InitElemSeg_2_9(m)
	m.DataEnd = 15863771
	m.ForkChild = Fn1129
	m.ForkModule = forkModule
	base.RestoreGlobals(m, globals)
	return m
}

// forkModule builds an instance from memory, a copy of m's memory of
// memSize bytes, and the rest of m's state (NewFromSnapshot): its
// globals, the memory's ceiling, its dropped data segments and dataEnd.
// Its wasi import is wasi and its other imports are m's. A shared
// memory is as long as m's: the memory's ceiling.
func forkModule(m *base.Module, wasi *base.WasiStubs, memory []byte, memSize uint64) *base.Module {
	c := NewFromSnapshot(wasi, m.Env, m.Wasmify, memory, memSize, base.SaveGlobals(m))
	c.MaxMem = m.MaxMem
	c.DataEnd = m.DataEnd
	return c
}
func initData_0(m *base.Module) {
	copy(m.Memory[8388608:], wasm2goData_data_bin[0:1407195])
	copy(m.Memory[9797824:], wasm2goData_data_bin[1407195:1495318])
	copy(m.Memory[9887488:], wasm2goData_data_bin[1495318:1516257])
	copy(m.Memory[9909888:], wasm2goData_data_bin[1516257:1767212])
	copy(m.Memory[10161888:], wasm2goData_data_bin[1767212:1786607])
	copy(m.Memory[10183136:], wasm2goData_data_bin[1786607:1786802])
	copy(m.Memory[10184800:], wasm2goData_data_bin[1786802:1826349])
	copy(m.Memory[10225888:], wasm2goData_data_bin[1826349:1855032])
	copy(m.Memory[10255728:], wasm2goData_data_bin[1855032:1859619])
	copy(m.Memory[10261536:], wasm2goData_data_bin[1859619:1874318])
	copy(m.Memory[10277824:], wasm2goData_data_bin[1874318:1889873])
	copy(m.Memory[10294496:], wasm2goData_data_bin[1889873:1894604])
	copy(m.Memory[10300256:], wasm2goData_data_bin[1894604:1900215])
	copy(m.Memory[10307408:], wasm2goData_data_bin[1900215:1930162])
	copy(m.Memory[10338864:], wasm2goData_data_bin[1930162:1942405])
	copy(m.Memory[10352704:], wasm2goData_data_bin[1942405:1949384])
	copy(m.Memory[10361168:], wasm2goData_data_bin[1949384:1972611])
	copy(m.Memory[10385616:], wasm2goData_data_bin[1972611:1986590])
	copy(m.Memory[10400832:], wasm2goData_data_bin[1986590:1988865])
	copy(m.Memory[10404192:], wasm2goData_data_bin[1988865:1998676])
	copy(m.Memory[10415712:], wasm2goData_data_bin[1998676:1999919])
	copy(m.Memory[10418368:], wasm2goData_data_bin[1999919:2006298])
	copy(m.Memory[10426032:], wasm2goData_data_bin[2006298:2009989])
	copy(m.Memory[10431280:], wasm2goData_data_bin[2009989:2021776])
	copy(m.Memory[10444304:], wasm2goData_data_bin[2021776:2063707])
	copy(m.Memory[10488048:], wasm2goData_data_bin[2063707:2063718])
	copy(m.Memory[10490105:], wasm2goData_data_bin[2063718:2064728])
	copy(m.Memory[10492145:], wasm2goData_data_bin[2064728:7111490])
	copy(m.Memory[15540929:], wasm2goData_data_bin[7111490:7111508])
	copy(m.Memory[15542272:], wasm2goData_data_bin[7111508:7118727])
	copy(m.Memory[15551200:], wasm2goData_data_bin[7118727:7129898])
	copy(m.Memory[15563400:], wasm2goData_data_bin[7129898:7142085])
	copy(m.Memory[15576888:], wasm2goData_data_bin[7142085:7142088])
	copy(m.Memory[15578128:], wasm2goData_data_bin[7142088:7146667])
	copy(m.Memory[15584576:], wasm2goData_data_bin[7146667:7150798])
	copy(m.Memory[15589840:], wasm2goData_data_bin[7150798:7155737])
	copy(m.Memory[15596640:], wasm2goData_data_bin[7155737:7163916])
	copy(m.Memory[15607296:], wasm2goData_data_bin[7163916:7166959])
	copy(m.Memory[15611616:], wasm2goData_data_bin[7166959:7166970])
	copy(m.Memory[15613560:], wasm2goData_data_bin[7166970:7166973])
	copy(m.Memory[15615992:], wasm2goData_data_bin[7166973:7171072])
	copy(m.Memory[15622112:], wasm2goData_data_bin[7171072:7171083])
	copy(m.Memory[15623648:], wasm2goData_data_bin[7171083:7171190])
	copy(m.Memory[15624960:], wasm2goData_data_bin[7171190:7175145])
	copy(m.Memory[15630768:], wasm2goData_data_bin[7175145:7183668])
	copy(m.Memory[15641216:], wasm2goData_data_bin[7183668:7183743])
	copy(m.Memory[15642952:], wasm2goData_data_bin[7183743:7183962])
	copy(m.Memory[15644424:], wasm2goData_data_bin[7183962:7184421])
	copy(m.Memory[15646672:], wasm2goData_data_bin[7184421:7184560])
	copy(m.Memory[15647856:], wasm2goData_data_bin[7184560:7184571])
	copy(m.Memory[15650136:], wasm2goData_data_bin[7184571:7185574])
	copy(m.Memory[15653080:], wasm2goData_data_bin[7185574:7185601])
	copy(m.Memory[15655120:], wasm2goData_data_bin[7185601:7186612])
	copy(m.Memory[15657992:], wasm2goData_data_bin[7186612:7186671])
	copy(m.Memory[15660608:], wasm2goData_data_bin[7186671:7186690])
	copy(m.Memory[15662576:], wasm2goData_data_bin[7186690:7188861])
	copy(m.Memory[15665920:], wasm2goData_data_bin[7188861:7193608])
	copy(m.Memory[15672064:], wasm2goData_data_bin[7193608:7197939])
	copy(m.Memory[15678209:], wasm2goData_data_bin[7197939:7198325])
	copy(m.Memory[15680201:], wasm2goData_data_bin[7198325:7198343])
	copy(m.Memory[15682241:], wasm2goData_data_bin[7198343:7198361])
	copy(m.Memory[15683840:], wasm2goData_data_bin[7198361:7211572])
	copy(m.Memory[15699200:], wasm2goData_data_bin[7211572:7213415])
	copy(m.Memory[15702272:], wasm2goData_data_bin[7213415:7216730])
	copy(m.Memory[15707456:], wasm2goData_data_bin[7216730:7220861])
	copy(m.Memory[15712720:], wasm2goData_data_bin[7220861:7223520])
	copy(m.Memory[15717168:], wasm2goData_data_bin[7223520:7224011])
	copy(m.Memory[15719168:], wasm2goData_data_bin[7224011:7224022])
	copy(m.Memory[15721520:], wasm2goData_data_bin[7224022:7228153])
	copy(m.Memory[15729120:], wasm2goData_data_bin[7228153:7231116])
	copy(m.Memory[15733545:], wasm2goData_data_bin[7231116:7232750])
	copy(m.Memory[15737160:], wasm2goData_data_bin[7232750:7232969])
	copy(m.Memory[15738832:], wasm2goData_data_bin[7232969:7233108])
	copy(m.Memory[15740016:], wasm2goData_data_bin[7233108:7233119])
	copy(m.Memory[15742296:], wasm2goData_data_bin[7233119:7234122])
	copy(m.Memory[15745240:], wasm2goData_data_bin[7234122:7234149])
	copy(m.Memory[15747280:], wasm2goData_data_bin[7234149:7235160])
	copy(m.Memory[15750152:], wasm2goData_data_bin[7235160:7235219])
	copy(m.Memory[15752768:], wasm2goData_data_bin[7235219:7235238])
	copy(m.Memory[15754736:], wasm2goData_data_bin[7235238:7240865])
	copy(m.Memory[15762201:], wasm2goData_data_bin[7240865:7262987])
	copy(m.Memory[15785945:], wasm2goData_data_bin[7262987:7272389])
	copy(m.Memory[15797384:], wasm2goData_data_bin[7272389:7275144])
	copy(m.Memory[15801984:], wasm2goData_data_bin[7275144:7280075])
	copy(m.Memory[15808776:], wasm2goData_data_bin[7280075:7284478])
	copy(m.Memory[15814344:], wasm2goData_data_bin[7284478:7287137])
	copy(m.Memory[15818792:], wasm2goData_data_bin[7287137:7287628])
	copy(m.Memory[15821232:], wasm2goData_data_bin[7287628:7288839])
	copy(m.Memory[15824000:], wasm2goData_data_bin[7288839:7291186])
	copy(m.Memory[15828712:], wasm2goData_data_bin[7291186:7292053])
	copy(m.Memory[15831081:], wasm2goData_data_bin[7292053:7293687])
	copy(m.Memory[15834768:], wasm2goData_data_bin[7293687:7293754])
	copy(m.Memory[15836360:], wasm2goData_data_bin[7293754:7293853])
	copy(m.Memory[15837576:], wasm2goData_data_bin[7293853:7293864])
	copy(m.Memory[15839800:], wasm2goData_data_bin[7293864:7294851])
	copy(m.Memory[15842761:], wasm2goData_data_bin[7294851:7294853])
	copy(m.Memory[15844800:], wasm2goData_data_bin[7294853:7295856])
	copy(m.Memory[15847664:], wasm2goData_data_bin[7295856:7295947])
	copy(m.Memory[15850264:], wasm2goData_data_bin[7295947:7295998])
	copy(m.Memory[15852216:], wasm2goData_data_bin[7295998:7307553])
}

// InstallDataSegments copies the module's active data segments into
// memory, as New does into the memory it allocates. NewWithMemory takes
// its memory as it is: a fresh memory (all zeros, InitialMemoryBytes
// long or more) gets the segments from here first.
func InstallDataSegments(memory []byte) {
	m := &base.Module{}
	m.Memory = memory
	initData_0(m)
}
func Initialize(m *base.Module) {
	Fn115(m)
}
func WasmAlloc(m *base.Module, l0 int64) int64 {
	return Fn116(m, l0)
}
func WasmFree(m *base.Module, l0 int64) {
	Fn117(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int64, l1 int64) int64 {
	return Fn160(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn162(m)
}
func WasmShutdown(m *base.Module) {
	Fn163(m)
}
func Wasip1PosixExtensionForkChild(m *base.Module) int32 {
	return Fn1129(m)
}
func Inv_0_0(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn118(m, l0, l1)
	return
}
func Inv_0_1(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn130(m, l0, l1)
	return
}
func Inv_0_2(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn131(m, l0, l1)
	return
}
func Inv_0_3(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn135(m, l0, l1)
	return
}
func Inv_0_4(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn136(m, l0, l1)
	return
}
func Inv_0_5(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn137(m, l0, l1)
	return
}
func Inv_0_6(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn138(m, l0, l1)
	return
}
func Inv_0_7(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn139(m, l0, l1)
	return
}
func Inv_0_8(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn140(m, l0, l1)
	return
}
func Inv_0_9(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn141(m, l0, l1)
	return
}
func Inv_0_10(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn142(m, l0, l1)
	return
}
func Inv_0_11(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn143(m, l0, l1)
	return
}
func Inv_0_12(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn144(m, l0, l1)
	return
}
func Inv_0_13(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn145(m, l0, l1)
	return
}
func Inv_0_14(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn147(m, l0, l1)
	return
}
func Inv_0_15(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn148(m, l0, l1)
	return
}
func Inv_0_16(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn149(m, l0, l1)
	return
}
func Inv_0_17(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn150(m, l0, l1)
	return
}
func Inv_0_18(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn151(m, l0, l1)
	return
}
func Inv_0_19(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn152(m, l0, l1)
	return
}
func Inv_0_20(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn153(m, l0, l1)
	return
}
func Inv_0_21(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn154(m, l0, l1)
	return
}
func Inv_0_22(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn156(m, l0, l1)
	return
}
func Inv_0_23(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn157(m, l0, l1)
	return
}
func Inv_0_24(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn158(m, l0, l1)
	return
}
func Inv_0_25(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn159(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
