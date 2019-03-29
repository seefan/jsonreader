/*
@Time : 2019-03-27 18:27
@Author : seefan
@File : jsonarr
@Software: jsonreader
*/
package jsonreader

type JsonArray struct {
	reader
	arr []JsonValue
}

func (j *JsonArray) Get(i int) JsonValue {
	return j.arr[i]
}
func (j *JsonArray) Size() int {
	return len(j.arr)
}
func ParseJsonArray(bs []byte) *JsonArray {
	j := &JsonArray{
		reader: *newReader(bs),
	}
	j.parse()
	return j
}
func (j *JsonArray) each(f func(int, JsonValue)) {
	if j.arr != nil {
		for i, v := range j.arr {
			f(i, v)
		}
	}
}
func (j *JsonArray) parse() {
	if !j.validArray() {
		j.end = -1
		return
	}
	//remove []
	j.start++
	j.end--
	start := j.start
	str := false
	depth := 0
	for j.start <= j.end {
		switch j.data[j.start] {
		case '[', '{':
			depth++
		case ']', '}':
			depth--
		case '"':
			str = !str
		case ',':
			if depth == 0 {
				j.arr = append(j.arr, JsonValue(j.data[start:j.start]))
				start = j.start + 1
			}
		}
		j.start++
	}
	if start < j.start {
		j.arr = append(j.arr, JsonValue(j.data[start:j.start]))
	}
}
