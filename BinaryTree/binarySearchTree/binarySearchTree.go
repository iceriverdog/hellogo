package binarySearchTree

type Tree []int

func NewTree(n int) Tree {
	arr := make([]int, n, n)

	arr[0] = 100
	return arr
}

func (tree Tree) Insert(value int) Tree {
	if tree[1] == 0 {
		tree[1] = value
		return tree
	}
	i := 1
	//去左还是去右，满足条件就退出，否则，循环下去
	for {
		if value > tree[i] {
			i = 2*i + 1
			if tree[i] == 0 {
				tree[i] = value
				return tree
			}
		}
		if value < tree[i] {
			i = 2 * i
			if tree[i] == 0 {
				tree[i] = value
				return tree
			}
		}
	}
}
