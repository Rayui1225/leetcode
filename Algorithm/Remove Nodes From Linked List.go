/************************************************************
*Operate just like a basic Monotonic Stack, but apply it to a linked list. Therefore, the stored data needs to be changed to addresses instead of values.
time complexity O(n)
space complexity O(n)
************************************************************/

func removeNodes(head *ListNode) *ListNode {
    var stack []*ListNode
    stack = append(stack,head)
    for head.Next != nil{
        head = head.Next
        val := head.Val
        n := len(stack)
        for n != 0 && val > stack[n-1].Val{
            stack = stack[:n-1]
            n = n-1
        }
        if n > 0{
            stack[n-1].Next = head 
        }
        stack = append(stack,head)
    }
    return stack[0]
}
/************************************************************
another way to operate like Monotonic Stack , reverse this linked list , when we meet number greater then last number
we can just go back to previous number until null or the number greater then current number
************************************************************/
