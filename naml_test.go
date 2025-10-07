package naml

import (
	"encoding/json"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    `title: "Hello"`,
			expected: "{\n\"title\": \"Hello\"\n}",
		},
		{
			input:    "title: \"Hello\"\npermalink: \"/test/\"",
			expected: "{\n\"title\": \"Hello\",\n\"permalink\": \"/test/\"\n}",
		},
		{
			input:    `cta: {"primary": {"title": "Click me"}}`,
			expected: "{\n\"cta\": {\"primary\": {\"title\": \"Click me\"}}\n}",
		},
		{
			input:    "title: \"Hello\"\npermalink: \"/test/\"\ncta: {\n  \"primary\": {\n  \"title\": \"Click me\"\n  }\n }",
			expected: "{\n\"title\": \"Hello\",\n\"permalink\": \"/test/\",\n\"cta\": {\n  \"primary\": {\n  \"title\": \"Click me\"\n  }\n }\n}",
		},
		{
			input:    `count: 42`,
			expected: "{\n\"count\": 42\n}",
		},
		{
			input:    `active: true`,
			expected: "{\n\"active\": true\n}",
		},
		{
			input:    `value: null`,
			expected: "{\n\"value\": null\n}",
		},
		{
			input:    "title: 11\n\npermalink: 22",
			expected: "{\n\"title\": 11,\n\"permalink\": 22\n}",
		},
		{
			input:    "# This is a comment\ntitle: 11\n# Another comment\npermalink: 22",
			expected: "{\n\"title\": 11,\n\"permalink\": 22\n}",
		},
		{
			input:    "topics: [11,\n  22]",
			expected: "{\n\"topics\": [11,\n  22]\n}",
		},
		{
			input:    "topics: [11,\n\t22]",
			expected: "{\n\"topics\": [11,\n\t22]\n}",
		},
		{
			input:    `title  :  "Hello"  `,
			expected: "{\n\"title\": \"Hello\"\n}",
		},
		{
			input:    `my key: "value"`,
			expected: "{\n\"my key\": \"value\"\n}",
		},
		{
			input:    `invalid line`,
			expected: `ERR: invalid NAML line 1 (missing colon): "invalid line"`,
		},
		{
			input:    `: "value"`,
			expected: `ERR: invalid NAML line 1 (empty key): ": \"value\""`,
		},
		{
			input:    `key:`,
			expected: `ERR: invalid NAML line 1 (empty value): "key:"`,
		},
		{
			input:    `key:   `,
			expected: `ERR: invalid NAML line 1 (empty value): "key:   "`,
		},
		{
			input:    `  key  :  "Hello"  `,
			expected: `ERR: invalid NAML line 1 (first line cannot be a continuation line): "  key  :  \"Hello\"  "`,
		},
	}

	for _, tt := range tests {
		result, err := Convert([]byte(tt.input))

		var got string
		if err != nil {
			got = "ERR: " + err.Error()
		} else {
			got = string(result)
		}

		if got != tt.expected {
			t.Errorf("** Convert(%q) = %q, want %q", tt.input, got, tt.expected)
		} else if err == nil {
			var dummy any
			err = json.Unmarshal(result, &dummy)
			if err != nil {
				t.Errorf("** json.Unmarshal(Convert(%q)) fails: %v", tt.input, err)
			}
		}
	}
}
