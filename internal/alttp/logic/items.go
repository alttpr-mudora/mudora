package logic

type Items struct {
	counts map[string]int
}

func NewItems() *Items {
	return &Items{counts: make(map[string]int)}
}

func (i *Items) Add(name string) {
	i.counts[name]++
}

func (i *Items) HasAtLeast(name string, n int) bool {
	if n <= 0 {
		return true
	}

	return i.counts[name] >= n
}

func (i *Items) Has(name string) bool {
	return i.HasAtLeast(name, 1)
}
