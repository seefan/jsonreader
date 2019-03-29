/*
@Time : 2019-03-27 12:36
@Author : seefan
@File : JsonValue
@Software: jsonreader
*/
package jsonreader

type reader struct {
	data  []byte
	start int
	end   int
}

func newReader(data []byte) *reader {
	r := &reader{
		data: data,
		end:  len(data) - 1,
	}
	return r
}

func (r *reader) IsValid() bool {
	return r.end >= r.start
}

func (r *reader) hasMore() bool {
	for r.start <= r.end {
		if r.data[r.start] == ',' {
			return true
		}
		r.start++
	}
	return false
}
func (r *reader) parseNumber() []byte {
	start := r.start
	for r.start < r.end {
		if r.data[r.start] <= ' ' || r.data[r.start] == ',' || r.data[r.start] == ']' || r.data[r.start] == '}' {
			break
		}
		r.start++
	}
	if r.IsValid() {
		return r.data[start : r.start+1]
	}
	r.end = -1
	return nil
}
func (r *reader) parseString() []byte {
	start := -1
	for r.start <= r.end {
		if start == -1 {
			r.skip()
		}
		if r.data[r.start] == '"' {
			if start == -1 {
				start = r.start + 1
			} else {
				return r.data[start:r.start]
			}
		} else if r.data[r.start] == '\\' {
			r.start++
		}
		r.start++
	}
	r.end = -1
	return nil
}
func (r *reader) validObject() bool {
	r.skip()
	r.skip(true)
	if !r.IsValid() {
		return false
	}
	return r.data[r.start] == '{' && r.data[r.end] == '}'
}
func (r *reader) validArray() bool {
	r.skip()
	r.skip(true)
	if !r.IsValid() {
		return false
	}
	return r.data[r.start] == '[' && r.data[r.end] == ']'
}
func (r *reader) skipSplit() {
	for r.start <= r.end {
		r.start++
		if r.data[r.start] == ':' {
			r.start++
			break
		}
	}
}
func (r *reader) parseValue() []byte {
	depth := 0
	start := r.start
	for r.start <= r.end {
		switch r.data[r.start] {
		case '"': //string
			if depth == 0 {
				return r.parseString()
			}
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': //number
			if depth == 0 {
				return r.parseNumber()
			}
		case ' ', '\t', '\r', '\n': //space
		case '{', '[': //object array begin
			depth++
		case '}', ']': //object array end
			depth--
			if depth == 0 {
				return r.data[start : r.start+1]
			}
		}
		r.start++
	}
	return nil
}
func (r *reader) skip(reverse ...bool) {
	if len(reverse) == 0 || reverse[0] == false {
		for i := r.end; i > r.start; i-- {
			if !isSpace(r.data[i]) {
				r.end = i
				break
			}
		}
	}
	if len(reverse) > 0 && reverse[0] {
		for i := r.start; i <= r.end; i++ {
			if !isSpace(r.data[i]) {
				r.start = i
				break
			}
		}
	}
}
func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}
