package fenwick

type FenwickTree struct {
	tree []int
}

func MakeFenwickTree(arr []int) *FenwickTree {
	i := 1
	tree := make([]int, len(arr)+1)
	for _, num := range arr {
		for j := 0; ; j++ {
			if i&(1<<j) == (1 << j) {
				var cur, k int = num, 1
				for range j {
					cur += tree[i-k]
					k <<= 1
				}
				tree[i] = cur
				break
			}
		}
		i++
	}
	return &FenwickTree{tree}
}

func (f *FenwickTree) Append(values ...int) {
	i := len(f.tree)
	if required := i + len(values); required > cap(f.tree) {
		f.increaseCapacity(required)
	} else {
		f.tree = f.tree[:required]
	}

	for _, num := range values {
		for j := 0; ; j++ {
			if i&(1<<j) == (1 << j) {
				var cur, k int = num, 1
				for range j {
					cur += f.tree[i-k]
					k <<= 1
				}
				f.tree[i] = cur
				break
			}
		}
		i++
	}
}

func (f *FenwickTree) increaseCapacity(requiredSize int) {
	newCap := cap(f.tree)
	for newCap < requiredSize {
		newCap <<= 1
	}

	newTree := make([]int, requiredSize, newCap)
	copy(newTree, f.tree)
	f.tree = newTree
}

func (f *FenwickTree) SumFirstK(k int) (sum int) {
	for k > 0 {
		sum += f.tree[k]
		k &= k - 1
	}
	return
}

func (f *FenwickTree) SumRange(from, to int) int {
	if from > to {
		return 0
	}
	return f.SumFirstK(to) - f.SumFirstK(from-1)
}

func (f *FenwickTree) Add(index, value int) {
	for index < len(f.tree) {
		f.tree[index] += value
		index += index & -index
	}
}
