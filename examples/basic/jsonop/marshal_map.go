package jsonop

import (
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// MarshalMap は、json.Marshal() で マップ をマーシャルした場合のサンプルです.
func MarshalMap() error {
	var (
		items = map[string]string{
			"golang": "fmt.Println",
			"java":   "System.out.println",
			"dotnet": "Console.WriteLine",
			"python": "print",
		}
	)

	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(&items); err != nil {
		return err
	}

	output.Stdoutl("[marshal]", string(buf))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_marshal_map

	   [Name] "json_marshal_map"
	   [marshal]            {"dotnet":"Console.WriteLine","golang":"fmt.Println","java":"System.out.println","python":"print"}


	   [Elapsed] 64.96µs
	*/

}
