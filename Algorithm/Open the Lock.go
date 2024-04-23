/************************************************************
* Use map to record all deadends
* Implement two functions, addNumber and subNumber, to calculate the results of incrementing or decrementing each digit of the original number.
* Use a queue (or alternatively, a stack) to store all numbers returned from addNumber and subNumber.(the first number to execute addNumber and subNumber is target )
* Pop each number from the queue and check if it is in the deadends or if it is '0000'. If yes, return the result or continue executing the for loop.
* If not, execute addNumber and subNumber for this number, and store the return values in another queue. 
* After all values are popped from the first queue, copy the contents from the second queue to the first queue and then clear the second queue.
* Repeat the process until the first queue is empty or you find '0000'.
************************************************************/

func openLock(deadends []string, target string) int {
    if target == "0000"{
        return 0
    }
    m := make(map[string]int)
    for _,v := range deadends{
        m[v] = 1
    }
    if m["0000"] == 1{
        return -1
    }
    res := 0
    var queue1 []string
    var queue2 []string
    queue1 = append(queue1,addnumber(target)...)
    queue1 = append(queue1,subnumber(target)...)
    for len(queue1) != 0{
        res += 1
        for _,v := range queue1{
        if v != "0000"{
        if m[v] != 1{     
        queue2 = append(queue2,addnumber(v)...)
        queue2 = append(queue2,subnumber(v)...)
        m[v] = 1
        }
        }else{
            return res
        }
        }
        queue1 = queue2
        queue2 = []string{}
    }
    return -1
}
func addnumber(s string) []string{
    var res []string
    for i:=0; i<4; i++{
        chars := []rune(s)
        currentVal := int(chars[i] - '0') 
        newVal := (currentVal + 1)%10
        chars[i] = rune(newVal + '0')
        res = append(res,string(chars))
    }
    return res 
}

func subnumber(s string) []string{
    var res []string
    for i:=0; i<4; i++{
        chars := []rune(s)
        currentVal := int(chars[i] - '0') 
        newVal := currentVal - 1
        if newVal == -1{
            newVal = 9
        }
        chars[i] = rune(newVal + '0')
        res = append(res,string(chars))
    }
    return res 
}
/************************************************************
something can be improved in this code
* I didn't need to change string to []rune and change back just directily compare []rune that enough
* subnumber I can use +9 instead of -1 and check greater than zero or not because +9 mod 10 = -1 mod 10
************************************************************/
