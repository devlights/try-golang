package interface_

import (
	"github.com/devlights/try-golang/basic/interface_/command"
)

// VerifyInterfaceCompliance は、インターフェースの実装を検証するやり方のサンプルです.
//
// REFERENCES::
//   - https://qiita.com/kskumgk63/items/423df2e5245da4b16c25
//   - https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
func VerifyInterfaceCompliance() error {

	cmd := &command.ListFileCommand{
		Dir:     ".",
		Pattern: "go.*",
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
