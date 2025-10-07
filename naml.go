// Package naml stands for Not Another Markup Language, or nano-YAML.
//
// NAML is a subset of YAML that only requires a JSON parser and a tiny
// conversion func -- each line is a key-value pair (just like HTTP headers)
// where the value is encoded in JSON.
package naml

import (
	"bytes"
	"fmt"
	"strconv"
)

// Convert converts NAML data to JSON which you can then pass to
// json.Unmarshal
func Convert(data []byte) ([]byte, error) {
	return AppendConvert(nil, data)
}

// AppendConvert converts NAML data to JSON which you can then pass to
// json.Unmarshal. Appends the result to the given buffer.
func AppendConvert(buf []byte, data []byte) ([]byte, error) {
	const commaNewline = ",\n"
	if buf == nil {
		buf = make([]byte, 0, len(data)*2)
	}
	buf = append(buf, "{\n"...)

	var lno, keyCount int
	for line := range bytes.Lines(data) {
		lno++
		trimmed := bytes.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		}
		if trimmed[0] == '#' {
			continue
		}
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1] // chomp
		}
		if line[0] == ' ' || line[0] == '\t' {
			// continuation of the prior line
			if keyCount == 0 {
				return nil, &SyntaxError{lno, string(line), "first line cannot be a continuation line"}
			}
			buf = buf[:len(buf)-len(commaNewline)]
			buf = append(buf, '\n')
			buf = append(buf, line...)
			buf = append(buf, commaNewline...)
			continue
		}

		i := bytes.IndexByte(line, ':')
		if i < 0 {
			return nil, &SyntaxError{lno, string(line), "missing colon"}
		}
		key := bytes.TrimSpace(line[:i])
		value := bytes.TrimSpace(line[i+1:])
		if len(key) == 0 {
			return nil, &SyntaxError{lno, string(line), "empty key"}
		}
		if len(value) == 0 {
			return nil, &SyntaxError{lno, string(line), "empty value"}
		}

		buf = strconv.AppendQuoteToGraphic(buf, string(key))
		buf = append(buf, ": "...)
		buf = append(buf, value...)
		buf = append(buf, commaNewline...)
		keyCount++
	}

	buf = append(buf[:len(buf)-len(commaNewline)], "\n}"...)
	return buf, nil
}

type SyntaxError struct {
	LineNo int
	Line   string
	Reason string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("invalid NAML line %d (%s): %q", e.LineNo, e.Reason, e.Line)
}
