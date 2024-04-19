/************************************************************
* A intuition thinking way select the largest difference between (x1/x2) and ((x1+1)/(x2+1))
* We can use priority queue to implement this requirment
************************************************************/
import "container/heap"
type ArrayHeap [][]float64
func (h ArrayHeap) Len() int { return len(h) }
func (h ArrayHeap) Less(i, j int) bool { return h[i][2] > h[j][2] }
func (h ArrayHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *ArrayHeap) Push(x interface{}) {
    *h = append(*h, x.([]float64))
}
func (h *ArrayHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    h := &ArrayHeap{}
    heap.Init(h)
    for _,v := range classes{
        v0 := float64(v[0])
        v1 := float64(v[1])
        heap.Push(h,[]float64{v0,v1,(v0+1)/(v1+1)-(v0/v1)})
    }
    for i:=0; i<extraStudents; i++{
        top := heap.Pop(h).([]float64)
        v0 := top[0]
        v1 := top[1]
        heap.Push(h,[]float64{v0+1,v1+1,(v0+2)/(v1+2)-(v0+1)/(v1+1)})
    }
    var res float64 
    for h.Len() > 0 {
        top := heap.Pop(h).([]float64)
        res += top[0]/top[1]
    }
    return res/float64(len(classes))
}
/************************************************************
something can be improved in this code
Don't need to store difference between (x1/x2) and ((x1+1)/(x2+1))
Just change Less function to h[i][0]+1 / h[i][1]+1 > h[j][0]+1 / h[j][1]+1
can save some space 
************************************************************
