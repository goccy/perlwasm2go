//go:build !arm64 && (!amd64 || !amd64.v2)

package p1

import (
	base "github.com/goccy/perlwasm2go/base"
	_ "unsafe"
)

//go:linkname Fn120 github.com/goccy/perlwasm2go/p2.Fn120
func Fn120(m *base.Module, l0 int64) int64

//go:linkname Fn122 github.com/goccy/perlwasm2go/p2.Fn122
func Fn122(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn126 github.com/goccy/perlwasm2go/p2.Fn126
func Fn126(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn132 github.com/goccy/perlwasm2go/p2.Fn132
func Fn132(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn134 github.com/goccy/perlwasm2go/p2.Fn134
func Fn134(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn146 github.com/goccy/perlwasm2go/p2.Fn146
func Fn146(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn155 github.com/goccy/perlwasm2go/p2.Fn155
func Fn155(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn168 github.com/goccy/perlwasm2go/p0.Fn168
func Fn168(m *base.Module, l0 int64) int64

//go:linkname Fn169 github.com/goccy/perlwasm2go/p2.Fn169
func Fn169(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn171 github.com/goccy/perlwasm2go/p0.Fn171
func Fn171(m *base.Module, l0 int64)

//go:linkname Fn184 github.com/goccy/perlwasm2go/p2.Fn184
func Fn184(m *base.Module, l0 int64)

//go:linkname Fn191 github.com/goccy/perlwasm2go/p0.Fn191
func Fn191(m *base.Module, l0 int64) int64

//go:linkname Fn194 github.com/goccy/perlwasm2go/p2.Fn194
func Fn194(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn195 github.com/goccy/perlwasm2go/p2.Fn195
func Fn195(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn197 github.com/goccy/perlwasm2go/p2.Fn197
func Fn197(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn198 github.com/goccy/perlwasm2go/p0.Fn198
func Fn198(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn210 github.com/goccy/perlwasm2go/p0.Fn210
func Fn210(m *base.Module, l0 int64) int64

//go:linkname Fn214 github.com/goccy/perlwasm2go/p0.Fn214
func Fn214(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn222 github.com/goccy/perlwasm2go/p2.Fn222
func Fn222(m *base.Module, l0 int32) int64

//go:linkname Fn224 github.com/goccy/perlwasm2go/p0.Fn224
func Fn224(m *base.Module, l0 int32) int64

//go:linkname Fn226 github.com/goccy/perlwasm2go/p0.Fn226
func Fn226(m *base.Module, l0 int64)

//go:linkname Fn227 github.com/goccy/perlwasm2go/p0.Fn227
func Fn227(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn230 github.com/goccy/perlwasm2go/p0.Fn230
func Fn230(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn231 github.com/goccy/perlwasm2go/p2.Fn231
func Fn231(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn232 github.com/goccy/perlwasm2go/p0.Fn232
func Fn232(m *base.Module, l0 int64) int64

//go:linkname Fn233 github.com/goccy/perlwasm2go/p2.Fn233
func Fn233(m *base.Module)

//go:linkname Fn234 github.com/goccy/perlwasm2go/p0.Fn234
func Fn234(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64) int64

//go:linkname Fn235 github.com/goccy/perlwasm2go/p2.Fn235
func Fn235(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn236 github.com/goccy/perlwasm2go/p2.Fn236
func Fn236(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn240 github.com/goccy/perlwasm2go/p2.Fn240
func Fn240(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn241 github.com/goccy/perlwasm2go/p2.Fn241
func Fn241(m *base.Module, l0 int64) int32

//go:linkname Fn242 github.com/goccy/perlwasm2go/p2.Fn242
func Fn242(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn244 github.com/goccy/perlwasm2go/p0.Fn244
func Fn244(m *base.Module, l0 int64) int32

//go:linkname Fn245 github.com/goccy/perlwasm2go/p2.Fn245
func Fn245(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn247 github.com/goccy/perlwasm2go/p2.Fn247
func Fn247(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn249 github.com/goccy/perlwasm2go/p0.Fn249
func Fn249(m *base.Module, l0 int64) int64

//go:linkname Fn250 github.com/goccy/perlwasm2go/p0.Fn250
func Fn250(m *base.Module, l0 int64) float64

//go:linkname Fn251 github.com/goccy/perlwasm2go/p2.Fn251
func Fn251(m *base.Module) int32

//go:linkname Fn254 github.com/goccy/perlwasm2go/p2.Fn254
func Fn254(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn255 github.com/goccy/perlwasm2go/p2.Fn255
func Fn255(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn256 github.com/goccy/perlwasm2go/p2.Fn256
func Fn256(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn262 github.com/goccy/perlwasm2go/p2.Fn262
func Fn262(m *base.Module, l0 int64) int64

//go:linkname Fn279 github.com/goccy/perlwasm2go/p0.Fn279
func Fn279(m *base.Module, l0 int64)

//go:linkname Fn283 github.com/goccy/perlwasm2go/p2.Fn283
func Fn283(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn285 github.com/goccy/perlwasm2go/p2.Fn285
func Fn285(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn291 github.com/goccy/perlwasm2go/p2.Fn291
func Fn291(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn294 github.com/goccy/perlwasm2go/p2.Fn294
func Fn294(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int64

//go:linkname Fn300 github.com/goccy/perlwasm2go/p2.Fn300
func Fn300(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn301 github.com/goccy/perlwasm2go/p0.Fn301
func Fn301(m *base.Module, l0 int64)

//go:linkname Fn324 github.com/goccy/perlwasm2go/p0.Fn324
func Fn324(m *base.Module)

//go:linkname Fn325 github.com/goccy/perlwasm2go/p2.Fn325
func Fn325(m *base.Module, l0 int64) float64

//go:linkname Fn345 github.com/goccy/perlwasm2go/p0.Fn345
func Fn345(m *base.Module)

//go:linkname Fn353 github.com/goccy/perlwasm2go/p0.Fn353
func Fn353(m *base.Module)

//go:linkname Fn355 github.com/goccy/perlwasm2go/p2.Fn355
func Fn355(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn357 github.com/goccy/perlwasm2go/p2.Fn357
func Fn357(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn358 github.com/goccy/perlwasm2go/p0.Fn358
func Fn358(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn359 github.com/goccy/perlwasm2go/p0.Fn359
func Fn359(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn360 github.com/goccy/perlwasm2go/p2.Fn360
func Fn360(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn361 github.com/goccy/perlwasm2go/p2.Fn361
func Fn361(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn364 github.com/goccy/perlwasm2go/p2.Fn364
func Fn364(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn365 github.com/goccy/perlwasm2go/p2.Fn365
func Fn365(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn366 github.com/goccy/perlwasm2go/p2.Fn366
func Fn366(m *base.Module, l0 int64)

//go:linkname Fn395 github.com/goccy/perlwasm2go/p2.Fn395
func Fn395(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int64) int32

//go:linkname Fn397 github.com/goccy/perlwasm2go/p2.Fn397
func Fn397(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn401 github.com/goccy/perlwasm2go/p2.Fn401
func Fn401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn427 github.com/goccy/perlwasm2go/p2.Fn427
func Fn427(m *base.Module, l0 int64) int64

//go:linkname Fn428 github.com/goccy/perlwasm2go/p2.Fn428
func Fn428(m *base.Module, l0 int64) int64

//go:linkname Fn991 github.com/goccy/perlwasm2go/p2.Fn991
func Fn991(m *base.Module, l0 int32) int32

//go:linkname Fn1001 github.com/goccy/perlwasm2go/p2.Fn1001
func Fn1001(m *base.Module)

//go:linkname Fn1002 github.com/goccy/perlwasm2go/p0.Fn1002
func Fn1002(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1003 github.com/goccy/perlwasm2go/p2.Fn1003
func Fn1003(m *base.Module, l0 int32) int64

//go:linkname Fn1006 github.com/goccy/perlwasm2go/p2.Fn1006
func Fn1006(m *base.Module, l0 int64)

//go:linkname Fn1008 github.com/goccy/perlwasm2go/p2.Fn1008
func Fn1008(m *base.Module)

//go:linkname Fn1012 github.com/goccy/perlwasm2go/p0.Fn1012
func Fn1012(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn1020 github.com/goccy/perlwasm2go/p0.Fn1020
func Fn1020(m *base.Module, l0 int32) int32

//go:linkname Fn1025 github.com/goccy/perlwasm2go/p2.Fn1025
func Fn1025(m *base.Module, l0 int64) int32

//go:linkname Fn1031 github.com/goccy/perlwasm2go/p0.Fn1031
func Fn1031(m *base.Module, l0 int32) int64

//go:linkname Fn1033 github.com/goccy/perlwasm2go/p0.Fn1033
func Fn1033(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn1054 github.com/goccy/perlwasm2go/p0.Fn1054
func Fn1054(m *base.Module, l0 int64) int64

//go:linkname Fn1088 github.com/goccy/perlwasm2go/p2.Fn1088
func Fn1088(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64) int32

//go:linkname Fn1093 github.com/goccy/perlwasm2go/p2.Fn1093
func Fn1093(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int32

//go:linkname Fn1094 github.com/goccy/perlwasm2go/p2.Fn1094
func Fn1094(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1101 github.com/goccy/perlwasm2go/p2.Fn1101
func Fn1101(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32) int64

//go:linkname Fn1110 github.com/goccy/perlwasm2go/p2.Fn1110
func Fn1110(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn1111 github.com/goccy/perlwasm2go/p2.Fn1111
func Fn1111(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1113 github.com/goccy/perlwasm2go/p2.Fn1113
func Fn1113(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1115 github.com/goccy/perlwasm2go/p2.Fn1115
func Fn1115(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1116 github.com/goccy/perlwasm2go/p2.Fn1116
func Fn1116(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1117 github.com/goccy/perlwasm2go/p2.Fn1117
func Fn1117(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1119 github.com/goccy/perlwasm2go/p2.Fn1119
func Fn1119(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn1120 github.com/goccy/perlwasm2go/p2.Fn1120
func Fn1120(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1122 github.com/goccy/perlwasm2go/p2.Fn1122
func Fn1122(m *base.Module, l0 int32) int32

//go:linkname Fn1125 github.com/goccy/perlwasm2go/p2.Fn1125
func Fn1125(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1126 github.com/goccy/perlwasm2go/p2.Fn1126
func Fn1126(m *base.Module, l0 int64) int32

//go:linkname Fn1130 github.com/goccy/perlwasm2go/p2.Fn1130
func Fn1130(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1133 github.com/goccy/perlwasm2go/p2.Fn1133
func Fn1133(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn1134 github.com/goccy/perlwasm2go/p2.Fn1134
func Fn1134(m *base.Module, l0 int32) int64

//go:linkname Fn1137 github.com/goccy/perlwasm2go/p2.Fn1137
func Fn1137(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1138 github.com/goccy/perlwasm2go/p2.Fn1138
func Fn1138(m *base.Module, l0 int64) int64

//go:linkname Fn1139 github.com/goccy/perlwasm2go/p2.Fn1139
func Fn1139(m *base.Module, l0 int64) int64

//go:linkname Fn1145 github.com/goccy/perlwasm2go/p2.Fn1145
func Fn1145(m *base.Module, l0 int64) int32

//go:linkname Fn1149 github.com/goccy/perlwasm2go/p2.Fn1149
func Fn1149(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn1150 github.com/goccy/perlwasm2go/p2.Fn1150
func Fn1150(m *base.Module, l0 int32)

//go:linkname Fn1151 github.com/goccy/perlwasm2go/p2.Fn1151
func Fn1151(m *base.Module)

//go:linkname Fn1152 github.com/goccy/perlwasm2go/p2.Fn1152
func Fn1152(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn1153 github.com/goccy/perlwasm2go/p2.Fn1153
func Fn1153(m *base.Module) int64

//go:linkname Fn1154 github.com/goccy/perlwasm2go/p2.Fn1154
func Fn1154(m *base.Module)

//go:linkname Fn1156 github.com/goccy/perlwasm2go/p2.Fn1156
func Fn1156(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn1159 github.com/goccy/perlwasm2go/p2.Fn1159
func Fn1159(m *base.Module)

//go:linkname Fn1162 github.com/goccy/perlwasm2go/p2.Fn1162
func Fn1162(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1163 github.com/goccy/perlwasm2go/p2.Fn1163
func Fn1163(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn1165 github.com/goccy/perlwasm2go/p2.Fn1165
func Fn1165(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1173 github.com/goccy/perlwasm2go/p2.Fn1173
func Fn1173(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn1176 github.com/goccy/perlwasm2go/p2.Fn1176
func Fn1176(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn1177 github.com/goccy/perlwasm2go/p2.Fn1177
func Fn1177(m *base.Module, l0 int64) int32

//go:linkname Fn1179 github.com/goccy/perlwasm2go/p2.Fn1179
func Fn1179(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1180 github.com/goccy/perlwasm2go/p0.Fn1180
func Fn1180(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1181 github.com/goccy/perlwasm2go/p2.Fn1181
func Fn1181(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1183 github.com/goccy/perlwasm2go/p2.Fn1183
func Fn1183(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1184 github.com/goccy/perlwasm2go/p2.Fn1184
func Fn1184(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn1188 github.com/goccy/perlwasm2go/p2.Fn1188
func Fn1188(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1190 github.com/goccy/perlwasm2go/p2.Fn1190
func Fn1190(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1192 github.com/goccy/perlwasm2go/p2.Fn1192
func Fn1192(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1193 github.com/goccy/perlwasm2go/p2.Fn1193
func Fn1193(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1194 github.com/goccy/perlwasm2go/p2.Fn1194
func Fn1194(m *base.Module, l0 int64) int32

//go:linkname Fn1274 github.com/goccy/perlwasm2go/p2.Fn1274
func Fn1274(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int64

//go:linkname Fn1276 github.com/goccy/perlwasm2go/p2.Fn1276
func Fn1276(m *base.Module, l0 int64) int64

//go:linkname Fn1277 github.com/goccy/perlwasm2go/p2.Fn1277
func Fn1277(m *base.Module, l0 int64) int64

//go:linkname Fn1278 github.com/goccy/perlwasm2go/p2.Fn1278
func Fn1278(m *base.Module, l0 int64) int64

//go:linkname Fn1280 github.com/goccy/perlwasm2go/p2.Fn1280
func Fn1280(m *base.Module, l0 int64) int32

//go:linkname Fn1281 github.com/goccy/perlwasm2go/p2.Fn1281
func Fn1281(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn1284 github.com/goccy/perlwasm2go/p0.Fn1284
func Fn1284(m *base.Module, l0 int64) int64

//go:linkname Fn1295 github.com/goccy/perlwasm2go/p2.Fn1295
func Fn1295(m *base.Module, l0 int32)

//go:linkname Fn1320 github.com/goccy/perlwasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1321 github.com/goccy/perlwasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1322 github.com/goccy/perlwasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64)

//go:linkname Fn1333 github.com/goccy/perlwasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1356 github.com/goccy/perlwasm2go/p2.Fn1356
func Fn1356(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1357 github.com/goccy/perlwasm2go/p2.Fn1357
func Fn1357(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1358 github.com/goccy/perlwasm2go/p0.Fn1358
func Fn1358(m *base.Module, l0 int64) int32

//go:linkname Fn1369 github.com/goccy/perlwasm2go/p2.Fn1369
func Fn1369(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1378 github.com/goccy/perlwasm2go/p0.Fn1378
func Fn1378(m *base.Module, l0 int64) int64

//go:linkname Fn1379 github.com/goccy/perlwasm2go/p2.Fn1379
func Fn1379(m *base.Module) int64

//go:linkname Fn1380 github.com/goccy/perlwasm2go/p2.Fn1380
func Fn1380(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1382 github.com/goccy/perlwasm2go/p2.Fn1382
func Fn1382(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1418 github.com/goccy/perlwasm2go/p2.Fn1418
func Fn1418(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1419 github.com/goccy/perlwasm2go/p2.Fn1419
func Fn1419(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1420 github.com/goccy/perlwasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1461 github.com/goccy/perlwasm2go/p0.Fn1461
func Fn1461(m *base.Module, l0 int32) int64

//go:linkname Fn1468 github.com/goccy/perlwasm2go/p2.Fn1468
func Fn1468(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1482 github.com/goccy/perlwasm2go/p2.Fn1482
func Fn1482(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1580 github.com/goccy/perlwasm2go/p0.Fn1580
func Fn1580(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1599 github.com/goccy/perlwasm2go/p0.Fn1599
func Fn1599(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1605 github.com/goccy/perlwasm2go/p2.Fn1605
func Fn1605(m *base.Module, l0 int64) int64

//go:linkname Fn1606 github.com/goccy/perlwasm2go/p2.Fn1606
func Fn1606(m *base.Module, l0 int64, l1 int64) float64

//go:linkname Fn1691 github.com/goccy/perlwasm2go/p2.Fn1691
func Fn1691(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn1692 github.com/goccy/perlwasm2go/p2.Fn1692
func Fn1692(m *base.Module) int64

//go:linkname Fn1693 github.com/goccy/perlwasm2go/p2.Fn1693
func Fn1693(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1695 github.com/goccy/perlwasm2go/p2.Fn1695
func Fn1695(m *base.Module, l0 int64) int32

//go:linkname Fn1697 github.com/goccy/perlwasm2go/p2.Fn1697
func Fn1697(m *base.Module, l0 int32) int64

//go:linkname Fn1713 github.com/goccy/perlwasm2go/p2.Fn1713
func Fn1713(m *base.Module, l0 int32) int64

//go:linkname Fn1721 github.com/goccy/perlwasm2go/p2.Fn1721
func Fn1721(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1723 github.com/goccy/perlwasm2go/p2.Fn1723
func Fn1723(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1729 github.com/goccy/perlwasm2go/p0.Fn1729
func Fn1729(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64) int64

//go:linkname Fn1745 github.com/goccy/perlwasm2go/p2.Fn1745
func Fn1745(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1746 github.com/goccy/perlwasm2go/p2.Fn1746
func Fn1746(m *base.Module, l0 int64) int32

//go:linkname Fn1800 github.com/goccy/perlwasm2go/p2.Fn1800
func Fn1800(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn1801 github.com/goccy/perlwasm2go/p2.Fn1801
func Fn1801(m *base.Module) int64

//go:linkname Fn1804 github.com/goccy/perlwasm2go/p0.Fn1804
func Fn1804(m *base.Module, l0 int64) int64

//go:linkname Fn1805 github.com/goccy/perlwasm2go/p2.Fn1805
func Fn1805(m *base.Module, l0 int32) int64

//go:linkname Fn1806 github.com/goccy/perlwasm2go/p2.Fn1806
func Fn1806(m *base.Module, l0 int32) int64

//go:linkname Fn1820 github.com/goccy/perlwasm2go/p2.Fn1820
func Fn1820(m *base.Module, l0 int64)

//go:linkname Fn1821 github.com/goccy/perlwasm2go/p2.Fn1821
func Fn1821(m *base.Module, l0 int64) int64

//go:linkname Fn1822 github.com/goccy/perlwasm2go/p2.Fn1822
func Fn1822(m *base.Module, l0 int64) int64

//go:linkname Fn1823 github.com/goccy/perlwasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int32) int64

//go:linkname Fn1826 github.com/goccy/perlwasm2go/p2.Fn1826
func Fn1826(m *base.Module, l0 int64)

//go:linkname Fn1827 github.com/goccy/perlwasm2go/p2.Fn1827
func Fn1827(m *base.Module, l0 int64)

//go:linkname Fn1831 github.com/goccy/perlwasm2go/p2.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1832 github.com/goccy/perlwasm2go/p2.Fn1832
func Fn1832(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int32

//go:linkname Fn1837 github.com/goccy/perlwasm2go/p2.Fn1837
func Fn1837(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1841 github.com/goccy/perlwasm2go/p2.Fn1841
func Fn1841(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1845 github.com/goccy/perlwasm2go/p2.Fn1845
func Fn1845(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1846 github.com/goccy/perlwasm2go/p2.Fn1846
func Fn1846(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1856 github.com/goccy/perlwasm2go/p2.Fn1856
func Fn1856(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1870 github.com/goccy/perlwasm2go/p2.Fn1870
func Fn1870(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1871 github.com/goccy/perlwasm2go/p2.Fn1871
func Fn1871(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1884 github.com/goccy/perlwasm2go/p2.Fn1884
func Fn1884(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32) int64

//go:linkname Fn1920 github.com/goccy/perlwasm2go/p2.Fn1920
func Fn1920(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1921 github.com/goccy/perlwasm2go/p2.Fn1921
func Fn1921(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn1953 github.com/goccy/perlwasm2go/p2.Fn1953
func Fn1953(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1954 github.com/goccy/perlwasm2go/p2.Fn1954
func Fn1954(m *base.Module, l0 int32) int64

//go:linkname Fn1955 github.com/goccy/perlwasm2go/p2.Fn1955
func Fn1955(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn1957 github.com/goccy/perlwasm2go/p2.Fn1957
func Fn1957(m *base.Module, l0 int64)

//go:linkname Fn1959 github.com/goccy/perlwasm2go/p0.Fn1959
func Fn1959(m *base.Module)

//go:linkname Fn1980 github.com/goccy/perlwasm2go/p0.Fn1980
func Fn1980(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2050 github.com/goccy/perlwasm2go/p2.Fn2050
func Fn2050(m *base.Module, l0 int32) int64

//go:linkname Fn2051 github.com/goccy/perlwasm2go/p2.Fn2051
func Fn2051(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2052 github.com/goccy/perlwasm2go/p2.Fn2052
func Fn2052(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2053 github.com/goccy/perlwasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2060 github.com/goccy/perlwasm2go/p2.Fn2060
func Fn2060(m *base.Module, l0 int64)

//go:linkname Fn2070 github.com/goccy/perlwasm2go/p2.Fn2070
func Fn2070(m *base.Module, l0 int64) int32

//go:linkname Fn2073 github.com/goccy/perlwasm2go/p2.Fn2073
func Fn2073(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn2074 github.com/goccy/perlwasm2go/p2.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn2079 github.com/goccy/perlwasm2go/p2.Fn2079
func Fn2079(m *base.Module, l0 int64)

//go:linkname Fn2080 github.com/goccy/perlwasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int64)

//go:linkname Fn2087 github.com/goccy/perlwasm2go/p2.Fn2087
func Fn2087(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2102 github.com/goccy/perlwasm2go/p2.Fn2102
func Fn2102(m *base.Module, l0 int64) int64

//go:linkname Fn2103 github.com/goccy/perlwasm2go/p2.Fn2103
func Fn2103(m *base.Module) int64

//go:linkname Fn2106 github.com/goccy/perlwasm2go/p0.Fn2106
func Fn2106(m *base.Module, l0 int64) int64

//go:linkname Fn2109 github.com/goccy/perlwasm2go/p0.Fn2109
func Fn2109(m *base.Module, l0 int64)

//go:linkname Fn2111 github.com/goccy/perlwasm2go/p0.Fn2111
func Fn2111(m *base.Module, l0 int64)

//go:linkname Fn2115 github.com/goccy/perlwasm2go/p0.Fn2115
func Fn2115(m *base.Module, l0 int64)

//go:linkname Fn2116 github.com/goccy/perlwasm2go/p2.Fn2116
func Fn2116(m *base.Module, l0 int64)

//go:linkname Fn2118 github.com/goccy/perlwasm2go/p2.Fn2118
func Fn2118(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2120 github.com/goccy/perlwasm2go/p0.Fn2120
func Fn2120(m *base.Module, l0 int64)

//go:linkname Fn2121 github.com/goccy/perlwasm2go/p0.Fn2121
func Fn2121(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int64

//go:linkname Fn2122 github.com/goccy/perlwasm2go/p0.Fn2122
func Fn2122(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2125 github.com/goccy/perlwasm2go/p0.Fn2125
func Fn2125(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2126 github.com/goccy/perlwasm2go/p2.Fn2126
func Fn2126(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2127 github.com/goccy/perlwasm2go/p0.Fn2127
func Fn2127(m *base.Module, l0 int64) int64

//go:linkname Fn2128 github.com/goccy/perlwasm2go/p0.Fn2128
func Fn2128(m *base.Module, l0 int64) int64

//go:linkname Fn2129 github.com/goccy/perlwasm2go/p0.Fn2129
func Fn2129(m *base.Module, l0 int64) int64

//go:linkname Fn2130 github.com/goccy/perlwasm2go/p0.Fn2130
func Fn2130(m *base.Module, l0 int64) int64

//go:linkname Fn2136 github.com/goccy/perlwasm2go/p0.Fn2136
func Fn2136(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2137 github.com/goccy/perlwasm2go/p0.Fn2137
func Fn2137(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2138 github.com/goccy/perlwasm2go/p0.Fn2138
func Fn2138(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn2142 github.com/goccy/perlwasm2go/p0.Fn2142
func Fn2142(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn2143 github.com/goccy/perlwasm2go/p2.Fn2143
func Fn2143(m *base.Module, l0 int64)

//go:linkname Fn2145 github.com/goccy/perlwasm2go/p0.Fn2145
func Fn2145(m *base.Module)

//go:linkname Fn2146 github.com/goccy/perlwasm2go/p0.Fn2146
func Fn2146(m *base.Module)

//go:linkname Fn2147 github.com/goccy/perlwasm2go/p0.Fn2147
func Fn2147(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2148 github.com/goccy/perlwasm2go/p0.Fn2148
func Fn2148(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2149 github.com/goccy/perlwasm2go/p0.Fn2149
func Fn2149(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2150 github.com/goccy/perlwasm2go/p0.Fn2150
func Fn2150(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn2151 github.com/goccy/perlwasm2go/p2.Fn2151
func Fn2151(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2153 github.com/goccy/perlwasm2go/p2.Fn2153
func Fn2153(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2155 github.com/goccy/perlwasm2go/p0.Fn2155
func Fn2155(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2156 github.com/goccy/perlwasm2go/p0.Fn2156
func Fn2156(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn2157 github.com/goccy/perlwasm2go/p0.Fn2157
func Fn2157(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2160 github.com/goccy/perlwasm2go/p0.Fn2160
func Fn2160(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2162 github.com/goccy/perlwasm2go/p0.Fn2162
func Fn2162(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn2163 github.com/goccy/perlwasm2go/p2.Fn2163
func Fn2163(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn2169 github.com/goccy/perlwasm2go/p0.Fn2169
func Fn2169(m *base.Module, l0 int64) int64

//go:linkname Fn2170 github.com/goccy/perlwasm2go/p2.Fn2170
func Fn2170(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2171 github.com/goccy/perlwasm2go/p2.Fn2171
func Fn2171(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2172 github.com/goccy/perlwasm2go/p2.Fn2172
func Fn2172(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2173 github.com/goccy/perlwasm2go/p0.Fn2173
func Fn2173(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2175 github.com/goccy/perlwasm2go/p0.Fn2175
func Fn2175(m *base.Module, l0 int32) int64

//go:linkname Fn2176 github.com/goccy/perlwasm2go/p2.Fn2176
func Fn2176(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2177 github.com/goccy/perlwasm2go/p2.Fn2177
func Fn2177(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2178 github.com/goccy/perlwasm2go/p2.Fn2178
func Fn2178(m *base.Module, l0 int64) int64

//go:linkname Fn2179 github.com/goccy/perlwasm2go/p0.Fn2179
func Fn2179(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn2181 github.com/goccy/perlwasm2go/p2.Fn2181
func Fn2181(m *base.Module, l0 int64) int64

//go:linkname Fn2182 github.com/goccy/perlwasm2go/p2.Fn2182
func Fn2182(m *base.Module, l0 int32) int32

//go:linkname Fn2184 github.com/goccy/perlwasm2go/p2.Fn2184
func Fn2184(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2185 github.com/goccy/perlwasm2go/p0.Fn2185
func Fn2185(m *base.Module, l0 int64) int64

//go:linkname Fn2186 github.com/goccy/perlwasm2go/p0.Fn2186
func Fn2186(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2187 github.com/goccy/perlwasm2go/p2.Fn2187
func Fn2187(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2188 github.com/goccy/perlwasm2go/p2.Fn2188
func Fn2188(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2189 github.com/goccy/perlwasm2go/p2.Fn2189
func Fn2189(m *base.Module) int64

//go:linkname Fn2191 github.com/goccy/perlwasm2go/p0.Fn2191
func Fn2191(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2192 github.com/goccy/perlwasm2go/p0.Fn2192
func Fn2192(m *base.Module, l0 int64)

//go:linkname Fn2193 github.com/goccy/perlwasm2go/p2.Fn2193
func Fn2193(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2194 github.com/goccy/perlwasm2go/p2.Fn2194
func Fn2194(m *base.Module, l0 int64) int64

//go:linkname Fn2195 github.com/goccy/perlwasm2go/p0.Fn2195
func Fn2195(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2196 github.com/goccy/perlwasm2go/p2.Fn2196
func Fn2196(m *base.Module, l0 int64) int64

//go:linkname Fn2198 github.com/goccy/perlwasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2201 github.com/goccy/perlwasm2go/p0.Fn2201
func Fn2201(m *base.Module) int64

//go:linkname Fn2202 github.com/goccy/perlwasm2go/p0.Fn2202
func Fn2202(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2203 github.com/goccy/perlwasm2go/p0.Fn2203
func Fn2203(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn2204 github.com/goccy/perlwasm2go/p2.Fn2204
func Fn2204(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2206 github.com/goccy/perlwasm2go/p0.Fn2206
func Fn2206(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2207 github.com/goccy/perlwasm2go/p0.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2208 github.com/goccy/perlwasm2go/p0.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2209 github.com/goccy/perlwasm2go/p0.Fn2209
func Fn2209(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2214 github.com/goccy/perlwasm2go/p2.Fn2214
func Fn2214(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2215 github.com/goccy/perlwasm2go/p2.Fn2215
func Fn2215(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2216 github.com/goccy/perlwasm2go/p0.Fn2216
func Fn2216(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2217 github.com/goccy/perlwasm2go/p2.Fn2217
func Fn2217(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2218 github.com/goccy/perlwasm2go/p0.Fn2218
func Fn2218(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2219 github.com/goccy/perlwasm2go/p2.Fn2219
func Fn2219(m *base.Module, l0 int64) int32

//go:linkname Fn2220 github.com/goccy/perlwasm2go/p2.Fn2220
func Fn2220(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2221 github.com/goccy/perlwasm2go/p0.Fn2221
func Fn2221(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2224 github.com/goccy/perlwasm2go/p2.Fn2224
func Fn2224(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2225 github.com/goccy/perlwasm2go/p2.Fn2225
func Fn2225(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn2230 github.com/goccy/perlwasm2go/p2.Fn2230
func Fn2230(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2231 github.com/goccy/perlwasm2go/p2.Fn2231
func Fn2231(m *base.Module, l0 int64) int64

//go:linkname Fn2232 github.com/goccy/perlwasm2go/p2.Fn2232
func Fn2232(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn2233 github.com/goccy/perlwasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int64) int64

//go:linkname Fn2234 github.com/goccy/perlwasm2go/p2.Fn2234
func Fn2234(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2236 github.com/goccy/perlwasm2go/p2.Fn2236
func Fn2236(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2237 github.com/goccy/perlwasm2go/p0.Fn2237
func Fn2237(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2244 github.com/goccy/perlwasm2go/p2.Fn2244
func Fn2244(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2245 github.com/goccy/perlwasm2go/p2.Fn2245
func Fn2245(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2246 github.com/goccy/perlwasm2go/p2.Fn2246
func Fn2246(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2247 github.com/goccy/perlwasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2249 github.com/goccy/perlwasm2go/p2.Fn2249
func Fn2249(m *base.Module, l0 int64) int64

//go:linkname Fn2250 github.com/goccy/perlwasm2go/p2.Fn2250
func Fn2250(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2251 github.com/goccy/perlwasm2go/p2.Fn2251
func Fn2251(m *base.Module, l0 int64) int64

//go:linkname Fn2252 github.com/goccy/perlwasm2go/p2.Fn2252
func Fn2252(m *base.Module, l0 int64) int64

//go:linkname Fn2253 github.com/goccy/perlwasm2go/p0.Fn2253
func Fn2253(m *base.Module, l0 int64) int64

//go:linkname Fn2254 github.com/goccy/perlwasm2go/p2.Fn2254
func Fn2254(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2255 github.com/goccy/perlwasm2go/p2.Fn2255
func Fn2255(m *base.Module, l0 int64) int64

//go:linkname Fn2256 github.com/goccy/perlwasm2go/p2.Fn2256
func Fn2256(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2260 github.com/goccy/perlwasm2go/p2.Fn2260
func Fn2260(m *base.Module, l0 int64) int64

//go:linkname Fn2261 github.com/goccy/perlwasm2go/p2.Fn2261
func Fn2261(m *base.Module, l0 int64)

//go:linkname Fn2262 github.com/goccy/perlwasm2go/p2.Fn2262
func Fn2262(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2263 github.com/goccy/perlwasm2go/p2.Fn2263
func Fn2263(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn2265 github.com/goccy/perlwasm2go/p2.Fn2265
func Fn2265(m *base.Module, l0 int64) int64

//go:linkname Fn2266 github.com/goccy/perlwasm2go/p2.Fn2266
func Fn2266(m *base.Module, l0 int64)

//go:linkname Fn2310 github.com/goccy/perlwasm2go/p2.Fn2310
func Fn2310(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2313 github.com/goccy/perlwasm2go/p2.Fn2313
func Fn2313(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2317 github.com/goccy/perlwasm2go/p0.Fn2317
func Fn2317(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2330 github.com/goccy/perlwasm2go/p2.Fn2330
func Fn2330(m *base.Module)

//go:linkname Fn2332 github.com/goccy/perlwasm2go/p2.Fn2332
func Fn2332(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2333 github.com/goccy/perlwasm2go/p2.Fn2333
func Fn2333(m *base.Module) int64

//go:linkname Fn2335 github.com/goccy/perlwasm2go/p2.Fn2335
func Fn2335(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn2336 github.com/goccy/perlwasm2go/p2.Fn2336
func Fn2336(m *base.Module) int64

//go:linkname Fn2339 github.com/goccy/perlwasm2go/p0.Fn2339
func Fn2339(m *base.Module, l0 int64) int64

//go:linkname Fn2341 github.com/goccy/perlwasm2go/p2.Fn2341
func Fn2341(m *base.Module, l0 int32) int64

//go:linkname Fn2343 github.com/goccy/perlwasm2go/p0.Fn2343
func Fn2343(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn2344 github.com/goccy/perlwasm2go/p0.Fn2344
func Fn2344(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2349 github.com/goccy/perlwasm2go/p2.Fn2349
func Fn2349(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn2350 github.com/goccy/perlwasm2go/p0.Fn2350
func Fn2350(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2352 github.com/goccy/perlwasm2go/p2.Fn2352
func Fn2352(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2353 github.com/goccy/perlwasm2go/p2.Fn2353
func Fn2353(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2354 github.com/goccy/perlwasm2go/p2.Fn2354
func Fn2354(m *base.Module)

//go:linkname Fn2355 github.com/goccy/perlwasm2go/p2.Fn2355
func Fn2355(m *base.Module)

//go:linkname Fn2356 github.com/goccy/perlwasm2go/p2.Fn2356
func Fn2356(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2357 github.com/goccy/perlwasm2go/p2.Fn2357
func Fn2357(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2359 github.com/goccy/perlwasm2go/p2.Fn2359
func Fn2359(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2361 github.com/goccy/perlwasm2go/p2.Fn2361
func Fn2361(m *base.Module)

//go:linkname Fn2362 github.com/goccy/perlwasm2go/p2.Fn2362
func Fn2362(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn2363 github.com/goccy/perlwasm2go/p2.Fn2363
func Fn2363(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2364 github.com/goccy/perlwasm2go/p0.Fn2364
func Fn2364(m *base.Module, l0 int32)

//go:linkname Fn2365 github.com/goccy/perlwasm2go/p2.Fn2365
func Fn2365(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2366 github.com/goccy/perlwasm2go/p0.Fn2366
func Fn2366(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2367 github.com/goccy/perlwasm2go/p0.Fn2367
func Fn2367(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2369 github.com/goccy/perlwasm2go/p2.Fn2369
func Fn2369(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2370 github.com/goccy/perlwasm2go/p2.Fn2370
func Fn2370(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2371 github.com/goccy/perlwasm2go/p0.Fn2371
func Fn2371(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2374 github.com/goccy/perlwasm2go/p2.Fn2374
func Fn2374(m *base.Module, l0 int64) int64

//go:linkname Fn2375 github.com/goccy/perlwasm2go/p2.Fn2375
func Fn2375(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2376 github.com/goccy/perlwasm2go/p2.Fn2376
func Fn2376(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2377 github.com/goccy/perlwasm2go/p2.Fn2377
func Fn2377(m *base.Module)

//go:linkname Fn2378 github.com/goccy/perlwasm2go/p2.Fn2378
func Fn2378(m *base.Module)

//go:linkname Fn2380 github.com/goccy/perlwasm2go/p0.Fn2380
func Fn2380(m *base.Module)

//go:linkname Fn2381 github.com/goccy/perlwasm2go/p2.Fn2381
func Fn2381(m *base.Module, l0 int64)

//go:linkname Fn2385 github.com/goccy/perlwasm2go/p0.Fn2385
func Fn2385(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2386 github.com/goccy/perlwasm2go/p0.Fn2386
func Fn2386(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2390 github.com/goccy/perlwasm2go/p0.Fn2390
func Fn2390(m *base.Module, l0 int64) int32

//go:linkname Fn2393 github.com/goccy/perlwasm2go/p2.Fn2393
func Fn2393(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2439 github.com/goccy/perlwasm2go/p0.Fn2439
func Fn2439(m *base.Module, l0 int64)

//go:linkname Fn2440 github.com/goccy/perlwasm2go/p0.Fn2440
func Fn2440(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2441 github.com/goccy/perlwasm2go/p0.Fn2441
func Fn2441(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2442 github.com/goccy/perlwasm2go/p0.Fn2442
func Fn2442(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2445 github.com/goccy/perlwasm2go/p0.Fn2445
func Fn2445(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2450 github.com/goccy/perlwasm2go/p0.Fn2450
func Fn2450(m *base.Module, l0 int64)

//go:linkname Fn2452 github.com/goccy/perlwasm2go/p2.Fn2452
func Fn2452(m *base.Module, l0 int64)

//go:linkname Fn2453 github.com/goccy/perlwasm2go/p2.Fn2453
func Fn2453(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2454 github.com/goccy/perlwasm2go/p0.Fn2454
func Fn2454(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2455 github.com/goccy/perlwasm2go/p0.Fn2455
func Fn2455(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2456 github.com/goccy/perlwasm2go/p2.Fn2456
func Fn2456(m *base.Module, l0 int64) int64

//go:linkname Fn2458 github.com/goccy/perlwasm2go/p0.Fn2458
func Fn2458(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2459 github.com/goccy/perlwasm2go/p0.Fn2459
func Fn2459(m *base.Module, l0 int64) int64

//go:linkname Fn2461 github.com/goccy/perlwasm2go/p0.Fn2461
func Fn2461(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2462 github.com/goccy/perlwasm2go/p2.Fn2462
func Fn2462(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2465 github.com/goccy/perlwasm2go/p2.Fn2465
func Fn2465(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2479 github.com/goccy/perlwasm2go/p0.Fn2479
func Fn2479(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn2493 github.com/goccy/perlwasm2go/p2.Fn2493
func Fn2493(m *base.Module)

//go:linkname Fn2494 github.com/goccy/perlwasm2go/p0.Fn2494
func Fn2494(m *base.Module, l0 int64)

//go:linkname Fn2495 github.com/goccy/perlwasm2go/p0.Fn2495
func Fn2495(m *base.Module, l0 int32) int64

//go:linkname Fn2499 github.com/goccy/perlwasm2go/p2.Fn2499
func Fn2499(m *base.Module, l0 int64)

//go:linkname Fn2502 github.com/goccy/perlwasm2go/p2.Fn2502
func Fn2502(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2504 github.com/goccy/perlwasm2go/p2.Fn2504
func Fn2504(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2505 github.com/goccy/perlwasm2go/p0.Fn2505
func Fn2505(m *base.Module, l0 int64)

//go:linkname Fn2506 github.com/goccy/perlwasm2go/p0.Fn2506
func Fn2506(m *base.Module, l0 int64) int64

//go:linkname Fn2508 github.com/goccy/perlwasm2go/p2.Fn2508
func Fn2508(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2512 github.com/goccy/perlwasm2go/p2.Fn2512
func Fn2512(m *base.Module, l0 int64)

//go:linkname Fn2517 github.com/goccy/perlwasm2go/p2.Fn2517
func Fn2517(m *base.Module, l0 int32)

//go:linkname Fn2521 github.com/goccy/perlwasm2go/p2.Fn2521
func Fn2521(m *base.Module, l0 int32) int32

//go:linkname Fn2527 github.com/goccy/perlwasm2go/p2.Fn2527
func Fn2527(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int32

//go:linkname Fn2529 github.com/goccy/perlwasm2go/p2.Fn2529
func Fn2529(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn2530 github.com/goccy/perlwasm2go/p2.Fn2530
func Fn2530(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32, l8 int32, l9 int32, l10 int32, l11 int64, l12 int64) int32

//go:linkname Fn2531 github.com/goccy/perlwasm2go/p2.Fn2531
func Fn2531(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2534 github.com/goccy/perlwasm2go/p2.Fn2534
func Fn2534(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2539 github.com/goccy/perlwasm2go/p0.Fn2539
func Fn2539(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2541 github.com/goccy/perlwasm2go/p2.Fn2541
func Fn2541(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn2545 github.com/goccy/perlwasm2go/p2.Fn2545
func Fn2545(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2546 github.com/goccy/perlwasm2go/p2.Fn2546
func Fn2546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2550 github.com/goccy/perlwasm2go/p2.Fn2550
func Fn2550(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2553 github.com/goccy/perlwasm2go/p0.Fn2553
func Fn2553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2554 github.com/goccy/perlwasm2go/p0.Fn2554
func Fn2554(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2555 github.com/goccy/perlwasm2go/p2.Fn2555
func Fn2555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2556 github.com/goccy/perlwasm2go/p2.Fn2556
func Fn2556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2558 github.com/goccy/perlwasm2go/p2.Fn2558
func Fn2558(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2559 github.com/goccy/perlwasm2go/p2.Fn2559
func Fn2559(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn2561 github.com/goccy/perlwasm2go/p2.Fn2561
func Fn2561(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64)

//go:linkname Fn2562 github.com/goccy/perlwasm2go/p2.Fn2562
func Fn2562(m *base.Module, l0 int64) int32

//go:linkname Fn2563 github.com/goccy/perlwasm2go/p2.Fn2563
func Fn2563(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2564 github.com/goccy/perlwasm2go/p2.Fn2564
func Fn2564(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2565 github.com/goccy/perlwasm2go/p2.Fn2565
func Fn2565(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn2566 github.com/goccy/perlwasm2go/p2.Fn2566
func Fn2566(m *base.Module, l0 int64) int64

//go:linkname Fn2567 github.com/goccy/perlwasm2go/p2.Fn2567
func Fn2567(m *base.Module, l0 int64) int64

//go:linkname Fn2569 github.com/goccy/perlwasm2go/p2.Fn2569
func Fn2569(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2570 github.com/goccy/perlwasm2go/p2.Fn2570
func Fn2570(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2571 github.com/goccy/perlwasm2go/p2.Fn2571
func Fn2571(m *base.Module, l0 int64) int64

//go:linkname Fn2572 github.com/goccy/perlwasm2go/p2.Fn2572
func Fn2572(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2576 github.com/goccy/perlwasm2go/p2.Fn2576
func Fn2576(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2578 github.com/goccy/perlwasm2go/p2.Fn2578
func Fn2578(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2581 github.com/goccy/perlwasm2go/p0.Fn2581
func Fn2581(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2583 github.com/goccy/perlwasm2go/p0.Fn2583
func Fn2583(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn2584 github.com/goccy/perlwasm2go/p2.Fn2584
func Fn2584(m *base.Module, l0 int64) int64

//go:linkname Fn2585 github.com/goccy/perlwasm2go/p2.Fn2585
func Fn2585(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2586 github.com/goccy/perlwasm2go/p0.Fn2586
func Fn2586(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2589 github.com/goccy/perlwasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2590 github.com/goccy/perlwasm2go/p0.Fn2590
func Fn2590(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2592 github.com/goccy/perlwasm2go/p0.Fn2592
func Fn2592(m *base.Module, l0 int64) int64

//go:linkname Fn2595 github.com/goccy/perlwasm2go/p0.Fn2595
func Fn2595(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2596 github.com/goccy/perlwasm2go/p0.Fn2596
func Fn2596(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn2598 github.com/goccy/perlwasm2go/p0.Fn2598
func Fn2598(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2600 github.com/goccy/perlwasm2go/p2.Fn2600
func Fn2600(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2601 github.com/goccy/perlwasm2go/p0.Fn2601
func Fn2601(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2603 github.com/goccy/perlwasm2go/p0.Fn2603
func Fn2603(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2605 github.com/goccy/perlwasm2go/p0.Fn2605
func Fn2605(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2606 github.com/goccy/perlwasm2go/p0.Fn2606
func Fn2606(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn2610 github.com/goccy/perlwasm2go/p0.Fn2610
func Fn2610(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn2611 github.com/goccy/perlwasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2612 github.com/goccy/perlwasm2go/p0.Fn2612
func Fn2612(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2613 github.com/goccy/perlwasm2go/p2.Fn2613
func Fn2613(m *base.Module, l0 int64)

//go:linkname Fn2615 github.com/goccy/perlwasm2go/p0.Fn2615
func Fn2615(m *base.Module, l0 int64) int64

//go:linkname Fn2616 github.com/goccy/perlwasm2go/p0.Fn2616
func Fn2616(m *base.Module, l0 int64)

//go:linkname Fn2618 github.com/goccy/perlwasm2go/p0.Fn2618
func Fn2618(m *base.Module, l0 int64) int32

//go:linkname Fn2619 github.com/goccy/perlwasm2go/p0.Fn2619
func Fn2619(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn2621 github.com/goccy/perlwasm2go/p0.Fn2621
func Fn2621(m *base.Module, l0 int32)

//go:linkname Fn2622 github.com/goccy/perlwasm2go/p2.Fn2622
func Fn2622(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2623 github.com/goccy/perlwasm2go/p0.Fn2623
func Fn2623(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2624 github.com/goccy/perlwasm2go/p2.Fn2624
func Fn2624(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2626 github.com/goccy/perlwasm2go/p0.Fn2626
func Fn2626(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int64, l7 int32) int64

//go:linkname Fn2629 github.com/goccy/perlwasm2go/p0.Fn2629
func Fn2629(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2640 github.com/goccy/perlwasm2go/p0.Fn2640
func Fn2640(m *base.Module, l0 int64)

//go:linkname Fn2641 github.com/goccy/perlwasm2go/p0.Fn2641
func Fn2641(m *base.Module, l0 int64)

//go:linkname Fn2646 github.com/goccy/perlwasm2go/p0.Fn2646
func Fn2646(m *base.Module, l0 int64) int32

//go:linkname Fn2647 github.com/goccy/perlwasm2go/p0.Fn2647
func Fn2647(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2648 github.com/goccy/perlwasm2go/p0.Fn2648
func Fn2648(m *base.Module, l0 int64)

//go:linkname Fn2649 github.com/goccy/perlwasm2go/p0.Fn2649
func Fn2649(m *base.Module, l0 int64) int64

//go:linkname Fn2650 github.com/goccy/perlwasm2go/p0.Fn2650
func Fn2650(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2653 github.com/goccy/perlwasm2go/p0.Fn2653
func Fn2653(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2654 github.com/goccy/perlwasm2go/p2.Fn2654
func Fn2654(m *base.Module, l0 int64) int64

//go:linkname Fn2655 github.com/goccy/perlwasm2go/p0.Fn2655
func Fn2655(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2656 github.com/goccy/perlwasm2go/p0.Fn2656
func Fn2656(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2659 github.com/goccy/perlwasm2go/p0.Fn2659
func Fn2659(m *base.Module, l0 int64)

//go:linkname Fn2663 github.com/goccy/perlwasm2go/p0.Fn2663
func Fn2663(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2664 github.com/goccy/perlwasm2go/p0.Fn2664
func Fn2664(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2665 github.com/goccy/perlwasm2go/p0.Fn2665
func Fn2665(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2667 github.com/goccy/perlwasm2go/p0.Fn2667
func Fn2667(m *base.Module, l0 int64) int64

//go:linkname Fn2670 github.com/goccy/perlwasm2go/p2.Fn2670
func Fn2670(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2672 github.com/goccy/perlwasm2go/p0.Fn2672
func Fn2672(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2676 github.com/goccy/perlwasm2go/p0.Fn2676
func Fn2676(m *base.Module, l0 int64)

//go:linkname Fn2679 github.com/goccy/perlwasm2go/p0.Fn2679
func Fn2679(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn2680 github.com/goccy/perlwasm2go/p0.Fn2680
func Fn2680(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn2681 github.com/goccy/perlwasm2go/p0.Fn2681
func Fn2681(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2683 github.com/goccy/perlwasm2go/p0.Fn2683
func Fn2683(m *base.Module, l0 int64) int64

//go:linkname Fn2685 github.com/goccy/perlwasm2go/p2.Fn2685
func Fn2685(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2686 github.com/goccy/perlwasm2go/p0.Fn2686
func Fn2686(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2687 github.com/goccy/perlwasm2go/p2.Fn2687
func Fn2687(m *base.Module)

//go:linkname Fn2689 github.com/goccy/perlwasm2go/p2.Fn2689
func Fn2689(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn2694 github.com/goccy/perlwasm2go/p2.Fn2694
func Fn2694(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2696 github.com/goccy/perlwasm2go/p2.Fn2696
func Fn2696(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn2697 github.com/goccy/perlwasm2go/p2.Fn2697
func Fn2697(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn2700 github.com/goccy/perlwasm2go/p2.Fn2700
func Fn2700(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2701 github.com/goccy/perlwasm2go/p2.Fn2701
func Fn2701(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2704 github.com/goccy/perlwasm2go/p0.Fn2704
func Fn2704(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2710 github.com/goccy/perlwasm2go/p0.Fn2710
func Fn2710(m *base.Module, l0 int64) int32

//go:linkname Fn2714 github.com/goccy/perlwasm2go/p0.Fn2714
func Fn2714(m *base.Module, l0 int64)

//go:linkname Fn2715 github.com/goccy/perlwasm2go/p0.Fn2715
func Fn2715(m *base.Module, l0 int64) int32

//go:linkname Fn2716 github.com/goccy/perlwasm2go/p0.Fn2716
func Fn2716(m *base.Module, l0 int64) int32

//go:linkname Fn2717 github.com/goccy/perlwasm2go/p2.Fn2717
func Fn2717(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2718 github.com/goccy/perlwasm2go/p2.Fn2718
func Fn2718(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int64

//go:linkname Fn2720 github.com/goccy/perlwasm2go/p2.Fn2720
func Fn2720(m *base.Module, l0 int64) int64

//go:linkname Fn2721 github.com/goccy/perlwasm2go/p2.Fn2721
func Fn2721(m *base.Module, l0 int64)

//go:linkname Fn2722 github.com/goccy/perlwasm2go/p0.Fn2722
func Fn2722(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2725 github.com/goccy/perlwasm2go/p0.Fn2725
func Fn2725(m *base.Module, l0 int64) int32

//go:linkname Fn2730 github.com/goccy/perlwasm2go/p2.Fn2730
func Fn2730(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2731 github.com/goccy/perlwasm2go/p2.Fn2731
func Fn2731(m *base.Module, l0 int32) int64

//go:linkname Fn2739 github.com/goccy/perlwasm2go/p2.Fn2739
func Fn2739(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2745 github.com/goccy/perlwasm2go/p2.Fn2745
func Fn2745(m *base.Module, l0 int32)

//go:linkname Fn2748 github.com/goccy/perlwasm2go/p0.Fn2748
func Fn2748(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn2799 github.com/goccy/perlwasm2go/p2.Fn2799
func Fn2799(m *base.Module, l0 int64) int64

//go:linkname Fn2809 github.com/goccy/perlwasm2go/p2.Fn2809
func Fn2809(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2810 github.com/goccy/perlwasm2go/p2.Fn2810
func Fn2810(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2811 github.com/goccy/perlwasm2go/p2.Fn2811
func Fn2811(m *base.Module, l0 int64) int64

//go:linkname Fn2812 github.com/goccy/perlwasm2go/p2.Fn2812
func Fn2812(m *base.Module, l0 int64)

//go:linkname Fn2813 github.com/goccy/perlwasm2go/p0.Fn2813
func Fn2813(m *base.Module, l0 int64) int64

//go:linkname Fn2816 github.com/goccy/perlwasm2go/p0.Fn2816
func Fn2816(m *base.Module, l0 int64)

//go:linkname Fn2818 github.com/goccy/perlwasm2go/p0.Fn2818
func Fn2818(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2821 github.com/goccy/perlwasm2go/p0.Fn2821
func Fn2821(m *base.Module, l0 int64)

//go:linkname Fn2827 github.com/goccy/perlwasm2go/p0.Fn2827
func Fn2827(m *base.Module, l0 int32)

//go:linkname Fn2828 github.com/goccy/perlwasm2go/p0.Fn2828
func Fn2828(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int64

//go:linkname Fn2829 github.com/goccy/perlwasm2go/p0.Fn2829
func Fn2829(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int64

//go:linkname Fn2833 github.com/goccy/perlwasm2go/p0.Fn2833
func Fn2833(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2834 github.com/goccy/perlwasm2go/p0.Fn2834
func Fn2834(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2836 github.com/goccy/perlwasm2go/p0.Fn2836
func Fn2836(m *base.Module, l0 int64) float64

//go:linkname Fn2840 github.com/goccy/perlwasm2go/p0.Fn2840
func Fn2840(m *base.Module, l0 int32) int64

//go:linkname Fn2841 github.com/goccy/perlwasm2go/p0.Fn2841
func Fn2841(m *base.Module, l0 int32) int64

//go:linkname Fn2842 github.com/goccy/perlwasm2go/p0.Fn2842
func Fn2842(m *base.Module, l0 int64) int64

//go:linkname Fn2844 github.com/goccy/perlwasm2go/p0.Fn2844
func Fn2844(m *base.Module, l0 int64)

//go:linkname Fn2846 github.com/goccy/perlwasm2go/p0.Fn2846
func Fn2846(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2847 github.com/goccy/perlwasm2go/p0.Fn2847
func Fn2847(m *base.Module, l0 int64)

//go:linkname Fn2848 github.com/goccy/perlwasm2go/p0.Fn2848
func Fn2848(m *base.Module, l0 int64)

//go:linkname Fn2850 github.com/goccy/perlwasm2go/p0.Fn2850
func Fn2850(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2852 github.com/goccy/perlwasm2go/p0.Fn2852
func Fn2852(m *base.Module, l0 int32) int64

//go:linkname Fn2856 github.com/goccy/perlwasm2go/p2.Fn2856
func Fn2856(m *base.Module, l0 int64) int64

//go:linkname Fn2857 github.com/goccy/perlwasm2go/p0.Fn2857
func Fn2857(m *base.Module, l0 int64)

//go:linkname Fn2858 github.com/goccy/perlwasm2go/p0.Fn2858
func Fn2858(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2860 github.com/goccy/perlwasm2go/p0.Fn2860
func Fn2860(m *base.Module, l0 int64)

//go:linkname Fn2861 github.com/goccy/perlwasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int32) int64

//go:linkname Fn2862 github.com/goccy/perlwasm2go/p0.Fn2862
func Fn2862(m *base.Module, l0 int64) int64

//go:linkname Fn2865 github.com/goccy/perlwasm2go/p2.Fn2865
func Fn2865(m *base.Module, l0 int64) int64

//go:linkname Fn2867 github.com/goccy/perlwasm2go/p2.Fn2867
func Fn2867(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int64

//go:linkname Fn2868 github.com/goccy/perlwasm2go/p2.Fn2868
func Fn2868(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2869 github.com/goccy/perlwasm2go/p0.Fn2869
func Fn2869(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2870 github.com/goccy/perlwasm2go/p0.Fn2870
func Fn2870(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2871 github.com/goccy/perlwasm2go/p2.Fn2871
func Fn2871(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2874 github.com/goccy/perlwasm2go/p2.Fn2874
func Fn2874(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2876 github.com/goccy/perlwasm2go/p0.Fn2876
func Fn2876(m *base.Module, l0 int64)

//go:linkname Fn2878 github.com/goccy/perlwasm2go/p0.Fn2878
func Fn2878(m *base.Module, l0 int64)

//go:linkname Fn2893 github.com/goccy/perlwasm2go/p2.Fn2893
func Fn2893(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2896 github.com/goccy/perlwasm2go/p0.Fn2896
func Fn2896(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2898 github.com/goccy/perlwasm2go/p2.Fn2898
func Fn2898(m *base.Module, l0 int64) int32

//go:linkname Fn2899 github.com/goccy/perlwasm2go/p2.Fn2899
func Fn2899(m *base.Module, l0 int64)

//go:linkname Fn2903 github.com/goccy/perlwasm2go/p2.Fn2903
func Fn2903(m *base.Module, l0 int64) int32

//go:linkname Fn2904 github.com/goccy/perlwasm2go/p2.Fn2904
func Fn2904(m *base.Module, l0 int32) int64

//go:linkname Fn2920 github.com/goccy/perlwasm2go/p0.Fn2920
func Fn2920(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2921 github.com/goccy/perlwasm2go/p0.Fn2921
func Fn2921(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int64) int64

//go:linkname Fn2925 github.com/goccy/perlwasm2go/p0.Fn2925
func Fn2925(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2926 github.com/goccy/perlwasm2go/p2.Fn2926
func Fn2926(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2928 github.com/goccy/perlwasm2go/p2.Fn2928
func Fn2928(m *base.Module, l0 int64) int32

//go:linkname Fn2933 github.com/goccy/perlwasm2go/p2.Fn2933
func Fn2933(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2935 github.com/goccy/perlwasm2go/p2.Fn2935
func Fn2935(m *base.Module, l0 int64) int64

//go:linkname Fn2936 github.com/goccy/perlwasm2go/p2.Fn2936
func Fn2936(m *base.Module, l0 int64) int64

//go:linkname Fn2937 github.com/goccy/perlwasm2go/p2.Fn2937
func Fn2937(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2938 github.com/goccy/perlwasm2go/p2.Fn2938
func Fn2938(m *base.Module, l0 int64) int32

//go:linkname Fn2939 github.com/goccy/perlwasm2go/p2.Fn2939
func Fn2939(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2941 github.com/goccy/perlwasm2go/p2.Fn2941
func Fn2941(m *base.Module, l0 int64) int64

//go:linkname Fn2943 github.com/goccy/perlwasm2go/p2.Fn2943
func Fn2943(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2946 github.com/goccy/perlwasm2go/p2.Fn2946
func Fn2946(m *base.Module, l0 int64) int32

//go:linkname Fn2948 github.com/goccy/perlwasm2go/p2.Fn2948
func Fn2948(m *base.Module, l0 int64) int32

//go:linkname Fn2950 github.com/goccy/perlwasm2go/p2.Fn2950
func Fn2950(m *base.Module, l0 int64)

//go:linkname Fn2958 github.com/goccy/perlwasm2go/p2.Fn2958
func Fn2958(m *base.Module, l0 int64) int64

//go:linkname Fn2959 github.com/goccy/perlwasm2go/p2.Fn2959
func Fn2959(m *base.Module, l0 int64) int64

//go:linkname Fn2965 github.com/goccy/perlwasm2go/p2.Fn2965
func Fn2965(m *base.Module)

//go:linkname Fn3030 github.com/goccy/perlwasm2go/p0.Fn3030
func Fn3030(m *base.Module) int64

//go:linkname Fn3031 github.com/goccy/perlwasm2go/p0.Fn3031
func Fn3031(m *base.Module) int64

//go:linkname Fn3032 github.com/goccy/perlwasm2go/p0.Fn3032
func Fn3032(m *base.Module) int64

//go:linkname Fn3033 github.com/goccy/perlwasm2go/p2.Fn3033
func Fn3033(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3034 github.com/goccy/perlwasm2go/p2.Fn3034
func Fn3034(m *base.Module, l0 int64) int32

//go:linkname Fn3035 github.com/goccy/perlwasm2go/p2.Fn3035
func Fn3035(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn3036 github.com/goccy/perlwasm2go/p2.Fn3036
func Fn3036(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn3037 github.com/goccy/perlwasm2go/p2.Fn3037
func Fn3037(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3039 github.com/goccy/perlwasm2go/p0.Fn3039
func Fn3039(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3060 github.com/goccy/perlwasm2go/p2.Fn3060
func Fn3060(m *base.Module)

//go:linkname Fn3062 github.com/goccy/perlwasm2go/p0.Fn3062
func Fn3062(m *base.Module, l0 int64)

//go:linkname Fn3066 github.com/goccy/perlwasm2go/p2.Fn3066
func Fn3066(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn3067 github.com/goccy/perlwasm2go/p0.Fn3067
func Fn3067(m *base.Module, l0 int32) int64

//go:linkname Fn3073 github.com/goccy/perlwasm2go/p2.Fn3073
func Fn3073(m *base.Module, l0 int32) int64

//go:linkname Fn3112 github.com/goccy/perlwasm2go/p2.Fn3112
func Fn3112(m *base.Module) int64

//go:linkname Fn3138 github.com/goccy/perlwasm2go/p2.Fn3138
func Fn3138(m *base.Module) int64

//go:linkname Fn3152 github.com/goccy/perlwasm2go/p2.Fn3152
func Fn3152(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3158 github.com/goccy/perlwasm2go/p2.Fn3158
func Fn3158(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3168 github.com/goccy/perlwasm2go/p2.Fn3168
func Fn3168(m *base.Module, l0 int64) int32

//go:linkname Fn3191 github.com/goccy/perlwasm2go/p2.Fn3191
func Fn3191(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3196 github.com/goccy/perlwasm2go/p2.Fn3196
func Fn3196(m *base.Module)

//go:linkname Fn3208 github.com/goccy/perlwasm2go/p2.Fn3208
func Fn3208(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3211 github.com/goccy/perlwasm2go/p2.Fn3211
func Fn3211(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3212 github.com/goccy/perlwasm2go/p2.Fn3212
func Fn3212(m *base.Module) int64

//go:linkname Fn3214 github.com/goccy/perlwasm2go/p2.Fn3214
func Fn3214(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3228 github.com/goccy/perlwasm2go/p2.Fn3228
func Fn3228(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn3233 github.com/goccy/perlwasm2go/p0.Fn3233
func Fn3233(m *base.Module, l0 int64)

//go:linkname Fn3234 github.com/goccy/perlwasm2go/p2.Fn3234
func Fn3234(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3244 github.com/goccy/perlwasm2go/p2.Fn3244
func Fn3244(m *base.Module) int32

//go:linkname Fn3245 github.com/goccy/perlwasm2go/p0.Fn3245
func Fn3245(m *base.Module, l0 int32)

//go:linkname Fn3254 github.com/goccy/perlwasm2go/p0.Fn3254
func Fn3254(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3257 github.com/goccy/perlwasm2go/p2.Fn3257
func Fn3257(m *base.Module, l0 int32) int64

//go:linkname Fn3260 github.com/goccy/perlwasm2go/p0.Fn3260
func Fn3260(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64

//go:linkname Fn3265 github.com/goccy/perlwasm2go/p2.Fn3265
func Fn3265(m *base.Module, l0 int64) int32

//go:linkname Fn3278 github.com/goccy/perlwasm2go/p2.Fn3278
func Fn3278(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64) int64

//go:linkname Fn3279 github.com/goccy/perlwasm2go/p2.Fn3279
func Fn3279(m *base.Module, l0 int64)

//go:linkname Fn3282 github.com/goccy/perlwasm2go/p2.Fn3282
func Fn3282(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn3284 github.com/goccy/perlwasm2go/p2.Fn3284
func Fn3284(m *base.Module, l0 int64)

//go:linkname Fn3285 github.com/goccy/perlwasm2go/p2.Fn3285
func Fn3285(m *base.Module, l0 int64) int64

//go:linkname Fn3302 github.com/goccy/perlwasm2go/p2.Fn3302
func Fn3302(m *base.Module)

//go:linkname Fn3303 github.com/goccy/perlwasm2go/p2.Fn3303
func Fn3303(m *base.Module, l0 int32) int64

//go:linkname Fn3304 github.com/goccy/perlwasm2go/p2.Fn3304
func Fn3304(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3305 github.com/goccy/perlwasm2go/p2.Fn3305
func Fn3305(m *base.Module, l0 int64) int64

//go:linkname Fn3306 github.com/goccy/perlwasm2go/p2.Fn3306
func Fn3306(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3307 github.com/goccy/perlwasm2go/p2.Fn3307
func Fn3307(m *base.Module)

//go:linkname Fn3332 github.com/goccy/perlwasm2go/p2.Fn3332
func Fn3332(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3334 github.com/goccy/perlwasm2go/p2.Fn3334
func Fn3334(m *base.Module, l0 int32) int64

//go:linkname Fn3335 github.com/goccy/perlwasm2go/p0.Fn3335
func Fn3335(m *base.Module, l0 int32) int64

//go:linkname Fn3339 github.com/goccy/perlwasm2go/p2.Fn3339
func Fn3339(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3342 github.com/goccy/perlwasm2go/p2.Fn3342
func Fn3342(m *base.Module) int64

//go:linkname Fn3358 github.com/goccy/perlwasm2go/p2.Fn3358
func Fn3358(m *base.Module, l0 int64) int32

//go:linkname Fn3361 github.com/goccy/perlwasm2go/p2.Fn3361
func Fn3361(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3367 github.com/goccy/perlwasm2go/p0.Fn3367
func Fn3367(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3369 github.com/goccy/perlwasm2go/p2.Fn3369
func Fn3369(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3370 github.com/goccy/perlwasm2go/p2.Fn3370
func Fn3370(m *base.Module, l0 int64)

//go:linkname Fn3371 github.com/goccy/perlwasm2go/p2.Fn3371
func Fn3371(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3382 github.com/goccy/perlwasm2go/p2.Fn3382
func Fn3382(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64) int32

//go:linkname Fn3383 github.com/goccy/perlwasm2go/p2.Fn3383
func Fn3383(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3385 github.com/goccy/perlwasm2go/p2.Fn3385
func Fn3385(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3389 github.com/goccy/perlwasm2go/p2.Fn3389
func Fn3389(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3390 github.com/goccy/perlwasm2go/p2.Fn3390
func Fn3390(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3391 github.com/goccy/perlwasm2go/p2.Fn3391
func Fn3391(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3395 github.com/goccy/perlwasm2go/p2.Fn3395
func Fn3395(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3396 github.com/goccy/perlwasm2go/p2.Fn3396
func Fn3396(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3400 github.com/goccy/perlwasm2go/p2.Fn3400
func Fn3400(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3424 github.com/goccy/perlwasm2go/p2.Fn3424
func Fn3424(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int64) int64

//go:linkname Fn3432 github.com/goccy/perlwasm2go/p2.Fn3432
func Fn3432(m *base.Module, l0 int32) int64

//go:linkname Fn3437 github.com/goccy/perlwasm2go/p0.Fn3437
func Fn3437(m *base.Module, l0 int64)

//go:linkname Fn3442 github.com/goccy/perlwasm2go/p2.Fn3442
func Fn3442(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3463 github.com/goccy/perlwasm2go/p2.Fn3463
func Fn3463(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3465 github.com/goccy/perlwasm2go/p2.Fn3465
func Fn3465(m *base.Module, l0 int32) int64

//go:linkname Fn3466 github.com/goccy/perlwasm2go/p2.Fn3466
func Fn3466(m *base.Module, l0 int64) int64

//go:linkname Fn3467 github.com/goccy/perlwasm2go/p2.Fn3467
func Fn3467(m *base.Module, l0 int64) int64

//go:linkname Fn3507 github.com/goccy/perlwasm2go/p2.Fn3507
func Fn3507(m *base.Module, l0 int64) int64

//go:linkname Fn3517 github.com/goccy/perlwasm2go/p2.Fn3517
func Fn3517(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3518 github.com/goccy/perlwasm2go/p2.Fn3518
func Fn3518(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn3519 github.com/goccy/perlwasm2go/p0.Fn3519
func Fn3519(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3520 github.com/goccy/perlwasm2go/p2.Fn3520
func Fn3520(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3522 github.com/goccy/perlwasm2go/p2.Fn3522
func Fn3522(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3523 github.com/goccy/perlwasm2go/p2.Fn3523
func Fn3523(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3524 github.com/goccy/perlwasm2go/p2.Fn3524
func Fn3524(m *base.Module)

//go:linkname Fn3525 github.com/goccy/perlwasm2go/p2.Fn3525
func Fn3525(m *base.Module) int64

//go:linkname Fn3528 github.com/goccy/perlwasm2go/p0.Fn3528
func Fn3528(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32, l7 int32) int64

//go:linkname Fn3529 github.com/goccy/perlwasm2go/p2.Fn3529
func Fn3529(m *base.Module, l0 int32) int64

//go:linkname Fn3531 github.com/goccy/perlwasm2go/p0.Fn3531
func Fn3531(m *base.Module, l0 int32) int64

//go:linkname Fn3533 github.com/goccy/perlwasm2go/p2.Fn3533
func Fn3533(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3536 github.com/goccy/perlwasm2go/p2.Fn3536
func Fn3536(m *base.Module, l0 int64) int64

//go:linkname Fn3541 github.com/goccy/perlwasm2go/p2.Fn3541
func Fn3541(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3545 github.com/goccy/perlwasm2go/p2.Fn3545
func Fn3545(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn3552 github.com/goccy/perlwasm2go/p2.Fn3552
func Fn3552(m *base.Module, l0 int32) int64

//go:linkname Fn3554 github.com/goccy/perlwasm2go/p2.Fn3554
func Fn3554(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3568 github.com/goccy/perlwasm2go/p2.Fn3568
func Fn3568(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3575 github.com/goccy/perlwasm2go/p0.Fn3575
func Fn3575(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3585 github.com/goccy/perlwasm2go/p2.Fn3585
func Fn3585(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3590 github.com/goccy/perlwasm2go/p2.Fn3590
func Fn3590(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3591 github.com/goccy/perlwasm2go/p2.Fn3591
func Fn3591(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3599 github.com/goccy/perlwasm2go/p0.Fn3599
func Fn3599(m *base.Module, l0 int64) int64

//go:linkname Fn3600 github.com/goccy/perlwasm2go/p2.Fn3600
func Fn3600(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3602 github.com/goccy/perlwasm2go/p2.Fn3602
func Fn3602(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3605 github.com/goccy/perlwasm2go/p2.Fn3605
func Fn3605(m *base.Module, l0 int64) int64

//go:linkname Fn3606 github.com/goccy/perlwasm2go/p2.Fn3606
func Fn3606(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3607 github.com/goccy/perlwasm2go/p2.Fn3607
func Fn3607(m *base.Module, l0 int64)

//go:linkname Fn3608 github.com/goccy/perlwasm2go/p0.Fn3608
func Fn3608(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3609 github.com/goccy/perlwasm2go/p2.Fn3609
func Fn3609(m *base.Module, l0 int64)

//go:linkname Fn3610 github.com/goccy/perlwasm2go/p2.Fn3610
func Fn3610(m *base.Module, l0 int64)

//go:linkname Fn3611 github.com/goccy/perlwasm2go/p2.Fn3611
func Fn3611(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn3630 github.com/goccy/perlwasm2go/p2.Fn3630
func Fn3630(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3633 github.com/goccy/perlwasm2go/p2.Fn3633
func Fn3633(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3635 github.com/goccy/perlwasm2go/p2.Fn3635
func Fn3635(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn3636 github.com/goccy/perlwasm2go/p2.Fn3636
func Fn3636(m *base.Module) int64

//go:linkname Fn3638 github.com/goccy/perlwasm2go/p2.Fn3638
func Fn3638(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3640 github.com/goccy/perlwasm2go/p2.Fn3640
func Fn3640(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3641 github.com/goccy/perlwasm2go/p2.Fn3641
func Fn3641(m *base.Module, l0 int64) int32

//go:linkname Fn3642 github.com/goccy/perlwasm2go/p2.Fn3642
func Fn3642(m *base.Module, l0 int64)

//go:linkname Fn3643 github.com/goccy/perlwasm2go/p2.Fn3643
func Fn3643(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3646 github.com/goccy/perlwasm2go/p2.Fn3646
func Fn3646(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3649 github.com/goccy/perlwasm2go/p2.Fn3649
func Fn3649(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3654 github.com/goccy/perlwasm2go/p2.Fn3654
func Fn3654(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int32) int32

//go:linkname Fn3655 github.com/goccy/perlwasm2go/p2.Fn3655
func Fn3655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32) int32

//go:linkname Fn3656 github.com/goccy/perlwasm2go/p2.Fn3656
func Fn3656(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn3657 github.com/goccy/perlwasm2go/p2.Fn3657
func Fn3657(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3658 github.com/goccy/perlwasm2go/p2.Fn3658
func Fn3658(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3659 github.com/goccy/perlwasm2go/p2.Fn3659
func Fn3659(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn3660 github.com/goccy/perlwasm2go/p2.Fn3660
func Fn3660(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3661 github.com/goccy/perlwasm2go/p2.Fn3661
func Fn3661(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3663 github.com/goccy/perlwasm2go/p2.Fn3663
func Fn3663(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3665 github.com/goccy/perlwasm2go/p2.Fn3665
func Fn3665(m *base.Module, l0 int32) int64

//go:linkname Fn3666 github.com/goccy/perlwasm2go/p2.Fn3666
func Fn3666(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn3667 github.com/goccy/perlwasm2go/p2.Fn3667
func Fn3667(m *base.Module, l0 int64) int64

//go:linkname Fn3672 github.com/goccy/perlwasm2go/p2.Fn3672
func Fn3672(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32

//go:linkname Fn3673 github.com/goccy/perlwasm2go/p2.Fn3673
func Fn3673(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3682 github.com/goccy/perlwasm2go/p2.Fn3682
func Fn3682(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3683 github.com/goccy/perlwasm2go/p2.Fn3683
func Fn3683(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3688 github.com/goccy/perlwasm2go/p0.Fn3688
func Fn3688(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3689 github.com/goccy/perlwasm2go/p0.Fn3689
func Fn3689(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn3691 github.com/goccy/perlwasm2go/p0.Fn3691
func Fn3691(m *base.Module) int32

//go:linkname Fn3692 github.com/goccy/perlwasm2go/p0.Fn3692
func Fn3692(m *base.Module)

//go:linkname Fn3693 github.com/goccy/perlwasm2go/p0.Fn3693
func Fn3693(m *base.Module)

//go:linkname Fn3694 github.com/goccy/perlwasm2go/p0.Fn3694
func Fn3694(m *base.Module, l0 int32)

//go:linkname Fn3695 github.com/goccy/perlwasm2go/p0.Fn3695
func Fn3695(m *base.Module, l0 int32)

//go:linkname Fn3696 github.com/goccy/perlwasm2go/p0.Fn3696
func Fn3696(m *base.Module) int64

//go:linkname Fn3698 github.com/goccy/perlwasm2go/p0.Fn3698
func Fn3698(m *base.Module, l0 int64) int64

//go:linkname Fn3699 github.com/goccy/perlwasm2go/p0.Fn3699
func Fn3699(m *base.Module)

//go:linkname Fn3700 github.com/goccy/perlwasm2go/p0.Fn3700
func Fn3700(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3701 github.com/goccy/perlwasm2go/p0.Fn3701
func Fn3701(m *base.Module, l0 int64) int64

//go:linkname Fn3702 github.com/goccy/perlwasm2go/p0.Fn3702
func Fn3702(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3703 github.com/goccy/perlwasm2go/p0.Fn3703
func Fn3703(m *base.Module, l0 int64)

//go:linkname Fn3704 github.com/goccy/perlwasm2go/p0.Fn3704
func Fn3704(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3705 github.com/goccy/perlwasm2go/p2.Fn3705
func Fn3705(m *base.Module, l0 int64)

//go:linkname Fn3706 github.com/goccy/perlwasm2go/p2.Fn3706
func Fn3706(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3707 github.com/goccy/perlwasm2go/p2.Fn3707
func Fn3707(m *base.Module, l0 int64) int64

//go:linkname Fn3709 github.com/goccy/perlwasm2go/p0.Fn3709
func Fn3709(m *base.Module, l0 int64)

//go:linkname Fn3710 github.com/goccy/perlwasm2go/p0.Fn3710
func Fn3710(m *base.Module, l0 int64)

//go:linkname Fn3711 github.com/goccy/perlwasm2go/p0.Fn3711
func Fn3711(m *base.Module, l0 int64)

//go:linkname Fn3712 github.com/goccy/perlwasm2go/p0.Fn3712
func Fn3712(m *base.Module, l0 int64)

//go:linkname Fn3713 github.com/goccy/perlwasm2go/p2.Fn3713
func Fn3713(m *base.Module, l0 int64)

//go:linkname Fn3714 github.com/goccy/perlwasm2go/p0.Fn3714
func Fn3714(m *base.Module, l0 int64)

//go:linkname Fn3715 github.com/goccy/perlwasm2go/p0.Fn3715
func Fn3715(m *base.Module, l0 int64)

//go:linkname Fn3716 github.com/goccy/perlwasm2go/p2.Fn3716
func Fn3716(m *base.Module, l0 int64)

//go:linkname Fn3717 github.com/goccy/perlwasm2go/p2.Fn3717
func Fn3717(m *base.Module, l0 int64)

//go:linkname Fn3718 github.com/goccy/perlwasm2go/p0.Fn3718
func Fn3718(m *base.Module, l0 int64)

//go:linkname Fn3719 github.com/goccy/perlwasm2go/p0.Fn3719
func Fn3719(m *base.Module, l0 int64)

//go:linkname Fn3720 github.com/goccy/perlwasm2go/p2.Fn3720
func Fn3720(m *base.Module, l0 int64)

//go:linkname Fn3721 github.com/goccy/perlwasm2go/p2.Fn3721
func Fn3721(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3724 github.com/goccy/perlwasm2go/p2.Fn3724
func Fn3724(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3725 github.com/goccy/perlwasm2go/p2.Fn3725
func Fn3725(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3726 github.com/goccy/perlwasm2go/p0.Fn3726
func Fn3726(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3727 github.com/goccy/perlwasm2go/p0.Fn3727
func Fn3727(m *base.Module)

//go:linkname Fn3728 github.com/goccy/perlwasm2go/p2.Fn3728
func Fn3728(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3729 github.com/goccy/perlwasm2go/p2.Fn3729
func Fn3729(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3730 github.com/goccy/perlwasm2go/p0.Fn3730
func Fn3730(m *base.Module)

//go:linkname Fn3732 github.com/goccy/perlwasm2go/p2.Fn3732
func Fn3732(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3734 github.com/goccy/perlwasm2go/p0.Fn3734
func Fn3734(m *base.Module) int64

//go:linkname Fn3736 github.com/goccy/perlwasm2go/p2.Fn3736
func Fn3736(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn3738 github.com/goccy/perlwasm2go/p2.Fn3738
func Fn3738(m *base.Module, l0 int64)

//go:linkname Fn3742 github.com/goccy/perlwasm2go/p0.Fn3742
func Fn3742(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3744 github.com/goccy/perlwasm2go/p0.Fn3744
func Fn3744(m *base.Module, l0 int32) int64

//go:linkname Fn3745 github.com/goccy/perlwasm2go/p0.Fn3745
func Fn3745(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3746 github.com/goccy/perlwasm2go/p0.Fn3746
func Fn3746(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3748 github.com/goccy/perlwasm2go/p0.Fn3748
func Fn3748(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3750 github.com/goccy/perlwasm2go/p0.Fn3750
func Fn3750(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3751 github.com/goccy/perlwasm2go/p0.Fn3751
func Fn3751(m *base.Module) int64

//go:linkname Fn3752 github.com/goccy/perlwasm2go/p0.Fn3752
func Fn3752(m *base.Module, l0 int64) int64

//go:linkname Fn3753 github.com/goccy/perlwasm2go/p0.Fn3753
func Fn3753(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3754 github.com/goccy/perlwasm2go/p0.Fn3754
func Fn3754(m *base.Module, l0 int32) int64

//go:linkname Fn3755 github.com/goccy/perlwasm2go/p0.Fn3755
func Fn3755(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn3757 github.com/goccy/perlwasm2go/p0.Fn3757
func Fn3757(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3758 github.com/goccy/perlwasm2go/p0.Fn3758
func Fn3758(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3759 github.com/goccy/perlwasm2go/p0.Fn3759
func Fn3759(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3760 github.com/goccy/perlwasm2go/p0.Fn3760
func Fn3760(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3761 github.com/goccy/perlwasm2go/p0.Fn3761
func Fn3761(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32)

//go:linkname Fn3762 github.com/goccy/perlwasm2go/p0.Fn3762
func Fn3762(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn3763 github.com/goccy/perlwasm2go/p2.Fn3763
func Fn3763(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3764 github.com/goccy/perlwasm2go/p0.Fn3764
func Fn3764(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3765 github.com/goccy/perlwasm2go/p0.Fn3765
func Fn3765(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3766 github.com/goccy/perlwasm2go/p0.Fn3766
func Fn3766(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn3767 github.com/goccy/perlwasm2go/p2.Fn3767
func Fn3767(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn3770 github.com/goccy/perlwasm2go/p2.Fn3770
func Fn3770(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3771 github.com/goccy/perlwasm2go/p2.Fn3771
func Fn3771(m *base.Module, l0 int64) int32

//go:linkname Fn3772 github.com/goccy/perlwasm2go/p0.Fn3772
func Fn3772(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3776 github.com/goccy/perlwasm2go/p0.Fn3776
func Fn3776(m *base.Module, l0 int64)

//go:linkname Fn3777 github.com/goccy/perlwasm2go/p0.Fn3777
func Fn3777(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3780 github.com/goccy/perlwasm2go/p0.Fn3780
func Fn3780(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3782 github.com/goccy/perlwasm2go/p0.Fn3782
func Fn3782(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn3785 github.com/goccy/perlwasm2go/p0.Fn3785
func Fn3785(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3786 github.com/goccy/perlwasm2go/p0.Fn3786
func Fn3786(m *base.Module, l0 int64, l1 int32) float64

//go:linkname Fn3788 github.com/goccy/perlwasm2go/p0.Fn3788
func Fn3788(m *base.Module, l0 int64) int64

//go:linkname Fn3790 github.com/goccy/perlwasm2go/p0.Fn3790
func Fn3790(m *base.Module, l0 int64) int64

//go:linkname Fn3792 github.com/goccy/perlwasm2go/p0.Fn3792
func Fn3792(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3793 github.com/goccy/perlwasm2go/p0.Fn3793
func Fn3793(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3794 github.com/goccy/perlwasm2go/p0.Fn3794
func Fn3794(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3795 github.com/goccy/perlwasm2go/p2.Fn3795
func Fn3795(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3796 github.com/goccy/perlwasm2go/p2.Fn3796
func Fn3796(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3797 github.com/goccy/perlwasm2go/p2.Fn3797
func Fn3797(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3798 github.com/goccy/perlwasm2go/p0.Fn3798
func Fn3798(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3799 github.com/goccy/perlwasm2go/p0.Fn3799
func Fn3799(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3800 github.com/goccy/perlwasm2go/p0.Fn3800
func Fn3800(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3801 github.com/goccy/perlwasm2go/p0.Fn3801
func Fn3801(m *base.Module, l0 int64) int32

//go:linkname Fn3808 github.com/goccy/perlwasm2go/p2.Fn3808
func Fn3808(m *base.Module, l0 int64)

//go:linkname Fn3809 github.com/goccy/perlwasm2go/p2.Fn3809
func Fn3809(m *base.Module, l0 int64) int32

//go:linkname Fn3811 github.com/goccy/perlwasm2go/p0.Fn3811
func Fn3811(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3818 github.com/goccy/perlwasm2go/p2.Fn3818
func Fn3818(m *base.Module, l0 int64)

//go:linkname Fn3819 github.com/goccy/perlwasm2go/p2.Fn3819
func Fn3819(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3820 github.com/goccy/perlwasm2go/p2.Fn3820
func Fn3820(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3821 github.com/goccy/perlwasm2go/p0.Fn3821
func Fn3821(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3822 github.com/goccy/perlwasm2go/p0.Fn3822
func Fn3822(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3823 github.com/goccy/perlwasm2go/p0.Fn3823
func Fn3823(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3824 github.com/goccy/perlwasm2go/p2.Fn3824
func Fn3824(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3825 github.com/goccy/perlwasm2go/p0.Fn3825
func Fn3825(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3827 github.com/goccy/perlwasm2go/p0.Fn3827
func Fn3827(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3829 github.com/goccy/perlwasm2go/p0.Fn3829
func Fn3829(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3830 github.com/goccy/perlwasm2go/p0.Fn3830
func Fn3830(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3831 github.com/goccy/perlwasm2go/p0.Fn3831
func Fn3831(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3832 github.com/goccy/perlwasm2go/p0.Fn3832
func Fn3832(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3833 github.com/goccy/perlwasm2go/p0.Fn3833
func Fn3833(m *base.Module, l0 int64) int64

//go:linkname Fn3834 github.com/goccy/perlwasm2go/p0.Fn3834
func Fn3834(m *base.Module, l0 int64)

//go:linkname Fn3835 github.com/goccy/perlwasm2go/p0.Fn3835
func Fn3835(m *base.Module, l0 int64) int64

//go:linkname Fn3836 github.com/goccy/perlwasm2go/p2.Fn3836
func Fn3836(m *base.Module, l0 int64) int64

//go:linkname Fn3840 github.com/goccy/perlwasm2go/p0.Fn3840
func Fn3840(m *base.Module, l0 int64) int64

//go:linkname Fn3841 github.com/goccy/perlwasm2go/p0.Fn3841
func Fn3841(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3842 github.com/goccy/perlwasm2go/p0.Fn3842
func Fn3842(m *base.Module, l0 int64)

//go:linkname Fn3843 github.com/goccy/perlwasm2go/p0.Fn3843
func Fn3843(m *base.Module, l0 int64) int64

//go:linkname Fn3844 github.com/goccy/perlwasm2go/p2.Fn3844
func Fn3844(m *base.Module, l0 int64) int64

//go:linkname Fn3845 github.com/goccy/perlwasm2go/p0.Fn3845
func Fn3845(m *base.Module, l0 int64) int64

//go:linkname Fn3846 github.com/goccy/perlwasm2go/p0.Fn3846
func Fn3846(m *base.Module, l0 int64) int64

//go:linkname Fn3847 github.com/goccy/perlwasm2go/p0.Fn3847
func Fn3847(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3851 github.com/goccy/perlwasm2go/p0.Fn3851
func Fn3851(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3852 github.com/goccy/perlwasm2go/p2.Fn3852
func Fn3852(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3858 github.com/goccy/perlwasm2go/p0.Fn3858
func Fn3858(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3863 github.com/goccy/perlwasm2go/p2.Fn3863
func Fn3863(m *base.Module, l0 int64)

//go:linkname Fn3865 github.com/goccy/perlwasm2go/p2.Fn3865
func Fn3865(m *base.Module, l0 int64)

//go:linkname Fn3867 github.com/goccy/perlwasm2go/p0.Fn3867
func Fn3867(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3868 github.com/goccy/perlwasm2go/p0.Fn3868
func Fn3868(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3869 github.com/goccy/perlwasm2go/p0.Fn3869
func Fn3869(m *base.Module, l0 int64) int64

//go:linkname Fn3870 github.com/goccy/perlwasm2go/p0.Fn3870
func Fn3870(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3871 github.com/goccy/perlwasm2go/p0.Fn3871
func Fn3871(m *base.Module, l0 int64) int64

//go:linkname Fn3877 github.com/goccy/perlwasm2go/p0.Fn3877
func Fn3877(m *base.Module, l0 int64) int32

//go:linkname Fn3878 github.com/goccy/perlwasm2go/p0.Fn3878
func Fn3878(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3879 github.com/goccy/perlwasm2go/p0.Fn3879
func Fn3879(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3881 github.com/goccy/perlwasm2go/p2.Fn3881
func Fn3881(m *base.Module, l0 float64) int64

//go:linkname Fn3882 github.com/goccy/perlwasm2go/p2.Fn3882
func Fn3882(m *base.Module, l0 int64) int64

//go:linkname Fn3883 github.com/goccy/perlwasm2go/p0.Fn3883
func Fn3883(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3884 github.com/goccy/perlwasm2go/p2.Fn3884
func Fn3884(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3885 github.com/goccy/perlwasm2go/p2.Fn3885
func Fn3885(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3886 github.com/goccy/perlwasm2go/p2.Fn3886
func Fn3886(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3887 github.com/goccy/perlwasm2go/p0.Fn3887
func Fn3887(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3888 github.com/goccy/perlwasm2go/p0.Fn3888
func Fn3888(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3889 github.com/goccy/perlwasm2go/p2.Fn3889
func Fn3889(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3892 github.com/goccy/perlwasm2go/p0.Fn3892
func Fn3892(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64)

//go:linkname Fn3895 github.com/goccy/perlwasm2go/p2.Fn3895
func Fn3895(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3906 github.com/goccy/perlwasm2go/p0.Fn3906
func Fn3906(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3907 github.com/goccy/perlwasm2go/p2.Fn3907
func Fn3907(m *base.Module)

//go:linkname Fn3908 github.com/goccy/perlwasm2go/p2.Fn3908
func Fn3908(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3909 github.com/goccy/perlwasm2go/p0.Fn3909
func Fn3909(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3910 github.com/goccy/perlwasm2go/p0.Fn3910
func Fn3910(m *base.Module, l0 int32) int64

//go:linkname Fn3911 github.com/goccy/perlwasm2go/p0.Fn3911
func Fn3911(m *base.Module, l0 int64)

//go:linkname Fn3912 github.com/goccy/perlwasm2go/p0.Fn3912
func Fn3912(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3913 github.com/goccy/perlwasm2go/p0.Fn3913
func Fn3913(m *base.Module, l0 int64)

//go:linkname Fn3915 github.com/goccy/perlwasm2go/p2.Fn3915
func Fn3915(m *base.Module, l0 int64) int64

//go:linkname Fn3923 github.com/goccy/perlwasm2go/p2.Fn3923
func Fn3923(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn3924 github.com/goccy/perlwasm2go/p2.Fn3924
func Fn3924(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3926 github.com/goccy/perlwasm2go/p2.Fn3926
func Fn3926(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3930 github.com/goccy/perlwasm2go/p0.Fn3930
func Fn3930(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3931 github.com/goccy/perlwasm2go/p2.Fn3931
func Fn3931(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3933 github.com/goccy/perlwasm2go/p2.Fn3933
func Fn3933(m *base.Module, l0 int32)

//go:linkname Fn3934 github.com/goccy/perlwasm2go/p2.Fn3934
func Fn3934(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3936 github.com/goccy/perlwasm2go/p0.Fn3936
func Fn3936(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3938 github.com/goccy/perlwasm2go/p2.Fn3938
func Fn3938(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3939 github.com/goccy/perlwasm2go/p2.Fn3939
func Fn3939(m *base.Module, l0 int64)

//go:linkname Fn3940 github.com/goccy/perlwasm2go/p2.Fn3940
func Fn3940(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3942 github.com/goccy/perlwasm2go/p2.Fn3942
func Fn3942(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3944 github.com/goccy/perlwasm2go/p2.Fn3944
func Fn3944(m *base.Module, l0 int64) int64

//go:linkname Fn3945 github.com/goccy/perlwasm2go/p2.Fn3945
func Fn3945(m *base.Module, l0 int64) int64

//go:linkname Fn3946 github.com/goccy/perlwasm2go/p0.Fn3946
func Fn3946(m *base.Module, l0 int64)

//go:linkname Fn3947 github.com/goccy/perlwasm2go/p2.Fn3947
func Fn3947(m *base.Module)

//go:linkname Fn3952 github.com/goccy/perlwasm2go/p2.Fn3952
func Fn3952(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn3953 github.com/goccy/perlwasm2go/p2.Fn3953
func Fn3953(m *base.Module)

//go:linkname Fn3958 github.com/goccy/perlwasm2go/p2.Fn3958
func Fn3958(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int64

//go:linkname Fn3962 github.com/goccy/perlwasm2go/p2.Fn3962
func Fn3962(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3963 github.com/goccy/perlwasm2go/p2.Fn3963
func Fn3963(m *base.Module) int32

//go:linkname Fn3965 github.com/goccy/perlwasm2go/p2.Fn3965
func Fn3965(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3966 github.com/goccy/perlwasm2go/p2.Fn3966
func Fn3966(m *base.Module, l0 int64) int64

//go:linkname Fn3967 github.com/goccy/perlwasm2go/p2.Fn3967
func Fn3967(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32) int64

//go:linkname Fn3968 github.com/goccy/perlwasm2go/p2.Fn3968
func Fn3968(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3971 github.com/goccy/perlwasm2go/p2.Fn3971
func Fn3971(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32

//go:linkname Fn3972 github.com/goccy/perlwasm2go/p2.Fn3972
func Fn3972(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int64

//go:linkname Fn3973 github.com/goccy/perlwasm2go/p0.Fn3973
func Fn3973(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3974 github.com/goccy/perlwasm2go/p0.Fn3974
func Fn3974(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3975 github.com/goccy/perlwasm2go/p2.Fn3975
func Fn3975(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3976 github.com/goccy/perlwasm2go/p2.Fn3976
func Fn3976(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3977 github.com/goccy/perlwasm2go/p0.Fn3977
func Fn3977(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3992 github.com/goccy/perlwasm2go/p2.Fn3992
func Fn3992(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3993 github.com/goccy/perlwasm2go/p2.Fn3993
func Fn3993(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3994 github.com/goccy/perlwasm2go/p2.Fn3994
func Fn3994(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3995 github.com/goccy/perlwasm2go/p2.Fn3995
func Fn3995(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3997 github.com/goccy/perlwasm2go/p2.Fn3997
func Fn3997(m *base.Module, l0 int32)

//go:linkname Fn3998 github.com/goccy/perlwasm2go/p2.Fn3998
func Fn3998(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn4000 github.com/goccy/perlwasm2go/p2.Fn4000
func Fn4000(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32

//go:linkname Fn4001 github.com/goccy/perlwasm2go/p2.Fn4001
func Fn4001(m *base.Module, l0 int64) int64

//go:linkname Fn4003 github.com/goccy/perlwasm2go/p2.Fn4003
func Fn4003(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn4006 github.com/goccy/perlwasm2go/p2.Fn4006
func Fn4006(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4008 github.com/goccy/perlwasm2go/p0.Fn4008
func Fn4008(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4009 github.com/goccy/perlwasm2go/p0.Fn4009
func Fn4009(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4011 github.com/goccy/perlwasm2go/p0.Fn4011
func Fn4011(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn4012 github.com/goccy/perlwasm2go/p0.Fn4012
func Fn4012(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int64) int32

//go:linkname Fn4013 github.com/goccy/perlwasm2go/p0.Fn4013
func Fn4013(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn4017 github.com/goccy/perlwasm2go/p2.Fn4017
func Fn4017(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4018 github.com/goccy/perlwasm2go/p2.Fn4018
func Fn4018(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4019 github.com/goccy/perlwasm2go/p2.Fn4019
func Fn4019(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4021 github.com/goccy/perlwasm2go/p0.Fn4021
func Fn4021(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4023 github.com/goccy/perlwasm2go/p0.Fn4023
func Fn4023(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4024 github.com/goccy/perlwasm2go/p0.Fn4024
func Fn4024(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn4025 github.com/goccy/perlwasm2go/p0.Fn4025
func Fn4025(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn4026 github.com/goccy/perlwasm2go/p2.Fn4026
func Fn4026(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4027 github.com/goccy/perlwasm2go/p0.Fn4027
func Fn4027(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4032 github.com/goccy/perlwasm2go/p2.Fn4032
func Fn4032(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn4034 github.com/goccy/perlwasm2go/p2.Fn4034
func Fn4034(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64) int64

//go:linkname Fn4039 github.com/goccy/perlwasm2go/p2.Fn4039
func Fn4039(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn4040 github.com/goccy/perlwasm2go/p2.Fn4040
func Fn4040(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn4042 github.com/goccy/perlwasm2go/p2.Fn4042
func Fn4042(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn4043 github.com/goccy/perlwasm2go/p2.Fn4043
func Fn4043(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4045 github.com/goccy/perlwasm2go/p0.Fn4045
func Fn4045(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4046 github.com/goccy/perlwasm2go/p2.Fn4046
func Fn4046(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4049 github.com/goccy/perlwasm2go/p2.Fn4049
func Fn4049(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn4051 github.com/goccy/perlwasm2go/p0.Fn4051
func Fn4051(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn4052 github.com/goccy/perlwasm2go/p2.Fn4052
func Fn4052(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int32) int32

//go:linkname Fn4053 github.com/goccy/perlwasm2go/p0.Fn4053
func Fn4053(m *base.Module, l0 int64) int64

//go:linkname Fn4055 github.com/goccy/perlwasm2go/p0.Fn4055
func Fn4055(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4056 github.com/goccy/perlwasm2go/p0.Fn4056
func Fn4056(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4057 github.com/goccy/perlwasm2go/p0.Fn4057
func Fn4057(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4065 github.com/goccy/perlwasm2go/p0.Fn4065
func Fn4065(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4070 github.com/goccy/perlwasm2go/p2.Fn4070
func Fn4070(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn4072 github.com/goccy/perlwasm2go/p0.Fn4072
func Fn4072(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn4073 github.com/goccy/perlwasm2go/p0.Fn4073
func Fn4073(m *base.Module)

//go:linkname Fn4074 github.com/goccy/perlwasm2go/p2.Fn4074
func Fn4074(m *base.Module, l0 int64)

//go:linkname Fn4075 github.com/goccy/perlwasm2go/p0.Fn4075
func Fn4075(m *base.Module, l0 int64)

//go:linkname Fn4078 github.com/goccy/perlwasm2go/p0.Fn4078
func Fn4078(m *base.Module)

//go:linkname Fn4081 github.com/goccy/perlwasm2go/p0.Fn4081
func Fn4081(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4082 github.com/goccy/perlwasm2go/p0.Fn4082
func Fn4082(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn4085 github.com/goccy/perlwasm2go/p2.Fn4085
func Fn4085(m *base.Module, l0 int32) int32

//go:linkname Fn4087 github.com/goccy/perlwasm2go/p0.Fn4087
func Fn4087(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn4089 github.com/goccy/perlwasm2go/p0.Fn4089
func Fn4089(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn4090 github.com/goccy/perlwasm2go/p2.Fn4090
func Fn4090(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4091 github.com/goccy/perlwasm2go/p2.Fn4091
func Fn4091(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4092 github.com/goccy/perlwasm2go/p2.Fn4092
func Fn4092(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn4093 github.com/goccy/perlwasm2go/p2.Fn4093
func Fn4093(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4098 github.com/goccy/perlwasm2go/p2.Fn4098
func Fn4098(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn4099 github.com/goccy/perlwasm2go/p2.Fn4099
func Fn4099(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn4100 github.com/goccy/perlwasm2go/p2.Fn4100
func Fn4100(m *base.Module, l0 int64)

//go:linkname Fn4103 github.com/goccy/perlwasm2go/p0.Fn4103
func Fn4103(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4106 github.com/goccy/perlwasm2go/p0.Fn4106
func Fn4106(m *base.Module, l0 int64) int64

//go:linkname Fn4107 github.com/goccy/perlwasm2go/p0.Fn4107
func Fn4107(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4108 github.com/goccy/perlwasm2go/p0.Fn4108
func Fn4108(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn4111 github.com/goccy/perlwasm2go/p2.Fn4111
func Fn4111(m *base.Module, l0 int64) int64

//go:linkname Fn4113 github.com/goccy/perlwasm2go/p0.Fn4113
func Fn4113(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4116 github.com/goccy/perlwasm2go/p2.Fn4116
func Fn4116(m *base.Module, l0 int64) int32

//go:linkname Fn4117 github.com/goccy/perlwasm2go/p2.Fn4117
func Fn4117(m *base.Module) int64

//go:linkname Fn4119 github.com/goccy/perlwasm2go/p2.Fn4119
func Fn4119(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn4120 github.com/goccy/perlwasm2go/p2.Fn4120
func Fn4120(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4123 github.com/goccy/perlwasm2go/p0.Fn4123
func Fn4123(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn4124 github.com/goccy/perlwasm2go/p2.Fn4124
func Fn4124(m *base.Module, l0 int64) int64

//go:linkname Fn4137 github.com/goccy/perlwasm2go/p2.Fn4137
func Fn4137(m *base.Module) int64

//go:linkname Fn4139 github.com/goccy/perlwasm2go/p2.Fn4139
func Fn4139(m *base.Module, l0 int32) int64

//go:linkname Fn4141 github.com/goccy/perlwasm2go/p2.Fn4141
func Fn4141(m *base.Module, l0 int32) int64

//go:linkname Fn4147 github.com/goccy/perlwasm2go/p2.Fn4147
func Fn4147(m *base.Module) int32

//go:linkname Fn4148 github.com/goccy/perlwasm2go/p2.Fn4148
func Fn4148(m *base.Module, l0 int64) int64

//go:linkname Fn4149 github.com/goccy/perlwasm2go/p2.Fn4149
func Fn4149(m *base.Module, l0 int32) int64

//go:linkname Fn4150 github.com/goccy/perlwasm2go/p2.Fn4150
func Fn4150(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn4153 github.com/goccy/perlwasm2go/p2.Fn4153
func Fn4153(m *base.Module, l0 int64) int64

//go:linkname Fn4154 github.com/goccy/perlwasm2go/p2.Fn4154
func Fn4154(m *base.Module, l0 int32) int64

//go:linkname Fn4156 github.com/goccy/perlwasm2go/p2.Fn4156
func Fn4156(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int32) int64

//go:linkname Fn4157 github.com/goccy/perlwasm2go/p2.Fn4157
func Fn4157(m *base.Module, l0 int64) int64

//go:linkname Fn4160 github.com/goccy/perlwasm2go/p2.Fn4160
func Fn4160(m *base.Module, l0 int64)

//go:linkname Fn4161 github.com/goccy/perlwasm2go/p2.Fn4161
func Fn4161(m *base.Module, l0 int64)

//go:linkname Fn4167 github.com/goccy/perlwasm2go/p2.Fn4167
func Fn4167(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn4171 github.com/goccy/perlwasm2go/p2.Fn4171
func Fn4171(m *base.Module, l0 int64)

//go:linkname Fn4172 github.com/goccy/perlwasm2go/p2.Fn4172
func Fn4172(m *base.Module, l0 int64) int64

//go:linkname Fn4176 github.com/goccy/perlwasm2go/p2.Fn4176
func Fn4176(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn4183 github.com/goccy/perlwasm2go/p2.Fn4183
func Fn4183(m *base.Module, l0 int32) int64

//go:linkname Fn4185 github.com/goccy/perlwasm2go/p2.Fn4185
func Fn4185(m *base.Module, l0 int64) int64

//go:linkname Fn4186 github.com/goccy/perlwasm2go/p2.Fn4186
func Fn4186(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4199 github.com/goccy/perlwasm2go/p0.Fn4199
func Fn4199(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn4206 github.com/goccy/perlwasm2go/p2.Fn4206
func Fn4206(m *base.Module, l0 int64)

//go:linkname Fn4210 github.com/goccy/perlwasm2go/p2.Fn4210
func Fn4210(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64) int64

//go:linkname Fn4215 github.com/goccy/perlwasm2go/p2.Fn4215
func Fn4215(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4216 github.com/goccy/perlwasm2go/p2.Fn4216
func Fn4216(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4217 github.com/goccy/perlwasm2go/p2.Fn4217
func Fn4217(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4221 github.com/goccy/perlwasm2go/p2.Fn4221
func Fn4221(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4223 github.com/goccy/perlwasm2go/p2.Fn4223
func Fn4223(m *base.Module, l0 int64) int64

//go:linkname Fn4250 github.com/goccy/perlwasm2go/p2.Fn4250
func Fn4250(m *base.Module, l0 int32) int64

//go:linkname Fn4251 github.com/goccy/perlwasm2go/p2.Fn4251
func Fn4251(m *base.Module, l0 int64) int64

//go:linkname Fn4252 github.com/goccy/perlwasm2go/p2.Fn4252
func Fn4252(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4253 github.com/goccy/perlwasm2go/p2.Fn4253
func Fn4253(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4254 github.com/goccy/perlwasm2go/p2.Fn4254
func Fn4254(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4255 github.com/goccy/perlwasm2go/p2.Fn4255
func Fn4255(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4256 github.com/goccy/perlwasm2go/p2.Fn4256
func Fn4256(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4257 github.com/goccy/perlwasm2go/p2.Fn4257
func Fn4257(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn4259 github.com/goccy/perlwasm2go/p2.Fn4259
func Fn4259(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64, l6 int64)

//go:linkname Fn4263 github.com/goccy/perlwasm2go/p0.Fn4263
func Fn4263(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn4269 github.com/goccy/perlwasm2go/p2.Fn4269
func Fn4269(m *base.Module, l0 int64) int64

//go:linkname Fn4270 github.com/goccy/perlwasm2go/p2.Fn4270
func Fn4270(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn4277 github.com/goccy/perlwasm2go/p2.Fn4277
func Fn4277(m *base.Module, l0 int64)

//go:linkname Fn4278 github.com/goccy/perlwasm2go/p2.Fn4278
func Fn4278(m *base.Module, l0 int64) int32

//go:linkname Fn4279 github.com/goccy/perlwasm2go/p2.Fn4279
func Fn4279(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4280 github.com/goccy/perlwasm2go/p2.Fn4280
func Fn4280(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4282 github.com/goccy/perlwasm2go/p2.Fn4282
func Fn4282(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn4285 github.com/goccy/perlwasm2go/p2.Fn4285
func Fn4285(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn4287 github.com/goccy/perlwasm2go/p2.Fn4287
func Fn4287(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn4288 github.com/goccy/perlwasm2go/p2.Fn4288
func Fn4288(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn4289 github.com/goccy/perlwasm2go/p2.Fn4289
func Fn4289(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4292 github.com/goccy/perlwasm2go/p2.Fn4292
func Fn4292(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4294 github.com/goccy/perlwasm2go/p2.Fn4294
func Fn4294(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int32) int32

//go:linkname Fn4295 github.com/goccy/perlwasm2go/p2.Fn4295
func Fn4295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32) int32

//go:linkname Fn4297 github.com/goccy/perlwasm2go/p2.Fn4297
func Fn4297(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn4298 github.com/goccy/perlwasm2go/p2.Fn4298
func Fn4298(m *base.Module, l0 int64) int32

//go:linkname Fn4299 github.com/goccy/perlwasm2go/p2.Fn4299
func Fn4299(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4300 github.com/goccy/perlwasm2go/p2.Fn4300
func Fn4300(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn4301 github.com/goccy/perlwasm2go/p2.Fn4301
func Fn4301(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn4303 github.com/goccy/perlwasm2go/p2.Fn4303
func Fn4303(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64)

//go:linkname Fn4304 github.com/goccy/perlwasm2go/p2.Fn4304
func Fn4304(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4307 github.com/goccy/perlwasm2go/p2.Fn4307
func Fn4307(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn4308 github.com/goccy/perlwasm2go/p2.Fn4308
func Fn4308(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32) int64

//go:linkname Fn4310 github.com/goccy/perlwasm2go/p2.Fn4310
func Fn4310(m *base.Module, l0 int32) int64

//go:linkname Fn4311 github.com/goccy/perlwasm2go/p2.Fn4311
func Fn4311(m *base.Module, l0 int64) int64

//go:linkname Fn4316 github.com/goccy/perlwasm2go/p2.Fn4316
func Fn4316(m *base.Module, l0 int64) int32

//go:linkname Fn4317 github.com/goccy/perlwasm2go/p2.Fn4317
func Fn4317(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32

//go:linkname Fn4318 github.com/goccy/perlwasm2go/p2.Fn4318
func Fn4318(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn4323 github.com/goccy/perlwasm2go/p2.Fn4323
func Fn4323(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4325 github.com/goccy/perlwasm2go/p2.Fn4325
func Fn4325(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4327 github.com/goccy/perlwasm2go/p2.Fn4327
func Fn4327(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4328 github.com/goccy/perlwasm2go/p2.Fn4328
func Fn4328(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int32)

//go:linkname Fn4329 github.com/goccy/perlwasm2go/p2.Fn4329
func Fn4329(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4331 github.com/goccy/perlwasm2go/p2.Fn4331
func Fn4331(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4332 github.com/goccy/perlwasm2go/p2.Fn4332
func Fn4332(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn4333 github.com/goccy/perlwasm2go/p2.Fn4333
func Fn4333(m *base.Module, l0 int64) int64

//go:linkname Fn4334 github.com/goccy/perlwasm2go/p2.Fn4334
func Fn4334(m *base.Module, l0 int64)

//go:linkname Fn4335 github.com/goccy/perlwasm2go/p2.Fn4335
func Fn4335(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4367 github.com/goccy/perlwasm2go/p2.Fn4367
func Fn4367(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn4375 github.com/goccy/perlwasm2go/p0.Fn4375
func Fn4375(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4384 github.com/goccy/perlwasm2go/p2.Fn4384
func Fn4384(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4389 github.com/goccy/perlwasm2go/p2.Fn4389
func Fn4389(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4394 github.com/goccy/perlwasm2go/p2.Fn4394
func Fn4394(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4395 github.com/goccy/perlwasm2go/p2.Fn4395
func Fn4395(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4396 github.com/goccy/perlwasm2go/p2.Fn4396
func Fn4396(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn4398 github.com/goccy/perlwasm2go/p2.Fn4398
func Fn4398(m *base.Module, l0 int64)

//go:linkname Fn4401 github.com/goccy/perlwasm2go/p2.Fn4401
func Fn4401(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4402 github.com/goccy/perlwasm2go/p2.Fn4402
func Fn4402(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4404 github.com/goccy/perlwasm2go/p2.Fn4404
func Fn4404(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4406 github.com/goccy/perlwasm2go/p2.Fn4406
func Fn4406(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4408 github.com/goccy/perlwasm2go/p2.Fn4408
func Fn4408(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4409 github.com/goccy/perlwasm2go/p2.Fn4409
func Fn4409(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4411 github.com/goccy/perlwasm2go/p2.Fn4411
func Fn4411(m *base.Module, l0 int64) int64

//go:linkname Fn4414 github.com/goccy/perlwasm2go/p2.Fn4414
func Fn4414(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4417 github.com/goccy/perlwasm2go/p2.Fn4417
func Fn4417(m *base.Module, l0 int64) int64

//go:linkname Fn4425 github.com/goccy/perlwasm2go/p2.Fn4425
func Fn4425(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4427 github.com/goccy/perlwasm2go/p2.Fn4427
func Fn4427(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn4435 github.com/goccy/perlwasm2go/p2.Fn4435
func Fn4435(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn4438 github.com/goccy/perlwasm2go/p2.Fn4438
func Fn4438(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4439 github.com/goccy/perlwasm2go/p2.Fn4439
func Fn4439(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4440 github.com/goccy/perlwasm2go/p2.Fn4440
func Fn4440(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4445 github.com/goccy/perlwasm2go/p2.Fn4445
func Fn4445(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4451 github.com/goccy/perlwasm2go/p2.Fn4451
func Fn4451(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4458 github.com/goccy/perlwasm2go/p2.Fn4458
func Fn4458(m *base.Module, l0 int64)

//go:linkname Fn4465 github.com/goccy/perlwasm2go/p2.Fn4465
func Fn4465(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4484 github.com/goccy/perlwasm2go/p2.Fn4484
func Fn4484(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4558 github.com/goccy/perlwasm2go/p0.Fn4558
func Fn4558(m *base.Module)

//go:linkname Fn4560 github.com/goccy/perlwasm2go/p0.Fn4560
func Fn4560(m *base.Module, l0 int64) int64

//go:linkname Fn4568 github.com/goccy/perlwasm2go/p2.Fn4568
func Fn4568(m *base.Module)

//go:linkname Fn4570 github.com/goccy/perlwasm2go/p0.Fn4570
func Fn4570(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn4572 github.com/goccy/perlwasm2go/p2.Fn4572
func Fn4572(m *base.Module, l0 int64) int64

//go:linkname Fn4574 github.com/goccy/perlwasm2go/p0.Fn4574
func Fn4574(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4575 github.com/goccy/perlwasm2go/p0.Fn4575
func Fn4575(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4576 github.com/goccy/perlwasm2go/p0.Fn4576
func Fn4576(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn4577 github.com/goccy/perlwasm2go/p2.Fn4577
func Fn4577(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn4578 github.com/goccy/perlwasm2go/p2.Fn4578
func Fn4578(m *base.Module, l0 int64) int32

//go:linkname Fn4581 github.com/goccy/perlwasm2go/p2.Fn4581
func Fn4581(m *base.Module, l0 int32)

//go:linkname Fn4584 github.com/goccy/perlwasm2go/p2.Fn4584
func Fn4584(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn4586 github.com/goccy/perlwasm2go/p2.Fn4586
func Fn4586(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn4590 github.com/goccy/perlwasm2go/p2.Fn4590
func Fn4590(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn4591 github.com/goccy/perlwasm2go/p2.Fn4591
func Fn4591(m *base.Module, l0 int32, l1 int64, l2 int64) int64

//go:linkname Fn4594 github.com/goccy/perlwasm2go/p2.Fn4594
func Fn4594(m *base.Module, l0 int32) int32

//go:linkname Fn4595 github.com/goccy/perlwasm2go/p2.Fn4595
func Fn4595(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn4596 github.com/goccy/perlwasm2go/p2.Fn4596
func Fn4596(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn4597 github.com/goccy/perlwasm2go/p2.Fn4597
func Fn4597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4598 github.com/goccy/perlwasm2go/p2.Fn4598
func Fn4598(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn4602 github.com/goccy/perlwasm2go/p2.Fn4602
func Fn4602(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn4605 github.com/goccy/perlwasm2go/p2.Fn4605
func Fn4605(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn4607 github.com/goccy/perlwasm2go/p2.Fn4607
func Fn4607(m *base.Module, l0 int64) int32

//go:linkname Fn4610 github.com/goccy/perlwasm2go/p2.Fn4610
func Fn4610(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4613 github.com/goccy/perlwasm2go/p2.Fn4613
func Fn4613(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4614 github.com/goccy/perlwasm2go/p2.Fn4614
func Fn4614(m *base.Module, l0 int64) int32

//go:linkname Fn4615 github.com/goccy/perlwasm2go/p2.Fn4615
func Fn4615(m *base.Module, l0 int64) int64

//go:linkname Fn4619 github.com/goccy/perlwasm2go/p2.Fn4619
func Fn4619(m *base.Module, l0 int64) int64

//go:linkname Fn4620 github.com/goccy/perlwasm2go/p2.Fn4620
func Fn4620(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn4624 github.com/goccy/perlwasm2go/p2.Fn4624
func Fn4624(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn4625 github.com/goccy/perlwasm2go/p2.Fn4625
func Fn4625(m *base.Module, l0 int64)

//go:linkname Fn4628 github.com/goccy/perlwasm2go/p2.Fn4628
func Fn4628(m *base.Module, l0 int64) int32

//go:linkname Fn4630 github.com/goccy/perlwasm2go/p2.Fn4630
func Fn4630(m *base.Module, l0 int32) int32

//go:linkname Fn4633 github.com/goccy/perlwasm2go/p2.Fn4633
func Fn4633(m *base.Module)

//go:linkname Fn4640 github.com/goccy/perlwasm2go/p2.Fn4640
func Fn4640(m *base.Module, l0 int64) int64

//go:linkname Fn4642 github.com/goccy/perlwasm2go/p2.Fn4642
func Fn4642(m *base.Module, l0 int64) int64

//go:linkname Fn4645 github.com/goccy/perlwasm2go/p2.Fn4645
func Fn4645(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn4659 github.com/goccy/perlwasm2go/p2.Fn4659
func Fn4659(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4662 github.com/goccy/perlwasm2go/p2.Fn4662
func Fn4662(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4664 github.com/goccy/perlwasm2go/p2.Fn4664
func Fn4664(m *base.Module, l0 int32) int32

//go:linkname Fn4668 github.com/goccy/perlwasm2go/p2.Fn4668
func Fn4668(m *base.Module, l0 int32) int32

//go:linkname Fn4670 github.com/goccy/perlwasm2go/p2.Fn4670
func Fn4670(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4671 github.com/goccy/perlwasm2go/p2.Fn4671
func Fn4671(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn4673 github.com/goccy/perlwasm2go/p2.Fn4673
func Fn4673(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn4676 github.com/goccy/perlwasm2go/p2.Fn4676
func Fn4676(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn4678 github.com/goccy/perlwasm2go/p2.Fn4678
func Fn4678(m *base.Module, l0 float64, l1 float64, l2 int32) float64

//go:linkname Fn4681 github.com/goccy/perlwasm2go/p2.Fn4681
func Fn4681(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn4682 github.com/goccy/perlwasm2go/p2.Fn4682
func Fn4682(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn4683 github.com/goccy/perlwasm2go/p2.Fn4683
func Fn4683(m *base.Module, l0 int32) float64

//go:linkname Fn4684 github.com/goccy/perlwasm2go/p2.Fn4684
func Fn4684(m *base.Module, l0 int32) float64

//go:linkname Fn4685 github.com/goccy/perlwasm2go/p2.Fn4685
func Fn4685(m *base.Module, l0 float64) float64

//go:linkname Fn4688 github.com/goccy/perlwasm2go/p2.Fn4688
func Fn4688(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn4693 github.com/goccy/perlwasm2go/p2.Fn4693
func Fn4693(m *base.Module, l0 int32, l1 float64, l2 int32) float64

//go:linkname Fn4694 github.com/goccy/perlwasm2go/p2.Fn4694
func Fn4694(m *base.Module, l0 float64) float64

//go:linkname Fn4696 github.com/goccy/perlwasm2go/p2.Fn4696
func Fn4696(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) float64

//go:linkname Fn4698 github.com/goccy/perlwasm2go/p2.Fn4698
func Fn4698(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn4700 github.com/goccy/perlwasm2go/p2.Fn4700
func Fn4700(m *base.Module, l0 float64) float64

//go:linkname Fn4707 github.com/goccy/perlwasm2go/p2.Fn4707
func Fn4707(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn4723 github.com/goccy/perlwasm2go/p2.Fn4723
func Fn4723(m *base.Module, l0 int64) int32

//go:linkname Fn4731 github.com/goccy/perlwasm2go/p2.Fn4731
func Fn4731(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4736 github.com/goccy/perlwasm2go/p2.Fn4736
func Fn4736(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn4737 github.com/goccy/perlwasm2go/p2.Fn4737
func Fn4737(m *base.Module)

//go:linkname Fn4738 github.com/goccy/perlwasm2go/p2.Fn4738
func Fn4738(m *base.Module, l0 int64)

//go:linkname Fn4740 github.com/goccy/perlwasm2go/p0.Fn4740
func Fn4740(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn4741 github.com/goccy/perlwasm2go/p0.Fn4741
func Fn4741(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn4749 github.com/goccy/perlwasm2go/p2.Fn4749
func Fn4749(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn4751 github.com/goccy/perlwasm2go/p2.Fn4751
func Fn4751(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4752 github.com/goccy/perlwasm2go/p2.Fn4752
func Fn4752(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn4753 github.com/goccy/perlwasm2go/p2.Fn4753
func Fn4753(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn4754 github.com/goccy/perlwasm2go/p2.Fn4754
func Fn4754(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn4756 github.com/goccy/perlwasm2go/p2.Fn4756
func Fn4756(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn4758 github.com/goccy/perlwasm2go/p2.Fn4758
func Fn4758(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4759 github.com/goccy/perlwasm2go/p2.Fn4759
func Fn4759(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4762 github.com/goccy/perlwasm2go/p2.Fn4762
func Fn4762(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4763 github.com/goccy/perlwasm2go/p2.Fn4763
func Fn4763(m *base.Module, l0 int64) int64

//go:linkname Fn4765 github.com/goccy/perlwasm2go/p2.Fn4765
func Fn4765(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn4766 github.com/goccy/perlwasm2go/p2.Fn4766
func Fn4766(m *base.Module, l0 int64) int64

//go:linkname Fn4769 github.com/goccy/perlwasm2go/p2.Fn4769
func Fn4769(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4772 github.com/goccy/perlwasm2go/p2.Fn4772
func Fn4772(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4775 github.com/goccy/perlwasm2go/p2.Fn4775
func Fn4775(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn4776 github.com/goccy/perlwasm2go/p2.Fn4776
func Fn4776(m *base.Module, l0 int64) int64

//go:linkname Fn4779 github.com/goccy/perlwasm2go/p2.Fn4779
func Fn4779(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn4795 github.com/goccy/perlwasm2go/p2.Fn4795
func Fn4795(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn4796 github.com/goccy/perlwasm2go/p2.Fn4796
func Fn4796(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn4829 github.com/goccy/perlwasm2go/p2.Fn4829
func Fn4829(m *base.Module, l0 int64)
