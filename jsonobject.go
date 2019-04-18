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
	j.index++
	j.end--
	for j.index <= j.end {
		j.skip()
		if !j.IsValid() {
			break
		}
		key := unescape(j.parseString())
		j.skipSplit()
		if !j.IsValid() {
			break
		}
		value := j.parseValue()
		j.val[key] = JsonValue(value)
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
func (j *JsonObject) C(name string) bool {
	return j.Contains(name)
}
func (j *JsonObject) Contains(name string) bool {
	if _, ok := j.val[name]; ok {
		return true
	}
	return false
}
func (j *JsonObject) V(name string) JsonValue {
	return j.GetValue(name)
}
func (j *JsonObject) GetValue(name string) JsonValue {
	return j.val[name]
}
func (j *JsonObject) O(name string) *JsonObject {
	return j.GetObject(name)
}
func (j *JsonObject) GetObject(name string) *JsonObject {
	if j.obj != nil {
		if v, ok := j.obj[name]; ok {
			return v
		}
	}
	return j.GetObjectForce(name)
}
func (j *JsonObject) GetObjectForce(name string) *JsonObject {
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
func (j *JsonObject) A(name string) *JsonArray {
	return j.GetArray(name)
}
func (j *JsonObject) GetArray(name string) *JsonArray {
	if j.arr != nil {
		if v, ok := j.arr[name]; ok {
			return v
		}
	}
	return j.GetArrayForce(name)
}
func (j *JsonObject) GetArrayForce(name string) *JsonArray {
	if v, ok := j.val[name]; ok {
		if j.arr == nil {
			j.arr = make(map[string]*JsonArray)
		}
		j.arr[name] = ParseJsonArray(v.Bytes())
		return j.arr[name]
	}
	return &JsonArray{}
}
