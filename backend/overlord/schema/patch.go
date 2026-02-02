package schema

import (
	"errors"
	"strings"
)

type PatchOp struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

func ValidatePatch(patch []PatchOp) error {
	for _, op := range patch {
		if !strings.HasPrefix(op.Path, "/") {
			return errors.New("invalid patch path")
		}
	}
	return nil
}
