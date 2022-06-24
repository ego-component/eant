package eant

// AntTree https://ant.design/components/tree-cn/#TreeNode-props
type AntTree struct {
	originOptions   AntTreeOptions
	deepCopyOptions AntTreeOptions
	options         AntTreeOptions
	err             error
	processTree     bool
}

func NewTree() *AntTree {
	return &AntTree{
		originOptions:   make(AntTreeOptions, 0),
		deepCopyOptions: make(AntTreeOptions, 0),
		options:         make(AntTreeOptions, 0),
	}
}

func (a *AntTree) SetOption(option *AntTreeOption) {
	a.originOptions = append(a.originOptions, option)
	a.deepCopyOptions = append(a.deepCopyOptions, option)
}

func (a *AntTree) WithTree() *AntTree {
	a.options = a.deepCopyOptions.toTree()
	return a
}

func (a *AntTree) WithLeafEnable() *AntTree {
	a.originOptions.withLeafEnable()
	return a
}

func (a *AntTree) GetOptions() AntTreeOptions {
	return a.options
}

type AntTreeOption struct {
	Title    string          `json:"title"`
	Pid      interface{}     `json:"pid"` // 父级id
	Key      interface{}     `json:"key"` // 自己的id
	Disabled bool            `json:"disabled"`
	IsLeaf   bool            `json:"isLeaf"`
	Children *AntTreeOptions `json:"children,omitempty"`
}

type AntTreeOptions []*AntTreeOption

func (m AntTreeOptions) toTree() AntTreeOptions {
	mTreeMap := make(map[interface{}]*AntTreeOption)
	for _, item := range m {
		mTreeMap[item.Key] = item
	}

	list := make(AntTreeOptions, 0)
	for _, item := range m {
		// 筛选出父级节点
		if item.Pid == 0 {
			list = append(list, item)
			continue
		}

		if pItem, ok := mTreeMap[item.Pid]; ok {
			// 如果存在父级节点，那么设置为叶子节点
			item.IsLeaf = true
			// 如果存在子级节点，那么设置叶子节点为false
			pItem.IsLeaf = false
			if pItem.Children == nil {
				children := AntTreeOptions{item}
				pItem.Children = &children
				continue
			}
			*pItem.Children = append(*pItem.Children, item)
		}
	}
	return list
}

// 只允许叶子节点被选择
func (m AntTreeOptions) withLeafEnable() AntTreeOptions {
	for _, value := range m {
		if value.Children != nil {
			// 父亲节点不让选中
			value.Disabled = true
			// 如果子节点存在继续遍历
			value.Children.withLeafEnable()
		}
	}
	return m
}
