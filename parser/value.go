package parser

// #include "query_types.h"
// #include <stdlib.h>
import "C"

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type ValueType int

const (
	ValueRegex        ValueType = C.VALUE_REGEX
	ValueInt                    = C.VALUE_INT
	ValueBool                   = C.VALUE_BOOLEAN
	ValueFloat                  = C.VALUE_FLOAT
	ValueString                 = C.VALUE_STRING
	ValueIntoName               = C.VALUE_INTO_NAME
	ValueTableName              = C.VALUE_TABLE_NAME
	ValueSimpleName             = C.VALUE_SIMPLE_NAME
	ValueDuration               = C.VALUE_DURATION
	ValueWildcard               = C.VALUE_WILDCARD
	ValueFunctionCall           = C.VALUE_FUNCTION_CALL
	ValueExpression             = C.VALUE_EXPRESSION
)

type Value struct {
	Name          string
	Alias         string
	Type          ValueType
	Elems         []*Value
	compiledRegex *regexp.Regexp
	IsInsensitive bool
}

func (self *Value) IsFunctionCall() bool {
	return self.Type == ValueFunctionCall
}

func (self *Value) GetCompiledRegex() (*regexp.Regexp, bool) {
	return self.compiledRegex, self.Type == ValueRegex
}

func (self *Value) GetString() string {
	buffer := bytes.NewBufferString("")
	switch self.Type {
	case ValueExpression:
		fmt.Fprintf(buffer, "%s %s %s", self.Elems[0].GetString(), self.Name, self.Elems[1].GetString())
	case ValueFunctionCall:
		fmt.Fprintf(buffer, "%s(%s)", self.Name, Values(self.Elems).GetString())
	case ValueString:
		fmt.Fprintf(buffer, "'%s'", self.Name)
	case ValueRegex:
		fmt.Fprintf(buffer, "/%s/", QueryFormatRegex(self.Name))
		if self.IsInsensitive {
			buffer.WriteString("i")
		}
	case ValueSimpleName:
		buffer.WriteString(QueryFormatSimpleName(self.Name))
	default:
		buffer.WriteString(self.Name)
	}

	if self.Alias != "" {
		fmt.Fprintf(buffer, " as %s", QueryFormatSimpleName(self.Alias))
	}

	return buffer.String()
}

func StringNeedsQuotes(r rune) bool {
	return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_');
}

func QueryFormatSimpleName(s string) string {
	needsQuotes := strings.IndexFunc(s, StringNeedsQuotes)

	// all characters are ok
	if (needsQuotes == -1) {
		return s;
	}

	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	return fmt.Sprintf("\"%s\"", s)
}

func QueryFormatRegex(s string) string {
	if (strings.Index(s, "/") == -1) {
		return s;
	}

	return strings.Replace(s, "/", "\\/", -1)
}
