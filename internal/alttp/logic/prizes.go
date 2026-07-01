package logic

var crystals = []string{
	"Crystal 1", "Crystal 2", "Crystal 3", "Crystal 4", "Crystal 5", "Crystal 6", "Crystal 7",
}

func (i *Items) CrystalCount() int {
	count := 0
	for _, name := range crystals {
		if i.Has(name) {
			count++
		}
	}

	return count
}
