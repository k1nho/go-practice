package DS

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	root *BSTNode
	size int
}

func (node BSTNode) String() string {
	return strconv.Itoa(node.val)
}

func (bstTree BST) String() string {
	sb := strings.Builder{}
	bstTree.Inorder(&sb, bstTree.root)
	return sb.String()
}

func (bstTree BST) Inorder(sb *strings.Builder, root *BSTNode) {
	// left root right
	if root != nil {
		bstTree.Inorder(sb, root.left)
		sb.WriteString(fmt.Sprintf("%s ", root))
		bstTree.Inorder(sb, root.right)
	}
}

// BST OPERATIONS
func (bstTree *BST) Insert(x int) {
	bstTree.root = bstTree.InsertNode(bstTree.root, x)
	bstTree.size++
}

func (bstTree *BST) InsertNode(root *BSTNode, x int) *BSTNode {
	if root == nil {
		return &BSTNode{value: x}
	}

	if x < root.val {
		root.left = bstTree.InsertNode(root.left, x)
	} else {
		root.right = bstTree.InsertNode(root.right, x)
	}

	return root
}

func (bstTree BST) Search(x int) (*BSTNode, bool) {
	return bstTree.SearchNode(bstTree.root, x)
}

func (bstTree BST) SearchNode(root *bstTree, x int) (*BSTNode, bool) {

	if root != nil {
		if x == root.val {
			return root, true
		} else if x < root.val {
			return SearchNode(root.left, x)
		} else {
			return SearchNode(root.right, x)
		}
	}
	return nil, false
}
