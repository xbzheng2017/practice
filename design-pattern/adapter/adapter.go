package adapter

type ScoreOperation interface {
	Sort([]float32) []float32
}

type QuickSortAlg struct{}

func (q *QuickSortAlg) QuickSort(scores []float32) []float32 {
	//fake implement
	return []float32{1.0, 2.0, 3.0}
}

type QuickSortStructAdapter struct {
	*QuickSortAlg
}

func (q *QuickSortStructAdapter) Sort(scores []float32) []float32 {
	return q.QuickSort(scores)
}

func NewQuickSortStructAdapter() *QuickSortStructAdapter {
	return new(QuickSortStructAdapter)
}

type QuickSortObjectAdapter struct {
	qk *QuickSortAlg
}

func (q *QuickSortObjectAdapter) Sort(scores []float32) []float32 {
	return q.qk.QuickSort(scores)
}

func NewQuickSortObjectAdapter() *QuickSortObjectAdapter {
	return new(QuickSortObjectAdapter)
}
