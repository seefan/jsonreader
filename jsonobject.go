/*
@Time : 2019-03-27 13:54
@Author : seefan
@File : JsonObject
@Software: jsonreader
*/
package jsonreader

type JsonObject struct {
	reader
	val map[string]JsonValue
	obj map[string]*JsonObject
	arr map[string]*JsonArray
}

func (j *JsonObject) parse() {
	if !j.validObject() {
		j.end = -1
		return
	}

	//remove {}
	j.start++
	j.end--
	for j.start <= j.end {
		j.skip()
		if !j.IsValid() {
			break
		}
		key := j.parseString()
		j.skipSplit()
		if !j.IsValid() {
			break
		}
		value := j.parseValue()
		j.val[string(key)] = JsonValue(value)
		if !j.hasMore() {
			break
		}
	}
}

func ParseJsonObject(bs []byte) *JsonObject {
	j := &JsonObject{
		reader: *newReader(bs),
		val:    make(map[string]JsonValue),
	}
	j.parse()
	return j
}
func (j *JsonObject) GetJsonValue(name string) JsonValue {
	if v, ok := j.val[name]; ok {
		return v
	}
	return "name not exists"
}
func (j *JsonObject) GetJsonObject(name string) *JsonObject {
	if j.obj != nil {
		if v, ok := j.obj[name]; ok {
			return v
		}
	}
	return j.GetJsonObjectForce(name)
}
func (j *JsonObject) GetJsonObjectForce(name string) *JsonObject {
	if v, ok := j.val[name]; ok {
		if j.obj == nil {
			j.obj = make(map[string]*JsonObject)
		}
		j.obj[name] = ParseJsonObject(v.Bytes())
		return j.obj[name]
	}
	return &JsonObject{
		val: make(map[string]JsonValue),
	}
}
func (j *JsonObject) GetJsonArray(name string) *JsonArray {
	if j.arr != nil {
		if v, ok := j.arr[name]; ok {
			return v
		}
	}
	return j.GetJsonArrayForce(name)
}
func (j *JsonObject) GetJsonArrayForce(name string) *JsonArray {
	if v, ok := j.val[name]; ok {
		if j.arr == nil {
			j.arr = make(map[string]*JsonArray)
		}
		j.arr[name] = ParseJsonArray(v.Bytes())
		return j.arr[name]
	}
	return &JsonArray{}
}
