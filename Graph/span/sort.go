package span

func sort(list []EdgeW) {
	lv := uint(1)
	for sz := len(list); sz != 0; sz /= 2 {
		lv++
	}
	doIntroSort(list, lv*2)
}
func doIntroSort(list []EdgeW, life uint) {
	if len(list) < 16 {
		simpleSort(list)
	} else if life == 0 {
		heapSort(list)
	} else {
		m := partition(list)
		doIntroSort(list[:m], life-1)
		doIntroSort(list[m:], life-1)
	}
}

func partition(list []EdgeW) int {
	pivot := list[len(list)/2].Weight
	a, b := 0, len(list)-1
	for {
		for list[a].Weight < pivot {
			a++
		}
		for list[b].Weight > pivot {
			b--
		}
		if a >= b {
			break
		}
		list[a], list[b] = list[b], list[a]
		a++
		b--
	}
	return a
}

func heapSort(list []EdgeW) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}
func down(list []EdgeW, root int) {
	key := list[root]
	kid, last := root*2+1, len(list)-1
	for kid < last {
		if list[kid+1].Weight > list[kid].Weight {
			kid++
		}
		if key.Weight >= list[kid].Weight {
			break
		}
		list[root] = list[kid]
		root, kid = kid, kid*2+1
	}
	if kid == last && key.Weight < list[kid].Weight {
		list[root], root = list[kid], kid
	}
	list[root] = key
}

func simpleSort(list []EdgeW) {
	if len(list) < 2 {
		return
	}
	best := 0
	for i := 1; i < len(list); i++ {
		if list[i].Weight < list[best].Weight {
			best = i
		}
	}
	list[0], list[best] = list[best], list[0]
	for i := 1; i < len(list); i++ {
		key, pos := list[i], i
		for list[pos-1].Weight > key.Weight {
			list[pos] = list[pos-1]
			pos--
		}
		list[pos] = key
	}
}
