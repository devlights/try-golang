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
}
