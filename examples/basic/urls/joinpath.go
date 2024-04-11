package urls

import (
	"net/url"

	"github.com/devlights/gomy/errs"
	"github.com/devlights/gomy/output"
)

// JoinPath -- Go1.19 から追加された url.JoinPath() についてのサンプルです.
func JoinPath() error {
	p, err := url.JoinPath("base", "child1", ".", "child2", "..", "child3")
	if err != nil {
		return err
	}

	output.Stdoutl("[url.JoinPath]", p)

	var (
		u  = errs.Drop(url.Parse("https://devlights.hatenablog.com/"))
		p2 = errs.Drop(url.JoinPath(u.String(), "entry", "2022", "08", "24", ".", "073000"))
	)

	output.Stdoutl("[url.JoinPath]", p2)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: url_joinpath

	   [Name] "url_joinpath"
	   [url.JoinPath]       base/child1/child3
	   [url.JoinPath]       https://devlights.hatenablog.com/entry/2022/08/24/073000


	   [Elapsed] 40.02µs
	*/

}
