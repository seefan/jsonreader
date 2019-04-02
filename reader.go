/*
@Time : 2019-03-27 12:36
@Author : seefan
@File : JsonValue
@Software: jsonreader
*/
package jsonreader

type reader struct {
	data      []byte
	index     int
	end       int
	LastError string
}

func newReader(data []byte) *reader {
	r := &reader{
		data: data,
		end:  len(data),
	}
	return r
}

func (r *reader) IsValid() bool {
	return r.end != -1
}

func (r *reader) hasMore() bool {
	for r.index < r.end {
		if r.data[r.index] == ',' {
			return true
		}
		r.index++
	}
	return false
}
func (r *reader) parseNumber() []byte {
	start := r.index
	for r.index < r.end {
		if r.data[r.index] <= ' ' || r.data[r.index] == ',' || r.data[r.index] == ']' || r.data[r.index] == '}' {
			break
		}
		r.index++
	}
	if r.IsValid() {
		return r.data[start:r.index]
	}
	r.LastError = "Number format error"
	r.end = -1
	return nil
}
func (r *reader) parseString() []byte {
	start := -1
	for r.index < r.end {
		if start == -1 {
			r.skip()
		}
		if r.data[r.index] == '"' {
			if start == -1 {
				start = r.index + 1
			} else {
				return r.data[start:r.index]
			}
		} else if r.data[r.index] == '\\' {
			if r.index+1 <= r.end && r.data[r.index+1] == '"' {
				r.index++
			}
		}
		r.index++
	}
	r.LastError = "String format error"
	r.end = -1
	return nil
}
func unescape(json []byte) string {
	var str = make([]byte, 0, len(json))
	for i := 0; i < len(json); i++ {
		switch {
		default:
			str = append(str, json[i])
		case json[i] < ' ': //skip
		case json[i] > '~': //skip
		case json[i] == '\\':
			i++
			if i >= len(json) {
				return string(str)
			}
			switch json[i] {
			default:
				return string(str)
			case '\\':
				str = append(str, '\\')
			case '/':
				str = append(str, '/')
			case 'b':
				str = append(str, '\b')
			case 'f':
				str = append(str, '\f')
			case 'n':
				str = append(str, '\n')
			case 'r':
				str = append(str, '\r')
			case 't':
				str = append(str, '\t')
			case '"':
				str = append(str, '"')
			}
		}
	}
	return string(str)
}
func (r *reader) validObject() bool {
	r.trim()
	if r.index < r.end {
		return r.data[r.index] == '{' && r.data[r.end-1] == '}'
	}
	return false
}
func (r *reader) validArray() bool {
	r.trim()
	if r.index < r.end {
		return r.data[r.index] == '[' && r.data[r.end-1] == ']'
	}
	return false
}
func (r *reader) skipSplit() {
	for r.index < r.end {
		r.index++
		if r.data[r.index] == ':' {
			r.index++
			break
		}
	}
}
func (r *reader) parseValue() []byte {
	depth := 0
	start := r.index
	for r.index < r.end {
		switch r.data[r.index] {
		case '"': //string
			if depth == 0 {
				return r.parseString()
			}
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': //number
			if depth == 0 {
				return r.parseNumber()
			}
		case '{', '[': //object array begin
			depth++
		case '}', ']': //object array end
			depth--
			if depth == 0 {
				return r.data[start : r.index+1]
			}
		}
		r.index++
	}
	return nil
}
func (r *reader) skip() {
	for i := r.index; i < r.end; i++ {
		if !isSpace(r.data[i]) {
			r.index = i
			break
		}
	}
}
func (r *reader) trim() {
	for i := r.end - 1; i > r.index; i-- {
		if !isSpace(r.data[i]) {
			r.end = i + 1
			break
		}
	}
	for i := r.index; i < r.end; i++ {
		if !isSpace(r.data[i]) {
			r.index = i
			break
		}
	}
}
func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}
