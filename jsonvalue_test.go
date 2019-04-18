/*
@Time : 2019-03-26 17:20
@Author : seefan
@File : jsonreader_test.go
@Software: jsonreader
*/
package jsonreader

import (
	"testing"
)

func TestJsonValue_GetValue(t *testing.T) {
	r := ParseJsonObject([]byte(" {\"key\" : 123,\"abc\":-1021e5 } "))
	v1 := r.GetValue("key")
	t.Log("v1=", v1)
	if v1.String() != "123" {
		t.Error("get value key != 123")
	}
	v2 := r.GetValue("abc")
	t.Log("v2=", v2)
	if v2 != "-1021e5" {
		t.Error("get value abc != -1021e5")
	}
}
func TestJsonValue_GetObjectAndValue(t *testing.T) {
	r := ParseJsonObject([]byte(" {\"data\" : {\"key\" : 123,\"abc\":-1021e5 }}"))
	v3 := r.GetObject("data").GetValue("key")
	t.Log("key=", v3)
	if v3 != "123" {
		t.Error("get value key!=123")
	}

}
func TestJsonValue_GetArray(t *testing.T) {
	arr := ParseJsonArray([]byte("[{\"data\" : {\"key\" : 123,\"abc\":-1021e5 } , \"value\":5 ,\"ars\":[1,2,3,4,{\"value\":5},\"6\",{\"value\":7}]},0,1,\"sfdada\"]"))
	arr.Each(func(i int, value JsonValue) {
		t.Log(i, value)
	})
	if arr.Size() != 4 {
		t.Error("array size error")
	}
}

func TestJsonValue_GetArrayObject(t *testing.T) {
	arr := ParseJsonArray([]byte("[{\"data\" : {\"key\" : 123,\"abc\":-1021e5 } , \"value\":5 ,\"ars\":[1,2,3,4,{\"value\":5},\"6\",{\"value\":7}]},0,1,\"sfdada\"]"))
	arr.Each(func(i int, value JsonValue) {
		if i == 0 {
			key := value.ParseJsonObject().GetObject("data").GetValue("key")
			if key != "123" {
				t.Error("get key value!=123")
			}
		}
		t.Log(i, value)
	})
}
func TestJsonValue_GetArrayString(t *testing.T) {
	arr := ParseJsonArray([]byte("[0,1,\"sfd\\\"ada\",{\"test\":\"ddd\"}]"))
	arr.Each(func(i int, value JsonValue) {
		t.Log(i, value)
	})
}
func TestJsonValue_Number(t *testing.T) {
	obj := ParseJsonObject([]byte("{\"key\":-1021e5}"))
	v := obj.GetValue("key").Float64()
	if v != -1.021e+08 {
		t.Error("number parse error")
	}
}
func TestJsonValue_Escape(t *testing.T) {
	obj := ParseJsonObject([]byte("{\"key\":\"\t\n\u6211\u559c\u6b22\u4f60\u554a\"}"))
	v := obj.GetValue("key")
	if v != "\t\n我喜欢你啊" {
		t.Error("escape parse error")
	}
}
func TestJsonValue_Bool(t *testing.T) {
	obj := ParseJsonObject([]byte("{\"key\":true,\"t\":false}"))
	v := obj.GetValue("key").Bool()
	if !v {
		t.Error(" parse error")
	}
}
func TestJsonValue_Null(t *testing.T) {
	obj := ParseJsonObject([]byte("{\"key\":null,\"t\":false}"))
	v := obj.GetValue("key")
	if v.IsNull() {
		t.Error(" parse error")
	}
}

func TestJsonValue_C(t *testing.T) {
	obj := ParseJsonObject([]byte("{\"contacts\":{ \"company\":{ \"address\":null, \"state_code\":null }, \"employees\":[ null] }}"))
	v := obj.V("contacts").ParseJsonObject().A("employees")
	if v.Size() == 0 {
		t.Error(" parse error")
	} else {
		t.Log(v.Get(0).String(), v.Get(0).IsNull())
	}
}
