package p06

import "fmt"

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%d", t.Val)
	}

	var str string
	str = fmt.Sprintf("%s%d", str, t.Val)
	str = fmt.Sprintf("%s(%s)", str, tree2str(t.Left))
	if t.Right != nil {
		str = fmt.Sprintf("%s(%s)", str, tree2str(t.Right))
	}

	return str
}
