// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import "fmt"

type queryTree struct {
	censusOperator
	list      bool   `queryProp:"list"`
	prefix    string `queryProp:"prefix"`
	start     string `queryProp:"start"`
	tree      []*queryTree
	treeField string
}

func newQueryTree(field string) *queryTree {
	return &queryTree{
		list:      false,
		prefix:    "",
		start:     "",
		tree:      make([]*queryTree, 0),
		treeField: field,
	}
}

func (t *queryTree) IsList(isList bool) {
	t.list = isList
}

func (t *queryTree) GroupPrefix(prefix string) {
	t.prefix = prefix
}

func (t *queryTree) StartField(field string) {
	t.start = field
}

func (t *queryTree) TreeField(field string) *queryTree {
	newTree := newQueryTree(field)
	t.tree = append(t.tree, newTree)
	return newTree
}

func (t *queryTree) String() string {
	baseString := t.BaseString()

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
