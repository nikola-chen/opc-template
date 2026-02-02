package schema

import (
	"io"
	"os"
)

func ApplyPatch(schemaPath string, patch []PatchOp) error {
	if err := backup(schemaPath); err != nil {
		// log error but maybe proceed? or fail? fail is safer.
		return err
	}

	// apply RFC6902 patch
	// v0.1: 使用现成 jsonpatch lib
	// v0.2: 自研约束引擎

	return nil
}

func backup(src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(src + ".prev.json")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
