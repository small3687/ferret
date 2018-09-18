package values

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/pkg/errors"
	"strings"
)

type String string

var EmptyString = String("")
var SpaceString = String(" ")

func NewString(input string) String {
	if input == "" {
		return EmptyString
	}

	return String(input)
}

func ParseString(input interface{}) (String, error) {
	if core.IsNil(input) {
		return EmptyString, nil
	}

	str, ok := input.(string)

	if ok {
		if str != "" {
			return String(str), nil
		}

		return EmptyString, nil
	}

	stringer, ok := input.(fmt.Stringer)

	if ok {
		return String(stringer.String()), nil
	}

	return EmptyString, errors.Wrap(core.ErrInvalidType, "expected 'string'")
}

func ParseStringP(input interface{}) String {
	res, err := ParseString(input)

	if err != nil {
		panic(err)
	}

	return res
}

func (t String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t String) Type() core.Type {
	return core.StringType
}

func (t String) String() string {
	return string(t)
}

func (t String) Compare(other core.Value) int {
	switch other.Type() {
	case core.StringType:
		return strings.Compare(string(t), other.Unwrap().(string))
	default:
		if other.Type() > core.DateTimeType {
			return -1
		}

		return 1
	}
}

func (t String) Unwrap() interface{} {
	return string(t)
}

func (t String) Hash() int {
	h := sha512.New()

	out, err := h.Write([]byte(t))

	if err != nil {
		return 0
	}

	return out
}

func (t String) Length() Int {
	return Int(len(t))
}

func (t String) Contains(other String) Boolean {
	return t.IndexOf(other) > -1
}

func (t String) IndexOf(other String) Int {
	return Int(strings.Index(string(t), string(other)))
}

func (t String) Concat(other core.Value) String {
	return String(string(t) + other.String())
}