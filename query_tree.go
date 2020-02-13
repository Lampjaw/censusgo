// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import "fmt"

type queryTree struct {
	censusOperator
	tree      []*queryTree
	treeField string
	List      bool   `queryProp:"list"`
	Prefix    string `queryProp:"prefix"`
	Start     string `queryProp:"start"`
}

func newQueryTree(field string) *queryTree {
	return &queryTree{
		List:      false,
		Prefix:    "",
		Start:     "",
		tree:      make([]*queryTree, 0),
		treeField: field,
	}
}

func (t *queryTree) IsList(isList bool) *queryTree {
	t.List = isList
	return t
}

func (t *queryTree) GroupPrefix(prefix string) *queryTree {
	t.Prefix = prefix
	return t
}

func (t *queryTree) StartField(field string) *queryTree {
	t.Start = field
	return t
}

func (t *queryTree) TreeField(field string) *queryTree {
	newTree := newQueryTree(field)
	t.tree = append(t.tree, newTree)
	return newTree
}

func (t *queryTree) String() string {
	baseString := t.baseString(t)

	if len(baseString) > 0 {
		baseString = "^" + baseString
	}

	subTreeString := ""
	for _, subTree := range t.tree {
		subTreeString += fmt.Sprintf("(%s)", subTree.String())
	}

	return t.treeField + baseString + subTreeString
}

func (t *queryTree) getKeyValueStringFormat() string {
	return "%s:%s"
}

func (t *queryTree) getPropertySpacer() string {
	return "^"
}

func (t *queryTree) getTermSpacer() string {
	return "'"
}
