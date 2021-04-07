package gotoant

// TreeSelect ant design树形结构 https://ant.design/components/tree-select-cn/#TreeNode-props
// Tree https://ant.design/components/tree-cn/#TreeNode-props
type AntTreeSelect struct {
	originOptions   AntTreeSelectOptions
	deepCopyOptions AntTreeSelectOptions
	options         AntTreeSelectOptions
	err             error
	processTree     bool
}

func NewTreeSelect() *AntTreeSelect {
	return &AntTreeSelect{
		originOptions:   make(AntTreeSelectOptions, 0),
		deepCopyOptions: make(AntTreeSelectOptions, 0),
		options:         make(AntTreeSelectOptions, 0),
	}
}

func (a *AntTreeSelect) SetOption(option *AntTreeSelectOption) {
	a.originOptions = append(a.originOptions, option)
	a.deepCopyOptions = append(a.deepCopyOptions, option)
}

func (a *AntTreeSelect) WithTree() *AntTreeSelect {
	a.options = a.deepCopyOptions.toTree()
	return a
}

func (a *AntTreeSelect) WithLeafEnable() *AntTreeSelect {
	a.originOptions.withLeafEnable()
	return a
}

func (a *AntTreeSelect) GetOptions() AntTreeSelectOptions {
	return a.options
}

type AntTreeSelectOption struct {
	Title    string                `json:"title"`
	Pid      interface{}           `json:"pid"` // 父级id
	Key      interface{}           `json:"key"` // 自己的id
	Disabled bool                  `json:"disabled"`
	IsLeaf   bool                  `json:"isLeaf"`
	Children *AntTreeSelectOptions `json:"children,omitempty"`
}

type AntTreeSelectOptions []*AntTreeSelectOption

func (m AntTreeSelectOptions) toTree() AntTreeSelectOptions {
	mTreeMap := make(map[interface{}]*AntTreeSelectOption)
	for _, item := range m {
		mTreeMap[item.Key] = item
	}

	list := make(AntTreeSelectOptions, 0)
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
				children := AntTreeSelectOptions{item}
				pItem.Children = &children
				continue
			}
			*pItem.Children = append(*pItem.Children, item)
		}
	}
	return list
}

// 只允许叶子节点被选择
func (m AntTreeSelectOptions) withLeafEnable() AntTreeSelectOptions {
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
