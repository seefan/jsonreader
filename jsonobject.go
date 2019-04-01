/*
@Time : 2019-03-27 13:54
@Author : seefan
@File : jsonObject
@Software: jsonreader
*/
package jsonreader

type jsonObject struct {
	reader
	val map[string]JsonValue
	obj map[string]*jsonObject
	arr map[string]*jsonArray
}

func (j *jsonObject) parse() {
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

func ParseJsonObject(bs []byte) *jsonObject {
	j := &jsonObject{
		reader: *newReader(bs),
		val:    make(map[string]JsonValue),
	}
	j.parse()
	return j
}
func (j *jsonObject) Contains(name string) bool {
	if _, ok := j.val[name]; ok {
		return true
	}
	return false
}
func (j *jsonObject) GetValue(name string) JsonValue {
	return j.val[name]
}
func (j *jsonObject) GetObject(name string) *jsonObject {
	if j.obj != nil {
		if v, ok := j.obj[name]; ok {
			return v
		}
	}
	return j.GetObjectForce(name)
}
func (j *jsonObject) GetObjectForce(name string) *jsonObject {
	if v, ok := j.val[name]; ok {
		if j.obj == nil {
			j.obj = make(map[string]*jsonObject)
		}
		j.obj[name] = ParseJsonObject(v.Bytes())
		return j.obj[name]
	}
	return &jsonObject{
		val: make(map[string]JsonValue),
	}
}
func (j *jsonObject) GetArray(name string) *jsonArray {
	if j.arr != nil {
		if v, ok := j.arr[name]; ok {
			return v
		}
	}
	return j.GetArrayForce(name)
}
func (j *jsonObject) GetArrayForce(name string) *jsonArray {
	if v, ok := j.val[name]; ok {
		if j.arr == nil {
			j.arr = make(map[string]*jsonArray)
		}
		j.arr[name] = ParseJsonArray(v.Bytes())
		return j.arr[name]
	}
	return &jsonArray{}
}
