//go:build !arm64 && (!amd64 || !amd64.v2)

package p1

import (
	base "github.com/goccy/perlwasm2go/base"
	_ "unsafe"
)

//go:linkname Fn164 github.com/goccy/perlwasm2go/p0.Fn164
func Fn164(m *base.Module, l0 int32) int32

//go:linkname Fn169 github.com/goccy/perlwasm2go/p0.Fn169
func Fn169(m *base.Module, l0 int32) int32

//go:linkname Fn175 github.com/goccy/perlwasm2go/p0.Fn175
func Fn175(m *base.Module, l0 int32)

//go:linkname Fn176 github.com/goccy/perlwasm2go/p0.Fn176
func Fn176(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn177 github.com/goccy/perlwasm2go/p0.Fn177
func Fn177(m *base.Module, l0 int32) int32

//go:linkname Fn178 github.com/goccy/perlwasm2go/p0.Fn178
func Fn178(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn298 github.com/goccy/perlwasm2go/p0.Fn298
func Fn298(m *base.Module, l0 int32) int32

//go:linkname Fn300 github.com/goccy/perlwasm2go/p0.Fn300
func Fn300(m *base.Module, l0 int32) int32

//go:linkname Fn301 github.com/goccy/perlwasm2go/p0.Fn301
func Fn301(m *base.Module, l0 int32) int32

//go:linkname Fn338 github.com/goccy/perlwasm2go/p0.Fn338
func Fn338(m *base.Module, l0 int32)

//go:linkname Fn344 github.com/goccy/perlwasm2go/p0.Fn344
func Fn344(m *base.Module, l0 int32) int32

//go:linkname Fn372 github.com/goccy/perlwasm2go/p0.Fn372
func Fn372(m *base.Module, l0 int32) int32

//go:linkname Fn396 github.com/goccy/perlwasm2go/p0.Fn396
func Fn396(m *base.Module, l0 int32) int32

//go:linkname Fn399 github.com/goccy/perlwasm2go/p0.Fn399
func Fn399(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)

//go:linkname Fn427 github.com/goccy/perlwasm2go/p0.Fn427
func Fn427(m *base.Module, l0 int32) int32

//go:linkname Fn451 github.com/goccy/perlwasm2go/p0.Fn451
func Fn451(m *base.Module, l0 int32) float64

//go:linkname Fn454 github.com/goccy/perlwasm2go/p0.Fn454
func Fn454(m *base.Module, l0 int32) int32

//go:linkname Fn492 github.com/goccy/perlwasm2go/p0.Fn492
func Fn492(m *base.Module)

//go:linkname Fn501 github.com/goccy/perlwasm2go/p0.Fn501
func Fn501(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn508 github.com/goccy/perlwasm2go/p0.Fn508
func Fn508(m *base.Module, l0 int32)

//go:linkname Fn599 github.com/goccy/perlwasm2go/p0.Fn599
func Fn599(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn618 github.com/goccy/perlwasm2go/p0.Fn618
func Fn618(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn754 github.com/goccy/perlwasm2go/p0.Fn754
func Fn754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn829 github.com/goccy/perlwasm2go/p0.Fn829
func Fn829(m *base.Module, l0 int32) int32

//go:linkname Fn851 github.com/goccy/perlwasm2go/p0.Fn851
func Fn851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn948 github.com/goccy/perlwasm2go/p0.Fn948
func Fn948(m *base.Module)

//go:linkname Fn952 github.com/goccy/perlwasm2go/p0.Fn952
func Fn952(m *base.Module, l0 int32) int32

//go:linkname Fn971 github.com/goccy/perlwasm2go/p0.Fn971
func Fn971(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1054 github.com/goccy/perlwasm2go/p0.Fn1054
func Fn1054(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1063 github.com/goccy/perlwasm2go/p0.Fn1063
func Fn1063(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1097 github.com/goccy/perlwasm2go/p0.Fn1097
func Fn1097(m *base.Module, l0 int32) int32

//go:linkname Fn1100 github.com/goccy/perlwasm2go/p0.Fn1100
func Fn1100(m *base.Module, l0 int32)

//go:linkname Fn1102 github.com/goccy/perlwasm2go/p0.Fn1102
func Fn1102(m *base.Module, l0 int32)

//go:linkname Fn1111 github.com/goccy/perlwasm2go/p0.Fn1111
func Fn1111(m *base.Module, l0 int32)

//go:linkname Fn1112 github.com/goccy/perlwasm2go/p0.Fn1112
func Fn1112(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1113 github.com/goccy/perlwasm2go/p0.Fn1113
func Fn1113(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1114 github.com/goccy/perlwasm2go/p0.Fn1114
func Fn1114(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1115 github.com/goccy/perlwasm2go/p0.Fn1115
func Fn1115(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1117 github.com/goccy/perlwasm2go/p0.Fn1117
func Fn1117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1119 github.com/goccy/perlwasm2go/p0.Fn1119
func Fn1119(m *base.Module, l0 int32) int32

//go:linkname Fn1120 github.com/goccy/perlwasm2go/p0.Fn1120
func Fn1120(m *base.Module, l0 int32) int32

//go:linkname Fn1121 github.com/goccy/perlwasm2go/p0.Fn1121
func Fn1121(m *base.Module, l0 int32) int32

//go:linkname Fn1122 github.com/goccy/perlwasm2go/p0.Fn1122
func Fn1122(m *base.Module, l0 int32) int32

//go:linkname Fn1125 github.com/goccy/perlwasm2go/p0.Fn1125
func Fn1125(m *base.Module, l0 int32) int32

//go:linkname Fn1127 github.com/goccy/perlwasm2go/p0.Fn1127
func Fn1127(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1128 github.com/goccy/perlwasm2go/p0.Fn1128
func Fn1128(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1129 github.com/goccy/perlwasm2go/p0.Fn1129
func Fn1129(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1130 github.com/goccy/perlwasm2go/p0.Fn1130
func Fn1130(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1132 github.com/goccy/perlwasm2go/p0.Fn1132
func Fn1132(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1134 github.com/goccy/perlwasm2go/p0.Fn1134
func Fn1134(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1137 github.com/goccy/perlwasm2go/p0.Fn1137
func Fn1137(m *base.Module)

//go:linkname Fn1138 github.com/goccy/perlwasm2go/p0.Fn1138
func Fn1138(m *base.Module)

//go:linkname Fn1139 github.com/goccy/perlwasm2go/p0.Fn1139
func Fn1139(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1140 github.com/goccy/perlwasm2go/p0.Fn1140
func Fn1140(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1141 github.com/goccy/perlwasm2go/p0.Fn1141
func Fn1141(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1142 github.com/goccy/perlwasm2go/p0.Fn1142
func Fn1142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1147 github.com/goccy/perlwasm2go/p0.Fn1147
func Fn1147(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1148 github.com/goccy/perlwasm2go/p0.Fn1148
func Fn1148(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1149 github.com/goccy/perlwasm2go/p0.Fn1149
func Fn1149(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1150 github.com/goccy/perlwasm2go/p0.Fn1150
func Fn1150(m *base.Module, l0 int32) int32

//go:linkname Fn1151 github.com/goccy/perlwasm2go/p0.Fn1151
func Fn1151(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1152 github.com/goccy/perlwasm2go/p0.Fn1152
func Fn1152(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1154 github.com/goccy/perlwasm2go/p0.Fn1154
func Fn1154(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1156 github.com/goccy/perlwasm2go/p0.Fn1156
func Fn1156(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1157 github.com/goccy/perlwasm2go/p0.Fn1157
func Fn1157(m *base.Module, l0 int32) int32

//go:linkname Fn1158 github.com/goccy/perlwasm2go/p0.Fn1158
func Fn1158(m *base.Module, l0 int32) int32

//go:linkname Fn1160 github.com/goccy/perlwasm2go/p0.Fn1160
func Fn1160(m *base.Module, l0 int32) int32

//go:linkname Fn1161 github.com/goccy/perlwasm2go/p0.Fn1161
func Fn1161(m *base.Module, l0 int32) int32

//go:linkname Fn1165 github.com/goccy/perlwasm2go/p0.Fn1165
func Fn1165(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1167 github.com/goccy/perlwasm2go/p0.Fn1167
func Fn1167(m *base.Module, l0 int32) int32

//go:linkname Fn1171 github.com/goccy/perlwasm2go/p0.Fn1171
func Fn1171(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1172 github.com/goccy/perlwasm2go/p0.Fn1172
func Fn1172(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1177 github.com/goccy/perlwasm2go/p0.Fn1177
func Fn1177(m *base.Module, l0 int32) int32

//go:linkname Fn1178 github.com/goccy/perlwasm2go/p0.Fn1178
func Fn1178(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1183 github.com/goccy/perlwasm2go/p0.Fn1183
func Fn1183(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1184 github.com/goccy/perlwasm2go/p0.Fn1184
func Fn1184(m *base.Module, l0 int32)

//go:linkname Fn1187 github.com/goccy/perlwasm2go/p0.Fn1187
func Fn1187(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1194 github.com/goccy/perlwasm2go/p0.Fn1194
func Fn1194(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1195 github.com/goccy/perlwasm2go/p0.Fn1195
func Fn1195(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1197 github.com/goccy/perlwasm2go/p0.Fn1197
func Fn1197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1198 github.com/goccy/perlwasm2go/p0.Fn1198
func Fn1198(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1199 github.com/goccy/perlwasm2go/p0.Fn1199
func Fn1199(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1200 github.com/goccy/perlwasm2go/p0.Fn1200
func Fn1200(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1201 github.com/goccy/perlwasm2go/p0.Fn1201
func Fn1201(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1203 github.com/goccy/perlwasm2go/p0.Fn1203
func Fn1203(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1214 github.com/goccy/perlwasm2go/p0.Fn1214
func Fn1214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1228 github.com/goccy/perlwasm2go/p0.Fn1228
func Fn1228(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1234 github.com/goccy/perlwasm2go/p0.Fn1234
func Fn1234(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1306 github.com/goccy/perlwasm2go/p0.Fn1306
func Fn1306(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1317 github.com/goccy/perlwasm2go/p0.Fn1317
func Fn1317(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1328 github.com/goccy/perlwasm2go/p0.Fn1328
func Fn1328(m *base.Module, l0 int32) int32

//go:linkname Fn1331 github.com/goccy/perlwasm2go/p0.Fn1331
func Fn1331(m *base.Module)

//go:linkname Fn1332 github.com/goccy/perlwasm2go/p0.Fn1332
func Fn1332(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1333 github.com/goccy/perlwasm2go/p0.Fn1333
func Fn1333(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1337 github.com/goccy/perlwasm2go/p0.Fn1337
func Fn1337(m *base.Module, l0 int32) int32

//go:linkname Fn1339 github.com/goccy/perlwasm2go/p0.Fn1339
func Fn1339(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1353 github.com/goccy/perlwasm2go/p0.Fn1353
func Fn1353(m *base.Module, l0 int32)

//go:linkname Fn1355 github.com/goccy/perlwasm2go/p0.Fn1355
func Fn1355(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1356 github.com/goccy/perlwasm2go/p0.Fn1356
func Fn1356(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1360 github.com/goccy/perlwasm2go/p0.Fn1360
func Fn1360(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1362 github.com/goccy/perlwasm2go/p0.Fn1362
func Fn1362(m *base.Module, l0 int32)

//go:linkname Fn1370 github.com/goccy/perlwasm2go/p0.Fn1370
func Fn1370(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1372 github.com/goccy/perlwasm2go/p0.Fn1372
func Fn1372(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1373 github.com/goccy/perlwasm2go/p0.Fn1373
func Fn1373(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1376 github.com/goccy/perlwasm2go/p0.Fn1376
func Fn1376(m *base.Module, l0 int32) int32

//go:linkname Fn1424 github.com/goccy/perlwasm2go/p0.Fn1424
func Fn1424(m *base.Module, l0 int32)

//go:linkname Fn1425 github.com/goccy/perlwasm2go/p0.Fn1425
func Fn1425(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1426 github.com/goccy/perlwasm2go/p0.Fn1426
func Fn1426(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1427 github.com/goccy/perlwasm2go/p0.Fn1427
func Fn1427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1428 github.com/goccy/perlwasm2go/p0.Fn1428
func Fn1428(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1429 github.com/goccy/perlwasm2go/p0.Fn1429
func Fn1429(m *base.Module, l0 int32) int32

//go:linkname Fn1430 github.com/goccy/perlwasm2go/p0.Fn1430
func Fn1430(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1432 github.com/goccy/perlwasm2go/p0.Fn1432
func Fn1432(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1433 github.com/goccy/perlwasm2go/p0.Fn1433
func Fn1433(m *base.Module, l0 int32)

//go:linkname Fn1437 github.com/goccy/perlwasm2go/p0.Fn1437
func Fn1437(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1438 github.com/goccy/perlwasm2go/p0.Fn1438
func Fn1438(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1441 github.com/goccy/perlwasm2go/p0.Fn1441
func Fn1441(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1442 github.com/goccy/perlwasm2go/p0.Fn1442
func Fn1442(m *base.Module, l0 int32) int32

//go:linkname Fn1443 github.com/goccy/perlwasm2go/p0.Fn1443
func Fn1443(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1448 github.com/goccy/perlwasm2go/p0.Fn1448
func Fn1448(m *base.Module)

//go:linkname Fn1449 github.com/goccy/perlwasm2go/p0.Fn1449
func Fn1449(m *base.Module)

//go:linkname Fn1461 github.com/goccy/perlwasm2go/p0.Fn1461
func Fn1461(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1463 github.com/goccy/perlwasm2go/p0.Fn1463
func Fn1463(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1470 github.com/goccy/perlwasm2go/p0.Fn1470
func Fn1470(m *base.Module, l0 int32)

//go:linkname Fn1471 github.com/goccy/perlwasm2go/p0.Fn1471
func Fn1471(m *base.Module, l0 int32)

//go:linkname Fn1478 github.com/goccy/perlwasm2go/p0.Fn1478
func Fn1478(m *base.Module, l0 int32) int32

//go:linkname Fn1489 github.com/goccy/perlwasm2go/p0.Fn1489
func Fn1489(m *base.Module, l0 int32) int32

//go:linkname Fn1517 github.com/goccy/perlwasm2go/p0.Fn1517
func Fn1517(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1519 github.com/goccy/perlwasm2go/p0.Fn1519
func Fn1519(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1534 github.com/goccy/perlwasm2go/p0.Fn1534
func Fn1534(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1535 github.com/goccy/perlwasm2go/p0.Fn1535
func Fn1535(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1541 github.com/goccy/perlwasm2go/p0.Fn1541
func Fn1541(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1549 github.com/goccy/perlwasm2go/p0.Fn1549
func Fn1549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1562 github.com/goccy/perlwasm2go/p0.Fn1562
func Fn1562(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1564 github.com/goccy/perlwasm2go/p0.Fn1564
func Fn1564(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1567 github.com/goccy/perlwasm2go/p0.Fn1567
func Fn1567(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1568 github.com/goccy/perlwasm2go/p0.Fn1568
func Fn1568(m *base.Module, l0 int32) int32

//go:linkname Fn1570 github.com/goccy/perlwasm2go/p0.Fn1570
func Fn1570(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1571 github.com/goccy/perlwasm2go/p0.Fn1571
func Fn1571(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1574 github.com/goccy/perlwasm2go/p0.Fn1574
func Fn1574(m *base.Module, l0 int32) int32

//go:linkname Fn1576 github.com/goccy/perlwasm2go/p0.Fn1576
func Fn1576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1578 github.com/goccy/perlwasm2go/p0.Fn1578
func Fn1578(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1579 github.com/goccy/perlwasm2go/p0.Fn1579
func Fn1579(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1581 github.com/goccy/perlwasm2go/p0.Fn1581
func Fn1581(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1583 github.com/goccy/perlwasm2go/p0.Fn1583
func Fn1583(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1584 github.com/goccy/perlwasm2go/p0.Fn1584
func Fn1584(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1585 github.com/goccy/perlwasm2go/p0.Fn1585
func Fn1585(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1587 github.com/goccy/perlwasm2go/p0.Fn1587
func Fn1587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1588 github.com/goccy/perlwasm2go/p0.Fn1588
func Fn1588(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1592 github.com/goccy/perlwasm2go/p0.Fn1592
func Fn1592(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1593 github.com/goccy/perlwasm2go/p0.Fn1593
func Fn1593(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1594 github.com/goccy/perlwasm2go/p0.Fn1594
func Fn1594(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1597 github.com/goccy/perlwasm2go/p0.Fn1597
func Fn1597(m *base.Module, l0 int32) int32

//go:linkname Fn1598 github.com/goccy/perlwasm2go/p0.Fn1598
func Fn1598(m *base.Module, l0 int32)

//go:linkname Fn1600 github.com/goccy/perlwasm2go/p0.Fn1600
func Fn1600(m *base.Module, l0 int32) int32

//go:linkname Fn1601 github.com/goccy/perlwasm2go/p0.Fn1601
func Fn1601(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1603 github.com/goccy/perlwasm2go/p0.Fn1603
func Fn1603(m *base.Module, l0 int32)

//go:linkname Fn1605 github.com/goccy/perlwasm2go/p0.Fn1605
func Fn1605(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1608 github.com/goccy/perlwasm2go/p0.Fn1608
func Fn1608(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn1609 github.com/goccy/perlwasm2go/p0.Fn1609
func Fn1609(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1611 github.com/goccy/perlwasm2go/p0.Fn1611
func Fn1611(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1612 github.com/goccy/perlwasm2go/p0.Fn1612
func Fn1612(m *base.Module) int32

//go:linkname Fn1615 github.com/goccy/perlwasm2go/p0.Fn1615
func Fn1615(m *base.Module, l0 int32) int32

//go:linkname Fn1617 github.com/goccy/perlwasm2go/p0.Fn1617
func Fn1617(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1619 github.com/goccy/perlwasm2go/p0.Fn1619
func Fn1619(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1621 github.com/goccy/perlwasm2go/p0.Fn1621
func Fn1621(m *base.Module, l0 int32)

//go:linkname Fn1622 github.com/goccy/perlwasm2go/p0.Fn1622
func Fn1622(m *base.Module, l0 int32)

//go:linkname Fn1625 github.com/goccy/perlwasm2go/p0.Fn1625
func Fn1625(m *base.Module, l0 int32) int32

//go:linkname Fn1626 github.com/goccy/perlwasm2go/p0.Fn1626
func Fn1626(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1627 github.com/goccy/perlwasm2go/p0.Fn1627
func Fn1627(m *base.Module, l0 int32)

//go:linkname Fn1628 github.com/goccy/perlwasm2go/p0.Fn1628
func Fn1628(m *base.Module, l0 int32) int32

//go:linkname Fn1629 github.com/goccy/perlwasm2go/p0.Fn1629
func Fn1629(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1630 github.com/goccy/perlwasm2go/p0.Fn1630
func Fn1630(m *base.Module, l0 int32) int32

//go:linkname Fn1632 github.com/goccy/perlwasm2go/p0.Fn1632
func Fn1632(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1634 github.com/goccy/perlwasm2go/p0.Fn1634
func Fn1634(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1635 github.com/goccy/perlwasm2go/p0.Fn1635
func Fn1635(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1636 github.com/goccy/perlwasm2go/p0.Fn1636
func Fn1636(m *base.Module, l0 int32) int32

//go:linkname Fn1638 github.com/goccy/perlwasm2go/p0.Fn1638
func Fn1638(m *base.Module, l0 int32)

//go:linkname Fn1642 github.com/goccy/perlwasm2go/p0.Fn1642
func Fn1642(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1643 github.com/goccy/perlwasm2go/p0.Fn1643
func Fn1643(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1644 github.com/goccy/perlwasm2go/p0.Fn1644
func Fn1644(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1648 github.com/goccy/perlwasm2go/p0.Fn1648
func Fn1648(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1650 github.com/goccy/perlwasm2go/p0.Fn1650
func Fn1650(m *base.Module, l0 int32) int32

//go:linkname Fn1651 github.com/goccy/perlwasm2go/p0.Fn1651
func Fn1651(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1653 github.com/goccy/perlwasm2go/p0.Fn1653
func Fn1653(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1655 github.com/goccy/perlwasm2go/p0.Fn1655
func Fn1655(m *base.Module, l0 int32)

//go:linkname Fn1658 github.com/goccy/perlwasm2go/p0.Fn1658
func Fn1658(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1659 github.com/goccy/perlwasm2go/p0.Fn1659
func Fn1659(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1660 github.com/goccy/perlwasm2go/p0.Fn1660
func Fn1660(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1661 github.com/goccy/perlwasm2go/p0.Fn1661
func Fn1661(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1662 github.com/goccy/perlwasm2go/p0.Fn1662
func Fn1662(m *base.Module, l0 int32) int32

//go:linkname Fn1665 github.com/goccy/perlwasm2go/p0.Fn1665
func Fn1665(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1684 github.com/goccy/perlwasm2go/p0.Fn1684
func Fn1684(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1691 github.com/goccy/perlwasm2go/p0.Fn1691
func Fn1691(m *base.Module, l0 int32)

//go:linkname Fn1695 github.com/goccy/perlwasm2go/p0.Fn1695
func Fn1695(m *base.Module, l0 int32)

//go:linkname Fn1696 github.com/goccy/perlwasm2go/p0.Fn1696
func Fn1696(m *base.Module, l0 int32) int32

//go:linkname Fn1697 github.com/goccy/perlwasm2go/p0.Fn1697
func Fn1697(m *base.Module, l0 int32)

//go:linkname Fn1703 github.com/goccy/perlwasm2go/p0.Fn1703
func Fn1703(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1704 github.com/goccy/perlwasm2go/p0.Fn1704
func Fn1704(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1705 github.com/goccy/perlwasm2go/p0.Fn1705
func Fn1705(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1706 github.com/goccy/perlwasm2go/p0.Fn1706
func Fn1706(m *base.Module, l0 int32)

//go:linkname Fn1729 github.com/goccy/perlwasm2go/p0.Fn1729
func Fn1729(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1738 github.com/goccy/perlwasm2go/p0.Fn1738
func Fn1738(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1741 github.com/goccy/perlwasm2go/p0.Fn1741
func Fn1741(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1743 github.com/goccy/perlwasm2go/p0.Fn1743
func Fn1743(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1744 github.com/goccy/perlwasm2go/p0.Fn1744
func Fn1744(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1749 github.com/goccy/perlwasm2go/p0.Fn1749
func Fn1749(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1750 github.com/goccy/perlwasm2go/p0.Fn1750
func Fn1750(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1794 github.com/goccy/perlwasm2go/p0.Fn1794
func Fn1794(m *base.Module, l0 int32) int32

//go:linkname Fn1795 github.com/goccy/perlwasm2go/p0.Fn1795
func Fn1795(m *base.Module, l0 int32) int32

//go:linkname Fn1796 github.com/goccy/perlwasm2go/p0.Fn1796
func Fn1796(m *base.Module, l0 int32) int32

//go:linkname Fn1797 github.com/goccy/perlwasm2go/p0.Fn1797
func Fn1797(m *base.Module, l0 int32)

//go:linkname Fn1799 github.com/goccy/perlwasm2go/p0.Fn1799
func Fn1799(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1802 github.com/goccy/perlwasm2go/p0.Fn1802
func Fn1802(m *base.Module, l0 int32)

//go:linkname Fn1808 github.com/goccy/perlwasm2go/p0.Fn1808
func Fn1808(m *base.Module, l0 int32)

//go:linkname Fn1809 github.com/goccy/perlwasm2go/p0.Fn1809
func Fn1809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1810 github.com/goccy/perlwasm2go/p0.Fn1810
func Fn1810(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1814 github.com/goccy/perlwasm2go/p0.Fn1814
func Fn1814(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1815 github.com/goccy/perlwasm2go/p0.Fn1815
func Fn1815(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1817 github.com/goccy/perlwasm2go/p0.Fn1817
func Fn1817(m *base.Module, l0 int32) float64

//go:linkname Fn1820 github.com/goccy/perlwasm2go/p0.Fn1820
func Fn1820(m *base.Module, l0 int32) int32

//go:linkname Fn1821 github.com/goccy/perlwasm2go/p0.Fn1821
func Fn1821(m *base.Module, l0 int32) int32

//go:linkname Fn1822 github.com/goccy/perlwasm2go/p0.Fn1822
func Fn1822(m *base.Module, l0 int32) int32

//go:linkname Fn1824 github.com/goccy/perlwasm2go/p0.Fn1824
func Fn1824(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1826 github.com/goccy/perlwasm2go/p0.Fn1826
func Fn1826(m *base.Module, l0 int32)

//go:linkname Fn1827 github.com/goccy/perlwasm2go/p0.Fn1827
func Fn1827(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1828 github.com/goccy/perlwasm2go/p0.Fn1828
func Fn1828(m *base.Module, l0 int32)

//go:linkname Fn1829 github.com/goccy/perlwasm2go/p0.Fn1829
func Fn1829(m *base.Module, l0 int32)

//go:linkname Fn1830 github.com/goccy/perlwasm2go/p0.Fn1830
func Fn1830(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1831 github.com/goccy/perlwasm2go/p0.Fn1831
func Fn1831(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1833 github.com/goccy/perlwasm2go/p0.Fn1833
func Fn1833(m *base.Module, l0 int32) int32

//go:linkname Fn1838 github.com/goccy/perlwasm2go/p0.Fn1838
func Fn1838(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1840 github.com/goccy/perlwasm2go/p0.Fn1840
func Fn1840(m *base.Module, l0 int32)

//go:linkname Fn1841 github.com/goccy/perlwasm2go/p0.Fn1841
func Fn1841(m *base.Module, l0 int32) int32

//go:linkname Fn1842 github.com/goccy/perlwasm2go/p0.Fn1842
func Fn1842(m *base.Module, l0 int32) int32

//go:linkname Fn1849 github.com/goccy/perlwasm2go/p0.Fn1849
func Fn1849(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1850 github.com/goccy/perlwasm2go/p0.Fn1850
func Fn1850(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1871 github.com/goccy/perlwasm2go/p0.Fn1871
func Fn1871(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1876 github.com/goccy/perlwasm2go/p0.Fn1876
func Fn1876(m *base.Module, l0 int32)

//go:linkname Fn1880 github.com/goccy/perlwasm2go/p0.Fn1880
func Fn1880(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1884 github.com/goccy/perlwasm2go/p0.Fn1884
func Fn1884(m *base.Module, l0 int32)

//go:linkname Fn1888 github.com/goccy/perlwasm2go/p0.Fn1888
func Fn1888(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1893 github.com/goccy/perlwasm2go/p0.Fn1893
func Fn1893(m *base.Module) int32

//go:linkname Fn1895 github.com/goccy/perlwasm2go/p0.Fn1895
func Fn1895(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1896 github.com/goccy/perlwasm2go/p0.Fn1896
func Fn1896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1899 github.com/goccy/perlwasm2go/p0.Fn1899
func Fn1899(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1900 github.com/goccy/perlwasm2go/p0.Fn1900
func Fn1900(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1905 github.com/goccy/perlwasm2go/p0.Fn1905
func Fn1905(m *base.Module, l0 int32) int32

//go:linkname Fn1907 github.com/goccy/perlwasm2go/p0.Fn1907
func Fn1907(m *base.Module, l0 int32) int32

//go:linkname Fn2003 github.com/goccy/perlwasm2go/p0.Fn2003
func Fn2003(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2004 github.com/goccy/perlwasm2go/p0.Fn2004
func Fn2004(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2022 github.com/goccy/perlwasm2go/p0.Fn2022
func Fn2022(m *base.Module, l0 int32) int32

//go:linkname Fn2027 github.com/goccy/perlwasm2go/p0.Fn2027
func Fn2027(m *base.Module, l0 int32)

//go:linkname Fn2032 github.com/goccy/perlwasm2go/p0.Fn2032
func Fn2032(m *base.Module, l0 int32) int32

//go:linkname Fn2036 github.com/goccy/perlwasm2go/p0.Fn2036
func Fn2036(m *base.Module, l0 int32)

//go:linkname Fn2148 github.com/goccy/perlwasm2go/p0.Fn2148
func Fn2148(m *base.Module, l0 int32) int32

//go:linkname Fn2199 github.com/goccy/perlwasm2go/p0.Fn2199
func Fn2199(m *base.Module, l0 int32)

//go:linkname Fn2211 github.com/goccy/perlwasm2go/p0.Fn2211
func Fn2211(m *base.Module, l0 int32)

//go:linkname Fn2212 github.com/goccy/perlwasm2go/p0.Fn2212
func Fn2212(m *base.Module, l0 int32)

//go:linkname Fn2213 github.com/goccy/perlwasm2go/p0.Fn2213
func Fn2213(m *base.Module, l0 int32)

//go:linkname Fn2214 github.com/goccy/perlwasm2go/p0.Fn2214
func Fn2214(m *base.Module, l0 int32)

//go:linkname Fn2215 github.com/goccy/perlwasm2go/p0.Fn2215
func Fn2215(m *base.Module, l0 int32)

//go:linkname Fn2218 github.com/goccy/perlwasm2go/p0.Fn2218
func Fn2218(m *base.Module, l0 int32)

//go:linkname Fn2220 github.com/goccy/perlwasm2go/p0.Fn2220
func Fn2220(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2226 github.com/goccy/perlwasm2go/p0.Fn2226
func Fn2226(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2254 github.com/goccy/perlwasm2go/p0.Fn2254
func Fn2254(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2263 github.com/goccy/perlwasm2go/p0.Fn2263
func Fn2263(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2302 github.com/goccy/perlwasm2go/p0.Fn2302
func Fn2302(m *base.Module, l0 int32) int32

//go:linkname Fn2334 github.com/goccy/perlwasm2go/p0.Fn2334
func Fn2334(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2346 github.com/goccy/perlwasm2go/p0.Fn2346
func Fn2346(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2353 github.com/goccy/perlwasm2go/p0.Fn2353
func Fn2353(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2402 github.com/goccy/perlwasm2go/p0.Fn2402
func Fn2402(m *base.Module, l0 int32)

//go:linkname Fn2483 github.com/goccy/perlwasm2go/p0.Fn2483
func Fn2483(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2491 github.com/goccy/perlwasm2go/p0.Fn2491
func Fn2491(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn2494 github.com/goccy/perlwasm2go/p0.Fn2494
func Fn2494(m *base.Module, l0 int32) int32

//go:linkname Fn2532 github.com/goccy/perlwasm2go/p0.Fn2532
func Fn2532(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn2538 github.com/goccy/perlwasm2go/p0.Fn2538
func Fn2538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2565 github.com/goccy/perlwasm2go/p0.Fn2565
func Fn2565(m *base.Module, l0 int32) int32

//go:linkname Fn2569 github.com/goccy/perlwasm2go/p0.Fn2569
func Fn2569(m *base.Module) int32

//go:linkname Fn2574 github.com/goccy/perlwasm2go/p0.Fn2574
func Fn2574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2604 github.com/goccy/perlwasm2go/p0.Fn2604
func Fn2604(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2609 github.com/goccy/perlwasm2go/p0.Fn2609
func Fn2609(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2653 github.com/goccy/perlwasm2go/p0.Fn2653
func Fn2653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2654 github.com/goccy/perlwasm2go/p0.Fn2654
func Fn2654(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2655 github.com/goccy/perlwasm2go/p0.Fn2655
func Fn2655(m *base.Module, l0 int32) int32

//go:linkname Fn2656 github.com/goccy/perlwasm2go/p0.Fn2656
func Fn2656(m *base.Module) int32

//go:linkname Fn2657 github.com/goccy/perlwasm2go/p0.Fn2657
func Fn2657(m *base.Module)

//go:linkname Fn2658 github.com/goccy/perlwasm2go/p0.Fn2658
func Fn2658(m *base.Module)

//go:linkname Fn2659 github.com/goccy/perlwasm2go/p0.Fn2659
func Fn2659(m *base.Module, l0 int32)

//go:linkname Fn2660 github.com/goccy/perlwasm2go/p0.Fn2660
func Fn2660(m *base.Module, l0 int32)

//go:linkname Fn2661 github.com/goccy/perlwasm2go/p0.Fn2661
func Fn2661(m *base.Module) int32

//go:linkname Fn2662 github.com/goccy/perlwasm2go/p0.Fn2662
func Fn2662(m *base.Module, l0 int32) int32

//go:linkname Fn2663 github.com/goccy/perlwasm2go/p0.Fn2663
func Fn2663(m *base.Module)

//go:linkname Fn2664 github.com/goccy/perlwasm2go/p0.Fn2664
func Fn2664(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2665 github.com/goccy/perlwasm2go/p0.Fn2665
func Fn2665(m *base.Module, l0 int32) int32

//go:linkname Fn2666 github.com/goccy/perlwasm2go/p0.Fn2666
func Fn2666(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2667 github.com/goccy/perlwasm2go/p0.Fn2667
func Fn2667(m *base.Module, l0 int32)

//go:linkname Fn2668 github.com/goccy/perlwasm2go/p0.Fn2668
func Fn2668(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2673 github.com/goccy/perlwasm2go/p0.Fn2673
func Fn2673(m *base.Module, l0 int32)

//go:linkname Fn2674 github.com/goccy/perlwasm2go/p0.Fn2674
func Fn2674(m *base.Module, l0 int32)

//go:linkname Fn2675 github.com/goccy/perlwasm2go/p0.Fn2675
func Fn2675(m *base.Module, l0 int32)

//go:linkname Fn2676 github.com/goccy/perlwasm2go/p0.Fn2676
func Fn2676(m *base.Module, l0 int32)

//go:linkname Fn2678 github.com/goccy/perlwasm2go/p0.Fn2678
func Fn2678(m *base.Module, l0 int32)

//go:linkname Fn2679 github.com/goccy/perlwasm2go/p0.Fn2679
func Fn2679(m *base.Module, l0 int32)

//go:linkname Fn2681 github.com/goccy/perlwasm2go/p0.Fn2681
func Fn2681(m *base.Module, l0 int32)

//go:linkname Fn2682 github.com/goccy/perlwasm2go/p0.Fn2682
func Fn2682(m *base.Module, l0 int32)

//go:linkname Fn2689 github.com/goccy/perlwasm2go/p0.Fn2689
func Fn2689(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2690 github.com/goccy/perlwasm2go/p0.Fn2690
func Fn2690(m *base.Module)

//go:linkname Fn2693 github.com/goccy/perlwasm2go/p0.Fn2693
func Fn2693(m *base.Module)

//go:linkname Fn2697 github.com/goccy/perlwasm2go/p0.Fn2697
func Fn2697(m *base.Module) int32

//go:linkname Fn2703 github.com/goccy/perlwasm2go/p0.Fn2703
func Fn2703(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2704 github.com/goccy/perlwasm2go/p0.Fn2704
func Fn2704(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2705 github.com/goccy/perlwasm2go/p0.Fn2705
func Fn2705(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2707 github.com/goccy/perlwasm2go/p0.Fn2707
func Fn2707(m *base.Module, l0 int32) int32

//go:linkname Fn2708 github.com/goccy/perlwasm2go/p0.Fn2708
func Fn2708(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2709 github.com/goccy/perlwasm2go/p0.Fn2709
func Fn2709(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2711 github.com/goccy/perlwasm2go/p0.Fn2711
func Fn2711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2712 github.com/goccy/perlwasm2go/p0.Fn2712
func Fn2712(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2713 github.com/goccy/perlwasm2go/p0.Fn2713
func Fn2713(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2714 github.com/goccy/perlwasm2go/p0.Fn2714
func Fn2714(m *base.Module) int32

//go:linkname Fn2715 github.com/goccy/perlwasm2go/p0.Fn2715
func Fn2715(m *base.Module, l0 int32) int32

//go:linkname Fn2716 github.com/goccy/perlwasm2go/p0.Fn2716
func Fn2716(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2717 github.com/goccy/perlwasm2go/p0.Fn2717
func Fn2717(m *base.Module, l0 int32) int32

//go:linkname Fn2718 github.com/goccy/perlwasm2go/p0.Fn2718
func Fn2718(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2719 github.com/goccy/perlwasm2go/p0.Fn2719
func Fn2719(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2720 github.com/goccy/perlwasm2go/p0.Fn2720
func Fn2720(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2721 github.com/goccy/perlwasm2go/p0.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2722 github.com/goccy/perlwasm2go/p0.Fn2722
func Fn2722(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2723 github.com/goccy/perlwasm2go/p0.Fn2723
func Fn2723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2724 github.com/goccy/perlwasm2go/p0.Fn2724
func Fn2724(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2725 github.com/goccy/perlwasm2go/p0.Fn2725
func Fn2725(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2727 github.com/goccy/perlwasm2go/p0.Fn2727
func Fn2727(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2728 github.com/goccy/perlwasm2go/p0.Fn2728
func Fn2728(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2729 github.com/goccy/perlwasm2go/p0.Fn2729
func Fn2729(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn2735 github.com/goccy/perlwasm2go/p0.Fn2735
func Fn2735(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2739 github.com/goccy/perlwasm2go/p0.Fn2739
func Fn2739(m *base.Module, l0 int32)

//go:linkname Fn2740 github.com/goccy/perlwasm2go/p0.Fn2740
func Fn2740(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2743 github.com/goccy/perlwasm2go/p0.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2745 github.com/goccy/perlwasm2go/p0.Fn2745
func Fn2745(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2748 github.com/goccy/perlwasm2go/p0.Fn2748
func Fn2748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2749 github.com/goccy/perlwasm2go/p0.Fn2749
func Fn2749(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn2750 github.com/goccy/perlwasm2go/p0.Fn2750
func Fn2750(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2751 github.com/goccy/perlwasm2go/p0.Fn2751
func Fn2751(m *base.Module, l0 int32) int32

//go:linkname Fn2753 github.com/goccy/perlwasm2go/p0.Fn2753
func Fn2753(m *base.Module, l0 int32) int32

//go:linkname Fn2755 github.com/goccy/perlwasm2go/p0.Fn2755
func Fn2755(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2756 github.com/goccy/perlwasm2go/p0.Fn2756
func Fn2756(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2757 github.com/goccy/perlwasm2go/p0.Fn2757
func Fn2757(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2761 github.com/goccy/perlwasm2go/p0.Fn2761
func Fn2761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2762 github.com/goccy/perlwasm2go/p0.Fn2762
func Fn2762(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2763 github.com/goccy/perlwasm2go/p0.Fn2763
func Fn2763(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2764 github.com/goccy/perlwasm2go/p0.Fn2764
func Fn2764(m *base.Module, l0 int32) int32

//go:linkname Fn2766 github.com/goccy/perlwasm2go/p0.Fn2766
func Fn2766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2768 github.com/goccy/perlwasm2go/p0.Fn2768
func Fn2768(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2769 github.com/goccy/perlwasm2go/p0.Fn2769
func Fn2769(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2770 github.com/goccy/perlwasm2go/p0.Fn2770
func Fn2770(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2773 github.com/goccy/perlwasm2go/p0.Fn2773
func Fn2773(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2774 github.com/goccy/perlwasm2go/p0.Fn2774
func Fn2774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2784 github.com/goccy/perlwasm2go/p0.Fn2784
func Fn2784(m *base.Module, l0 int32)

//go:linkname Fn2785 github.com/goccy/perlwasm2go/p0.Fn2785
func Fn2785(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2786 github.com/goccy/perlwasm2go/p0.Fn2786
func Fn2786(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2788 github.com/goccy/perlwasm2go/p0.Fn2788
func Fn2788(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2790 github.com/goccy/perlwasm2go/p0.Fn2790
func Fn2790(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2792 github.com/goccy/perlwasm2go/p0.Fn2792
func Fn2792(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2793 github.com/goccy/perlwasm2go/p0.Fn2793
func Fn2793(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2794 github.com/goccy/perlwasm2go/p0.Fn2794
func Fn2794(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2795 github.com/goccy/perlwasm2go/p0.Fn2795
func Fn2795(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2796 github.com/goccy/perlwasm2go/p0.Fn2796
func Fn2796(m *base.Module, l0 int32) int32

//go:linkname Fn2797 github.com/goccy/perlwasm2go/p0.Fn2797
func Fn2797(m *base.Module, l0 int32)

//go:linkname Fn2798 github.com/goccy/perlwasm2go/p0.Fn2798
func Fn2798(m *base.Module, l0 int32) int32

//go:linkname Fn2800 github.com/goccy/perlwasm2go/p0.Fn2800
func Fn2800(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2801 github.com/goccy/perlwasm2go/p0.Fn2801
func Fn2801(m *base.Module, l0 int32) int32

//go:linkname Fn2802 github.com/goccy/perlwasm2go/p0.Fn2802
func Fn2802(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2804 github.com/goccy/perlwasm2go/p0.Fn2804
func Fn2804(m *base.Module, l0 int32) int32

//go:linkname Fn2806 github.com/goccy/perlwasm2go/p0.Fn2806
func Fn2806(m *base.Module, l0 int32) int32

//go:linkname Fn2807 github.com/goccy/perlwasm2go/p0.Fn2807
func Fn2807(m *base.Module, l0 int32) int32

//go:linkname Fn2808 github.com/goccy/perlwasm2go/p0.Fn2808
func Fn2808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2812 github.com/goccy/perlwasm2go/p0.Fn2812
func Fn2812(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2816 github.com/goccy/perlwasm2go/p0.Fn2816
func Fn2816(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2825 github.com/goccy/perlwasm2go/p0.Fn2825
func Fn2825(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2826 github.com/goccy/perlwasm2go/p0.Fn2826
func Fn2826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2827 github.com/goccy/perlwasm2go/p0.Fn2827
func Fn2827(m *base.Module, l0 int32) int32

//go:linkname Fn2828 github.com/goccy/perlwasm2go/p0.Fn2828
func Fn2828(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2829 github.com/goccy/perlwasm2go/p0.Fn2829
func Fn2829(m *base.Module, l0 int32) int32

//go:linkname Fn2835 github.com/goccy/perlwasm2go/p0.Fn2835
func Fn2835(m *base.Module, l0 int32) int32

//go:linkname Fn2837 github.com/goccy/perlwasm2go/p0.Fn2837
func Fn2837(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2838 github.com/goccy/perlwasm2go/p0.Fn2838
func Fn2838(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2841 github.com/goccy/perlwasm2go/p0.Fn2841
func Fn2841(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2844 github.com/goccy/perlwasm2go/p0.Fn2844
func Fn2844(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2845 github.com/goccy/perlwasm2go/p0.Fn2845
func Fn2845(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2849 github.com/goccy/perlwasm2go/p0.Fn2849
func Fn2849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2858 github.com/goccy/perlwasm2go/p0.Fn2858
func Fn2858(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2861 github.com/goccy/perlwasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2862 github.com/goccy/perlwasm2go/p0.Fn2862
func Fn2862(m *base.Module, l0 int32) int32

//go:linkname Fn2863 github.com/goccy/perlwasm2go/p0.Fn2863
func Fn2863(m *base.Module, l0 int32)

//go:linkname Fn2864 github.com/goccy/perlwasm2go/p0.Fn2864
func Fn2864(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2875 github.com/goccy/perlwasm2go/p0.Fn2875
func Fn2875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2880 github.com/goccy/perlwasm2go/p0.Fn2880
func Fn2880(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2890 github.com/goccy/perlwasm2go/p0.Fn2890
func Fn2890(m *base.Module, l0 int32)

//go:linkname Fn2917 github.com/goccy/perlwasm2go/p0.Fn2917
func Fn2917(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2918 github.com/goccy/perlwasm2go/p0.Fn2918
func Fn2918(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2925 github.com/goccy/perlwasm2go/p0.Fn2925
func Fn2925(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2942 github.com/goccy/perlwasm2go/p0.Fn2942
func Fn2942(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2943 github.com/goccy/perlwasm2go/p0.Fn2943
func Fn2943(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2945 github.com/goccy/perlwasm2go/p0.Fn2945
func Fn2945(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2946 github.com/goccy/perlwasm2go/p0.Fn2946
func Fn2946(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2947 github.com/goccy/perlwasm2go/p0.Fn2947
func Fn2947(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2951 github.com/goccy/perlwasm2go/p0.Fn2951
func Fn2951(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2953 github.com/goccy/perlwasm2go/p0.Fn2953
func Fn2953(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2954 github.com/goccy/perlwasm2go/p0.Fn2954
func Fn2954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2955 github.com/goccy/perlwasm2go/p0.Fn2955
func Fn2955(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2957 github.com/goccy/perlwasm2go/p0.Fn2957
func Fn2957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2969 github.com/goccy/perlwasm2go/p0.Fn2969
func Fn2969(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2970 github.com/goccy/perlwasm2go/p0.Fn2970
func Fn2970(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2975 github.com/goccy/perlwasm2go/p0.Fn2975
func Fn2975(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2976 github.com/goccy/perlwasm2go/p0.Fn2976
func Fn2976(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2978 github.com/goccy/perlwasm2go/p0.Fn2978
func Fn2978(m *base.Module, l0 int32) int32

//go:linkname Fn2979 github.com/goccy/perlwasm2go/p0.Fn2979
func Fn2979(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2980 github.com/goccy/perlwasm2go/p0.Fn2980
func Fn2980(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2981 github.com/goccy/perlwasm2go/p0.Fn2981
func Fn2981(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2982 github.com/goccy/perlwasm2go/p0.Fn2982
func Fn2982(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2991 github.com/goccy/perlwasm2go/p0.Fn2991
func Fn2991(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2993 github.com/goccy/perlwasm2go/p0.Fn2993
func Fn2993(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2994 github.com/goccy/perlwasm2go/p0.Fn2994
func Fn2994(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2995 github.com/goccy/perlwasm2go/p0.Fn2995
func Fn2995(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2997 github.com/goccy/perlwasm2go/p0.Fn2997
func Fn2997(m *base.Module, l0 int32)

//go:linkname Fn2998 github.com/goccy/perlwasm2go/p0.Fn2998
func Fn2998(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2999 github.com/goccy/perlwasm2go/p0.Fn2999
func Fn2999(m *base.Module)

//go:linkname Fn3001 github.com/goccy/perlwasm2go/p0.Fn3001
func Fn3001(m *base.Module, l0 int32)

//go:linkname Fn3003 github.com/goccy/perlwasm2go/p0.Fn3003
func Fn3003(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3004 github.com/goccy/perlwasm2go/p0.Fn3004
func Fn3004(m *base.Module)

//go:linkname Fn3006 github.com/goccy/perlwasm2go/p0.Fn3006
func Fn3006(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3007 github.com/goccy/perlwasm2go/p0.Fn3007
func Fn3007(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3011 github.com/goccy/perlwasm2go/p0.Fn3011
func Fn3011(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3012 github.com/goccy/perlwasm2go/p0.Fn3012
func Fn3012(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3014 github.com/goccy/perlwasm2go/p0.Fn3014
func Fn3014(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3022 github.com/goccy/perlwasm2go/p0.Fn3022
func Fn3022(m *base.Module, l0 int32) int32

//go:linkname Fn3029 github.com/goccy/perlwasm2go/p0.Fn3029
func Fn3029(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3032 github.com/goccy/perlwasm2go/p0.Fn3032
func Fn3032(m *base.Module, l0 int32) int32

//go:linkname Fn3033 github.com/goccy/perlwasm2go/p0.Fn3033
func Fn3033(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3034 github.com/goccy/perlwasm2go/p0.Fn3034
func Fn3034(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3035 github.com/goccy/perlwasm2go/p0.Fn3035
func Fn3035(m *base.Module, l0 int32) int32

//go:linkname Fn3036 github.com/goccy/perlwasm2go/p0.Fn3036
func Fn3036(m *base.Module, l0 int32) int32

//go:linkname Fn3039 github.com/goccy/perlwasm2go/p0.Fn3039
func Fn3039(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3077 github.com/goccy/perlwasm2go/p0.Fn3077
func Fn3077(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3124 github.com/goccy/perlwasm2go/p0.Fn3124
func Fn3124(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3190 github.com/goccy/perlwasm2go/p0.Fn3190
func Fn3190(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3375 github.com/goccy/perlwasm2go/p0.Fn3375
func Fn3375(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3378 github.com/goccy/perlwasm2go/p0.Fn3378
func Fn3378(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3480 github.com/goccy/perlwasm2go/p0.Fn3480
func Fn3480(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3485 github.com/goccy/perlwasm2go/p0.Fn3485
func Fn3485(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn3487 github.com/goccy/perlwasm2go/p0.Fn3487
func Fn3487(m *base.Module, l0 int32) int32

//go:linkname Fn3511 github.com/goccy/perlwasm2go/p0.Fn3511
func Fn3511(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3515 github.com/goccy/perlwasm2go/p0.Fn3515
func Fn3515(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3519 github.com/goccy/perlwasm2go/p0.Fn3519
func Fn3519(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3530 github.com/goccy/perlwasm2go/p0.Fn3530
func Fn3530(m *base.Module, l0 int32) int32

//go:linkname Fn3538 github.com/goccy/perlwasm2go/p0.Fn3538
func Fn3538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3550 github.com/goccy/perlwasm2go/p0.Fn3550
func Fn3550(m *base.Module, l0 int32, l1 int32) int32
