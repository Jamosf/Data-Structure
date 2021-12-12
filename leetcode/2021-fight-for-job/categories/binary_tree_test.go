package categories

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

// tag-[二叉树]
// 第一题
// leetcode617: 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	m := &TreeNode{}
	m.Val = root1.Val + root2.Val
	m.Left = mergeTrees(root1.Left, root2.Left)
	m.Right = mergeTrees(root1.Right, root2.Right)

	return m
}

var (
	direction = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited   [][]bool
)

type pos struct {
	x int
	y int
}

// tag-[二叉树]
// 第一题
// leetcode116: 填充每个节点的下一个右侧节点指针
func connect(root *NodeC) *NodeC {
	if root == nil {
		return nil
	}
	queue := []*NodeC{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i < len(tmp)-1 {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}
	return root
}

func Test_connect(t *testing.T) {
	fmt.Println(connect(&NodeC{1, &NodeC{2, &NodeC{4, nil, nil, nil},
		&NodeC{5, nil, nil, nil}, nil}, &NodeC{3, &NodeC{6, nil, nil, nil}, &NodeC{7, nil, nil, nil}, nil}, nil}))
}

// tag-[二叉树]
// 第十二题
// leetcode 剑指offer 27: 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

// tag-[二叉树]
// 第十题
// leetcode 剑指offer 54: 二叉搜索树的第K大节点
func kthLargest(root *TreeNode, k int) int {
	var res int
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Right)
		if k == 0 {
			return
		}
		if k--; k == 0 {
			res = root.Val
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}

// tag-[二叉树]
// 第一题
// leetcode144: 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)

	return ret
}

// tag-[二叉树]
// 第二题
// leetcode94: 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)

	return ret
}

// tag-[二叉树]
// 第三题
// leetcode145: 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, postorderTraversal(root.Left)...)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)

	return ret
}

// tag-[二叉树]
// 第五题
// leetcode104: 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// tag-[二叉树]
// 第六题
// leetcode101: 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check(left.Right, right.Left) && check(left.Left, right.Right)
}

// tag-[二叉树]
// 第七题
// leetcode226: 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// tag-[二叉树]
// 第八题
// leetcode112: 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

// tag-[二叉树]
// 第九题
// leetcode102: 二叉树的层序遍历
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		levelNum := queue.Len()
		var tmp []int
		for i := 0; i < levelNum; i++ {
			v := queue.Front()
			queue.Remove(v)
			value := v.Value.(*TreeNode)
			tmp = append(tmp, value.Val)
			if value.Left != nil {
				queue.PushBack(value.Left)
			}
			if value.Right != nil {
				queue.PushBack(value.Right)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

// tag-[二叉树]
// 第五题
// leetcode700: 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

// tag-[二叉树]
// 第一题
// leetcode701: 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// tag-[二叉树]
// 第二题
// leetcode98: 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return isValidBST(root.Right)
}

// tag-[二叉树]
// leetcode98: 验证二叉搜索树
// 方法2
func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val >= upper || root.Val <= lower {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// tag-[二叉树]
// 第三题
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return dfs(root, m, k)
}

func dfs(root *TreeNode, m map[int]struct{}, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[root.Val]; ok {
		return true
	}
	m[k-root.Val] = struct{}{}
	return dfs(root.Left, m, k) || dfs(root.Right, m, k)
}

// tag-[二叉树]
// 第四题
// leetcode235: 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		if left := lowestCommonAncestor(root.Left, p, q); left != nil {
			return left
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		if right := lowestCommonAncestor(root.Right, p, q); right != nil {
			return right
		}
	}
	return root
}

// tag-[二叉树]
// 第三题
// leetcode 剑指offer32-II: 从上到下打印二叉树
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		levelNum := len(queue)
		var tmp []int
		for i := 0; i < levelNum; i++ {
			value := queue[0]
			tmp = append(tmp, value.Val)
			queue = queue[1:]
			if value.Left != nil {
				queue = append(queue, value.Left)
			}
			if value.Right != nil {
				queue = append(queue, value.Right)
			}
		}
		result = append(result, tmp)
	}
	return result
}

// tag-[二叉树]
// 第一题
// leetcode 剑指offer55-II: 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if minusAbs(depth(root.Left), depth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

// tag-[二叉树]
// 第二题
// leetcode 剑指offer 68-II: 二叉树的最近公共祖先
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}

// tag-[二叉树]
// 第六题
// leetcode 剑指offer26: 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if isEqual1(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isEqual1(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	return a.Val == b.Val && isEqual1(a.Left, b.Left) && isEqual1(a.Right, b.Right)
}

// tag-[二叉树]
// 第二题
// leetcode654：最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx, maxn := maxNumAndIdx(nums)
	root := &TreeNode{Val: maxn}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}

func maxNumAndIdx(nums []int) (int, int) {
	idx, maxn := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxn {
			maxn = nums[i]
			idx = i
		}
	}
	return idx, maxn
}

func Test_constructMaximumBinaryTree(t *testing.T) {
	r := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(r)
}

// tag-[二叉树]
// leetcode1993: 树上的操作
type lockStatus struct {
	user     int
	isLocked bool
	parent   int
	child    []int
}

type LockingTree struct {
	lockStates map[int]*lockStatus // 节点是否被上锁，以及上锁者是谁
	parent     []int
}

func ConstructorLockingTree(parent []int) LockingTree {
	l := LockingTree{lockStates: make(map[int]*lockStatus), parent: make([]int, len(parent))}
	for i := range parent {
		_, ok := l.lockStates[i]
		if !ok {
			l.lockStates[i] = &lockStatus{parent: parent[i]}
		}
		vv, ok := l.lockStates[parent[i]]
		if !ok {
			l.lockStates[parent[i]] = &lockStatus{child: append([]int{}, i)}
		} else {
			vv.child = append(vv.child, i)
		}
		l.parent[i] = parent[i]
	}
	return l
}

func (l *LockingTree) Lock(num int, user int) bool {
	v, ok := l.lockStates[num]
	if ok && v.isLocked {
		return false
	}
	v.isLocked = true
	v.user = user
	return true
}

func (l *LockingTree) Unlock(num int, user int) bool {
	v, ok := l.lockStates[num]
	if ok && v.isLocked && v.user == user {
		v.isLocked = false
		return true
	}
	return false
}

func (l *LockingTree) Upgrade(num int, user int) bool {
	v, ok := l.lockStates[num]
	if !ok || v.isLocked {
		return false
	}
	for l.parent[num] != -1 {
		num = l.parent[num]
		if v, ok := l.lockStates[num]; ok && v.isLocked {
			return false
		}
	}
	// 子节点被上锁
	flag := false
	l.isChildLocked(v, &flag)
	if !flag {
		return false
	}
	v.isLocked = true
	v.user = user
	return true
}

func (l *LockingTree) isChildLocked(v *lockStatus, flag *bool) {
	for _, child := range v.child {
		v, ok := l.lockStates[child]
		if ok && v.isLocked {
			*flag = true
			v.isLocked = false
		}
		l.isChildLocked(v, flag)
	}
}

func Test_lock(t *testing.T) {
	lockTree := ConstructorLockingTree([]int{-1, 0, 0, 1, 1, 2, 2})
	a := lockTree.Lock(2, 2)
	b := lockTree.Unlock(2, 3)
	c := lockTree.Unlock(2, 2)
	d := lockTree.Lock(4, 5)
	e := lockTree.Upgrade(0, 1)
	f := lockTree.Lock(0, 1)
	fmt.Println(a, b, c, d, e, f)
}

// tag-[二叉树]
// 第五题
// leetcode114: 二叉树展开为链表
func flatten(root *TreeNode) {
	dummy := &TreeNode{}
	p := dummy
	var traval func(r *TreeNode)
	traval = func(r *TreeNode) {
		if r == nil {
			return
		}
		dummy.Right = &TreeNode{Val: r.Val}
		dummy = dummy.Right
		traval(r.Left)
		traval(r.Right)
	}
	traval(root)
	if root == nil {
		return
	}
	root.Left = nil
	root.Right = p.Right.Right
}

func Test_flatten(t *testing.T) {
	flatten(&TreeNode{Right: &TreeNode{1, nil, nil}, Left: &TreeNode{2, nil, nil}, Val: 0})
}

// tag-[二叉树]
// leetcode543: 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	maxn := 0
	var depth func(r *TreeNode) int
	depth = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		L := depth(r.Left)
		R := depth(r.Right)
		maxn = max(maxn, L+R-1)
		return 1 + max(L, R)
	}
	depth(root)
	return maxn
}

// tag-[二叉树]
// 第一题
// leetcode105: 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	r := &TreeNode{Val: preorder[0]}
	idx := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			idx = i
		}
	}
	r.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	r.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return r
}

func Test_buildTree(t *testing.T) {
	r := buildTree([]int{-1}, []int{-1})
	fmt.Println(r)
}

// tag-[二叉树]
// 第二题
// leetcode236: 二叉树的最近公共祖先
// 二叉树
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	// 左边没有找到
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

// tag-[二叉树]
// 第二题
// leetcode538: 二叉搜索树转换为累加树
// 二叉树、反向中序遍历
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Right)
		sum += r.Val
		r.Val = sum
		dfs(r.Left)
	}
	dfs(root)
	return root
}

// tag-[二叉树]
// leetcode106, 可以使用map优化查找节点的idx
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	n := len(postorder)
	root := &TreeNode{Val: postorder[n-1]}
	idx := n - 1
	for idx >= 0 && inorder[idx] != postorder[n-1] {
		idx--
	}
	root.Right = buildTree1(inorder[idx+1:], postorder[idx:n-1])
	root.Left = buildTree1(inorder[:idx], postorder[:idx])
	return root
}

// tag-[二叉树]
// leetcode99:二叉搜索树中序遍历,找出错位的两个节点并交换
// 时间复杂度：O(n)
// 空间复杂度：O(logn), 即为树的高度
func recoverTree(root *TreeNode) {
	var firstMax, lastMin *TreeNode
	pre := &TreeNode{Val: math.MinInt32}
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		if r.Val < pre.Val {
			lastMin = r
			if firstMax == nil {
				firstMax = pre
			}
		}
		pre = r
		inorder(r.Right)
	}
	inorder(root)
	if firstMax != nil && lastMin != nil {
		firstMax.Val, lastMin.Val = lastMin.Val, firstMax.Val
	}
}

// tag-[二叉树]
// leetcode99: 中序遍历栈的写法，先将跟节点和所有的左节点压栈，然后逐个取出左节点，压入右节点。
func recoverTree_(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

// morris遍历的代码，通过左子树的最右边链接根节点，可以在遍历左边结束后找回到根节点，实现回溯的功能，但是又没有栈的开销。
func morrisInorder(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			// 如果没有左子树，则直接走右子树
			cur = cur.Right
		} else {
			pre := cur.Left
			// 遍历到最右
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == cur {
				// 意味着前驱节点的右节点已被设置，该次遍历为回溯
				// 左边已经搞定，接下来需要处理右边
				pre.Right = nil
				cur = cur.Right
			} else {
				// 第一次访问前驱节点，设置线索，即右节点为当前节点
				pre.Right = cur
				cur = cur.Left
			}
		}
	}
}

// tag-[二叉树]
// leetcode99:morris解法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func recoverTree__(root *TreeNode) {
	var x, y, pred, predecessor *TreeNode // pred记录上一个
	cur := root
	for cur != nil {
		if cur.Left == nil {
			if pred != nil && cur.Val < pred.Val {
				y = cur
				if x == nil {
					x = pred
				}
			}
			pred = cur
			cur = cur.Right
		} else {
			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
			predecessor = cur.Left
			for predecessor.Right != nil && predecessor.Right != cur {
				predecessor = predecessor.Right
			}

			// 让 predecessor 的右指针指向 root，继续遍历左子树
			if predecessor.Right != nil { // 说明左子树已经访问完了，我们需要断开链接
				if pred != nil && cur.Val < pred.Val {
					y = cur
					if x == nil {
						x = pred
					}
				}
				pred = cur
				predecessor.Right = nil
				cur = cur.Right
			} else {
				predecessor.Right = cur
				cur = cur.Left
			}
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func Test_recoverTree(t *testing.T) {
	recoverTree(&TreeNode{Val: 1, Left: &TreeNode{3, nil, &TreeNode{2, nil, nil}}})
}

// tag-[二叉树]
// leetcode437: 不一定从根节点开始的路径和
func pathSumIII(root *TreeNode, targetSum int) int {
	sum := 0
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		sum += pathSumCnt(r, targetSum)
		inorder(r.Left)
		inorder(r.Right)
	}
	inorder(root)
	return sum
}

// tag-[二叉树]
// leetcode230: 二叉搜索树中k小的元素
// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		k--
		if k == 0 {
			ans = r.Val
		}
		inorder(r.Right)
	}
	inorder(root)
	return ans
}

// tag-[二叉树]
// leetcode331: 验证二叉树的前序序列化
func isValidSerialization(preorder string) bool {
	n := len(preorder)
	stack := []int{1}
	for i := 0; i < n; {
		if len(stack) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			i++
		} else {
			for i < n && preorder[i] != ',' {
				i++
			}
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, 2)
		}
	}
	return len(stack) == 0
}

// tag-[二叉树]
// 二叉树的前序遍历迭代遍历
func preorderTraversal_(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := make([]int, 0)
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		v := l.Back()
		l.Remove(v)
		r := v.Value.(*TreeNode)
		out = append(out, r.Val)
		if r.Right != nil {
			l.PushBack(r.Right)
		}
		if r.Left != nil {
			l.PushBack(r.Left)
		}
	}
	return out
}

// tag-[二叉树]
// leetcode297: 二叉树的序列化和反序列化（本提可以用前序遍历和后序遍历、层序遍历来解决，但是中序遍历无法解）
// 前序遍历
type Codec struct {
}

func ConstructorC() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var preorder func(node *TreeNode) string
	preorder = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		return strconv.Itoa(node.Val) + "," + preorder(node.Left) + "," + preorder(node.Right)
	}
	return preorder(root)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	var preorder func(s *[]string) *TreeNode
	preorder = func(s *[]string) *TreeNode {
		if len(*s) == 0 {
			return nil
		}
		if (*s)[0] == "#" {
			*s = (*s)[1:]
			return nil
		}
		v, _ := strconv.Atoi((*s)[0])
		*s = (*s)[1:]
		root := &TreeNode{Val: v}
		root.Left = preorder(s)
		root.Right = preorder(s)
		return root
	}
	return preorder(&ss)
}

func Test_codec(t *testing.T) {
	c := &Codec{}
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	r := c.deserialize(c.serialize(root))
	fmt.Println(r)
}

// tag-[二叉树]
// leetcode297: 二叉树的序列化和反序列化
// 层序遍历
type Codec_ struct {
}

func ConstructorC_() Codec_ {
	return Codec_{}
}

// Serializes a tree to a single string.
func (c *Codec_) serialize(root *TreeNode) string {
	var ss []string
	q := []*TreeNode{root}
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		if v == nil {
			ss = append(ss, "#")
			continue
		}
		ss = append(ss, strconv.Itoa(v.Val))
		q = append(q, v.Left)
		q = append(q, v.Right)
	}
	return strings.Join(ss, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec_) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	v, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: v}
	q := []*TreeNode{root}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		left, right := ss[1], ss[2]
		if left != "#" {
			v, _ := strconv.Atoi(left)
			node.Left = &TreeNode{Val: v}
			q = append(q, node.Left)
		}
		if right != "#" {
			v, _ := strconv.Atoi(right)
			node.Right = &TreeNode{Val: v}
			q = append(q, node.Right)
		}
		ss = ss[2:]
	}
	return root
}

func Test_codec_(t *testing.T) {
	c := &Codec_{}
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	r := c.deserialize(c.serialize(root))
	fmt.Println(r)
}
