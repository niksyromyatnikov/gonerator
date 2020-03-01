package blocks

import (
	"fmt"
	"strconv"
)


/* String Type */

type String struct {
	string
	noQuotes bool
}

func NewString(val string) *String{
	return &String{string: val}
}

func (S String) GetName() string{
	return "String"
}

func (S String) GetType() string{
	return "string"
}

func (S String) Code(scopes ...ScopeInterface) string{
	if S.noQuotes {
		return S.string
	}

	return "\"" + S.string + "\""
}

func (S String) SetValue(val string) {
	S.string = val
}

func (S *String) NoQuotes() *String{
	S.noQuotes = true
	return S
}


/* Byte Type */

type Byte struct {
	byte
}

func NewByte(val byte) *Byte{
	return &Byte{val}
}

func (B Byte) GetName() string{
	return "byte"
}

func (B Byte) GetType() string{
	return "byte"
}

func (B Byte) Code(scopes ...ScopeInterface) string{
	return strconv.Itoa(int(B.byte))
}


/* Rune */

type Rune struct {
	rune
}

func NewRune(val rune) *Rune{
	return &Rune{val}
}

func (R Rune) GetName() string{
	return "rune"
}

func (R Rune) GetType() string{
	return "rune"
}

func (R Rune) Code(scopes ...ScopeInterface) string{
	return strconv.Itoa(int(R.rune))
}


/* Int */

type Int struct {
	int
}

func NewInt(val int) *Int{
	return &Int{val}
}

func (I Int) GetName() string{
	return "int"
}

func (I Int) GetType() string{
	return "int"
}

func (I Int) Code(scopes ...ScopeInterface) string{
	return strconv.Itoa(I.int)
}


/* Int8 */

type Int8 struct {
	int8
}

func NewInt8(val int8) *Int8{
	return &Int8{val}
}

func (I Int8) GetName() string{
	return "int8"
}

func (I Int8) GetType() string{
	return "int8"
}

func (I Int8) Code(scopes ...ScopeInterface) string{
	return strconv.Itoa(int(I.int8))
}


/* Int16 */

type Int16 struct {
	int16
}

func NewInt16(val int16) *Int16{
	return &Int16{val}
}

func (I Int16) GetName() string{
	return "int16"
}

func (I Int16) GetType() string{
	return "int16"
}

func (I Int16) Code(scopes ...ScopeInterface) string{
	return strconv.Itoa(int(I.int16))
}


/* Int32 */

type Int32 struct {
	int32
}

func NewInt32(val int32) *Int32{
	return &Int32{val}
}

func (I Int32) GetName() string{
	return "int32"
}

func (I Int32) GetType() string{
	return "int32"
}

func (I Int32) Code(scopes ...ScopeInterface) string{
	return strconv.Itoa(int(I.int32))
}


/* Int64 */

type Int64 struct {
	int64
}

func NewInt64(val int64) *Int64{
	return &Int64{val}
}

func (I Int64) GetName() string{
	return "int64"
}

func (I Int64) GetType() string{
	return "int64"
}

func (I Int64) Code(scopes ...ScopeInterface) string{
	return strconv.Itoa(int(I.int64))
}


/* Float32 */

type Float32 struct {
	float32
}

func NewFloat32(val float32) *Float32{
	return &Float32{val}
}

func (F Float32) GetName() string{
	return "float32"
}

func (F Float32) GetType() string{
	return "float32"
}

func (F Float32) Code(scopes ...ScopeInterface) string{
	return fmt.Sprintf("%f", F.float32)
}


/* Float64 */

type Float64 struct {
	float64
}

func NewFloat64(val float64) *Float64{
	return &Float64{val}
}

func (F Float64) GetName() string{
	return "float64"
}

func (F Float64) GetType() string{
	return "float64"
}

func (F Float64) Code(scopes ...ScopeInterface) string{
	return fmt.Sprintf("%f", F.float64)
}





