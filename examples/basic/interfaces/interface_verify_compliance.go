package interfaces

import (
	"github.com/devlights/try-golang/examples/basic/interfaces/command"
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: interface_verify_compliance

	   [Name] "interface_verify_compliance"
	   go.mod
	   go.sum


	   [Elapsed] 165.59µs
	*/

}
