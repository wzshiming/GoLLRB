package llrb

import (
	"bytes"
	"fmt"
	"io"
)

func (t *LLRB) String() string {
	buf := bytes.NewBuffer(nil)
	fprintNodes(buf, t.root, nil, edgeRoot)
	return buf.String()
}

func fprintNodes(wr io.Writer, n *Node, pre edges, edge edge) {
	if n == nil {
		return
	}

	{
		es := edgeSpace
		if edge == edgeRight {
			es = edgeLink
		}
		fprintNodes(wr, n.Left, append(pre, es), edgeLeft)
	}

	fmt.Fprintf(wr, "%s%v\n", append(pre, edge), n.Item)

	{
		es := edgeSpace
		if edge == edgeLeft {
			es = edgeLink
		}
		fprintNodes(wr, n.Right, append(pre, es), edgeRight)
	}
}

type edges []edge

func (e edges) String() string {
	buf := bytes.NewBuffer(nil)
	for _, v := range e {
		buf.WriteString(v.String())
	}
	return buf.String()
}

type edge uint

const (
	_ = edge(iota)
	edgeSpace
	edgeLink
	edgeRoot
	edgeLeft
	edgeRight
)

func (e edge) String() string {
	switch e {
	case edgeSpace:
		return `　`
	case edgeLink:
		return `│`
	case edgeRoot:
		return `─>`
	case edgeLeft:
		return `┌>`
	case edgeRight:
		return `└>`
	default:
		return ``
	}
}
