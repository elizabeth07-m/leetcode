type set struct {
	pass, total int
}

func (s set) gain(i int) float64 {
	return float64(s.pass+i)/float64(s.total+i) - float64(s.pass)/float64(s.total)
}

type MaxHeap []set

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].gain(1) > m[j].gain(1)
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(set))
}

func (m *MaxHeap) Pop() any {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}


func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	maxheap := &MaxHeap{}
	heap.Init(maxheap)

	for _, class := range classes {
		heap.Push(maxheap, set{
			pass:  class[0],
			total: class[1],
		})
	}
	
	for extraStudents > 0 {
		curr := heap.Pop(maxheap).(set)
		heap.Push(maxheap, set{curr.pass + 1, curr.total + 1})
		extraStudents--
	}
	
	res := 0.0
	for maxheap.Len() > 0 {
		curr := heap.Pop(maxheap).(set)
		res += float64(curr.pass) / float64(curr.total)
	}
	
	return res / float64(len(classes))
}