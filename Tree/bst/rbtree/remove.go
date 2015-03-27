package rbtree

//成功返回true，没有返回false。
//红黑树删除过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tree *Tree) Remove(key int) bool {
	tree.path.clear()
	var target = tree.root
	for target != nil && key != target.key {
		if key < target.key {
			tree.path.push(target, true)
			target = target.left
		} else {
			tree.path.push(target, false)
			target = target.right
		}
	}
	if target == nil {
		return false
	}

	var victim, orphan *node = nil, nil
	if target.left == nil {
		victim, orphan = target, target.right
	} else if target.right == nil {
		victim, orphan = target, target.left
	} else {
		tree.path.push(target, false)
		victim = target.right
		for victim.left != nil {
			tree.path.push(victim, true)
			victim = victim.left
		}
		orphan = victim.right
	}

	if tree.path.isEmpty() { //此时victim==target
		tree.root = orphan
		if tree.root != nil {
			tree.root.black = true
		}
	} else {
		tree.hookSubTree(orphan)
		if victim.black { //红victim随便删，黑的要考虑
			if orphan != nil && !orphan.black {
				orphan.black = true //红子变黑顶上
			} else {
				tree.adjustAfterDelete()
			}
		}
		target.key = victim.key //李代桃僵
	}
	return true
}

//----------------红叔模式----------------
//|        bG        |        bU        |
//|       /  \       |       /  \       |
//|     bO    rU     |     rG    bR     |
//|          /  \    |    /  \          |
//|        bL    bR  |  bO    bL        |

//------------------双黑------------------
//|        xG        |        bG        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bO    rU     |
//|          /  \    |          /  \    |
//|        bL    bR  |        bL    bR  |

//----------------中黑外红----------------
//|        xG        |        xU        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bR     |
//|          /  \    |    /  \          |
//|        bL    rR  |  bO    bL        |

//------------------中红------------------
//|        xG        |        xL        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bU     |
//|          /  \    |    /  \  /  \    |
//|        rL    xR  |  bO   u  v   xR  |
//|       /  \       |                  |
//|      u    v      |                  |

func (tree *Tree) adjustAfterDelete() {
	var G, lf = tree.path.pop()
	for { //剩下情况：victim黑，orphan也黑，此时victim(orphan顶替)的兄弟必然存在
		if lf {
			var U = G.right //U != nil
			var L, R = U.left, U.right
			if !U.black { //红P下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G.right, U.left = L, G
				U.black, G.black = true, false
				tree.hookSubTree(U)
				tree.path.push(U, lf)
				continue //变出黑U后再行解决
			} else {
				if L == nil || L.black {
					if R == nil || R.black { //双黑，变色解决
						U.black = false
						if G.black && !tree.path.isEmpty() {
							G, lf = tree.path.pop()
							continue //上溯
						}
						G.black = true
					} else { //中黑外红
						G.right, U.left = L, G
						U.black, G.black, R.black = G.black, true, true
						tree.hookSubTree(U)
					}
				} else { //中红
					U.left, G.right = L.right, L.left
					L.right, L.left = U, G
					L.black, G.black = G.black, true
					tree.hookSubTree(L)
				}
			}
		} else {
			var U = G.left //U != nil
			var R, L = U.right, U.left
			if !U.black { //红P下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G.left, U.right = R, G
				U.black, G.black = true, false
				tree.hookSubTree(U)
				tree.path.push(U, lf)
				continue //变出黑U后再行解决
			} else {
				if R == nil || R.black {
					if L == nil || L.black { //双黑，变色解决
						U.black = false
						if G.black && !tree.path.isEmpty() {
							G, lf = tree.path.pop()
							continue //上溯
						}
						G.black = true
					} else { //中黑外红
						G.left, U.right = R, G
						U.black, G.black, L.black = G.black, true, true
						tree.hookSubTree(U)
					}
				} else { //中红
					U.right, G.left = R.left, R.right
					R.left, R.right = U, G
					R.black, G.black = G.black, true
					tree.hookSubTree(R)
				}
			}
		}
		break //个别情况需要循环
	}
}
