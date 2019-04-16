/*
@Time : 2019-03-27 18:27
@Author : seefan
@File : jsonarr
@Software: jsonreader
*/
package jsonreader

type jsonArray struct {
	reader
	arr []JsonValue
}

func (j *jsonArray) Get(i int) JsonValue {
	return j.arr[i]
}
func (j *jsonArray) Size() int {
	return len(j.arr)
}
func ParseJsonArray(bs []byte) *jsonArray {
	j := &jsonArray{
		reader: *newReader(bs),
	}
	j.parse()
	return j
}
func (j *jsonArray) Each(f func(int, JsonValue)) {
	if j.arr != nil {
		for i, v := range j.arr {
			f(i, v)
		}
	}
}
func (j *jsonArray) parse() {
	if !j.validArray() {
		j.LastError = "JsonArray format error"
		j.end = -1
		return
	}
	//remove []
	j.index++
	j.end--
	start := j.index
	str := 0
	depth := 0
	for j.index < j.end {
		switch j.data[j.index] {
		case '[', '{':
			depth++
		case ']', '}':
			depth--
		case '"':
			str++
		case '\\':
			if j.index+1 < j.end && j.data[j.index+1] == '"' {
				j.index++
			}
		case ',':
			if depth == 0 && str%2 == 0 {
				if j.data[start] == '"' {
					j.arr = append(j.arr, JsonValue(j.data[start+1:j.index-1]))
				} else {
					j.arr = append(j.arr, JsonValue(j.data[start:j.index]))
				}
				start = j.index + 1
			}
		}
		j.index++
	}
	if start < j.index {
		if j.data[start] == '"' {
			j.arr = append(j.arr, JsonValue(j.data[start+1:j.index-1]))
		} else {
			j.arr = append(j.arr, JsonValue(j.data[start:j.index]))
		}
	}
}
