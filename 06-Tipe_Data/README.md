# Tipe Data pada GoLang

Ada beberapa jenis tipe data pada GoLang, diantaranya :

- Tipe data Number
- Tipe data Boolean
- Tipe data String
- Nilai `nil`, `nil` sebetulnya bukan tipe data. Variabel dengan nilai nil, artinya nilainya kosong. Biasanya digunakan pada pointer, function, slice, map, channel, dan interface.

## Tipe Data Number

Ada 2 jenis tipe data Number, yaitu :

- Integer
- Floating Point

Lebih detailnya versi tipe data int :

- int8 (-128 s/d 127)
- int16 (-32768 s/d 32767)
- int32 (-2147483648 s/d 2147483647)
- int64 (-9223372036854775808 s/d 9223372036854775807)
- uint8 (0 s/d 255)
- uint16 (0 s/d 65535)
- uint32 (0 s/d 4294967295)
- uint64 (0 s/d 18446744073709551615)

Lebih detailnya versi tipe data float :

- float32 (1.18x10^-38 s/d 3.4x10^38)
- float64 (2.23x10^-308 s/d 1.80x10^308)
- complex64
- complex128

Alias :

- byte (uint8)
- rune (int32)
- int (Minimal int32 {tergantung dari bit sistem operasi})
- uint (Minimal int32 {tergantung dari bit sistem operasi})
