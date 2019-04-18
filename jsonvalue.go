/*
@Time : 2019-03-26 16:36
@Author : seefan
@File : jsonreader
@Software: jsonreader
*/
package jsonreader

import (
	"strconv"
	"time"
)

//扩展值，原始类型为 string
type JsonValue string

//返回 string 的值
func (v JsonValue) String() string {
	return string(v)
}

//返回 int64 的值
func (v JsonValue) Int64() int64 {
	if f, e := strconv.ParseInt(string(v), 10, 64); e == nil {
		return f
	}
	return 0
}

//返回 int32 的值
func (v JsonValue) Int32() int32 {
	return int32(v.Int64())
}

//返回 int16 的值
func (v JsonValue) Int16() int16 {
	return int16(v.Int64())
}

//返回 int8 的值
func (v JsonValue) Int8() int8 {
	return int8(v.Int64())
}

//返回 int 的值
func (v JsonValue) Int() int {
	return int(v.Int64())
}

//返回 uint64 的值
func (v JsonValue) UInt64() uint64 {
	if f, e := strconv.ParseUint(string(v), 10, 64); e == nil {
		return f
	}
	return 0
}

//返回 uint32 类型的值
func (v JsonValue) UInt32() uint32 {
	return uint32(v.UInt64())
}

//返回 uint16 类型的值
func (v JsonValue) UInt16() uint16 {
	return uint16(v.UInt64())
}

//返回 uint8 类型的值
func (v JsonValue) UInt8() uint8 {
	return uint8(v.UInt64())
}

//返回 byte 类型的值
func (v JsonValue) Byte() byte {
	return v.UInt8()
}

//返回 uint 类型的值
func (v JsonValue) UInt() uint {
	return uint(v.UInt64())
}

//返回 float64 类型的值
func (v JsonValue) Float64() float64 {
	if f, e := strconv.ParseFloat(string(v), 64); e == nil {
		return f
	}
	return 0
}

//返回 float32 类型的值
func (v JsonValue) Float32() float32 {
	return float32(v.Float64())
}

//返回 bool 类型的值
func (v JsonValue) Bool() bool {
	return v == "true"
}

//return is null
// v=='null'
func (v JsonValue) IsNull() bool {
	return v == "null"
}

//返回 time.Time 类型的值
func (v JsonValue) Time() time.Time {
	return time.Unix(v.Int64(), 0)
}

//返回 time.Duration 类型的值
func (v JsonValue) Duration() time.Duration {
	return time.Duration(v.Int64())
}

//返回 []byte 类型的值
func (v JsonValue) Bytes() []byte {
	return []byte(v)
}

//判断是否为空
func (v JsonValue) IsEmpty() bool {
	return v == "" || v.IsNull()
}
func (v JsonValue) ParseJsonObject() *jsonObject {
	return ParseJsonObject([]byte(v))
}
func (v JsonValue) ParseJsonArray() *jsonArray {
	return ParseJsonArray([]byte(v))
}
