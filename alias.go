package wasm2go

import (
	base "github.com/goccy/perlwasm2go/base"
	_ "unsafe"
)

//go:linkname Fn128 github.com/goccy/perlwasm2go/p1.Fn128
func Fn128(m *base.Module)

//go:linkname Fn129 github.com/goccy/perlwasm2go/p1.Fn129
func Fn129(m *base.Module, l0 int32) int32

//go:linkname Fn130 github.com/goccy/perlwasm2go/p1.Fn130
func Fn130(m *base.Module, l0 int32)

//go:linkname Fn131 github.com/goccy/perlwasm2go/p1.Fn131
func Fn131(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn138 github.com/goccy/perlwasm2go/p1.Fn138
func Fn138(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn145 github.com/goccy/perlwasm2go/p1.Fn145
func Fn145(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn146 github.com/goccy/perlwasm2go/p0.Fn146
func Fn146(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn147 github.com/goccy/perlwasm2go/p1.Fn147
func Fn147(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn149 github.com/goccy/perlwasm2go/p1.Fn149
func Fn149(m *base.Module) int32

//go:linkname Fn150 github.com/goccy/perlwasm2go/p1.Fn150
func Fn150(m *base.Module)

//go:linkname InitElemSeg_0_0 github.com/goccy/perlwasm2go/p0.InitElemSeg_0_0
func InitElemSeg_0_0(m *base.Module)

//go:linkname InitElemSeg_1_0 github.com/goccy/perlwasm2go/p1.InitElemSeg_1_0
func InitElemSeg_1_0(m *base.Module)

//go:linkname InitElemSeg_1_1 github.com/goccy/perlwasm2go/p1.InitElemSeg_1_1
func InitElemSeg_1_1(m *base.Module)

//go:linkname InitElemSeg_1_2 github.com/goccy/perlwasm2go/p1.InitElemSeg_1_2
func InitElemSeg_1_2(m *base.Module)

//go:linkname InitElemSeg_1_3 github.com/goccy/perlwasm2go/p1.InitElemSeg_1_3
func InitElemSeg_1_3(m *base.Module)

//go:linkname InitElemSeg_1_4 github.com/goccy/perlwasm2go/p1.InitElemSeg_1_4
func InitElemSeg_1_4(m *base.Module)

//go:linkname InitElemSeg_1_5 github.com/goccy/perlwasm2go/p1.InitElemSeg_1_5
func InitElemSeg_1_5(m *base.Module)
