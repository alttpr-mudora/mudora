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

func (i *Items) Clone() *Items {
	counts := make(map[string]int, len(i.counts))
	for name, n := range i.counts {
		counts[name] = n
	}
	return &Items{counts: counts}
}

func (i *Items) Count(name string) int {
	return i.counts[name]
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
