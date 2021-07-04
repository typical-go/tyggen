package typgen

import (
	"fmt"
	"strings"
)

type (
	// Signature for code generation by annotation
	Signature struct {
		TagName string
	}
)

var _ fmt.Stringer = (*Signature)(nil)

func (s Signature) String() string {
	var out strings.Builder
	fmt.Fprint(&out, "DO NOT EDIT. ")
	if s.TagName != "" {
		fmt.Fprintf(&out, "This file generated due to '%s' annotation", s.TagName)
	} else {
		fmt.Fprint(&out, "Autogenerated by Typical-Go")
	}

	return out.String()
}