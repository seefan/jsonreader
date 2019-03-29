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
	r := ParseJsonObject([]byte(" {\"data\" : {\"key\" : 123,\"abc\":-1021e5 } , \"value\":5 ,\"ars\":[1,2,3,4,{\"value\":5},\"6\",{\"value\":7}]}"))
	t.Log(r.GetJsonObject("data").GetJsonValue("abc").String())

	r.GetJsonArray("ars").each(func(i int, value JsonValue) {
		t.Log(i, value)
	})
	arr := ParseJsonArray([]byte("[{\"data\" : {\"key\" : 123,\"abc\":-1021e5 } , \"value\":5 ,\"ars\":[1,2,3,4,{\"value\":5},\"6\",{\"value\":7}]},0,1,\"sfdada\"]"))
	arr.each(func(i int, value JsonValue) {
		if i == 0 {
			t.Log(value.JsonObject().GetJsonObject("data").GetJsonValue("key"))
		}
		t.Log(i, value)
	})
}
