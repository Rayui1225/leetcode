/************************************************************
* I initially used the difference between the two element in costs  as a sorting key
* and then calculated the sum of the first 'n' elements for city A and the last 'n' elements for city B.
* Although this method passed the test data, I wasn't sure why it was correct.
* Later, I found a good explanation: Assume that all members initially plan to go to city A, giving us a total sum of costs. 
* Then, we decide to send 'n' members to city B instead. If a member 'i' switches to city B, the extra cost incurred is costs[i][0] - costs[i][1], which could be negative. 
* To minimize the total extra cost, we should select the 'n' members with the smallest values of costs[i][0] - costs[i][1]
************************************************************/

func twoCitySchedCost(costs [][]int) int {
    n := len(costs)/2
    for k,v := range costs{
        costs[k]=append(costs[k],v[0]-v[1])
    }
    sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})
    res := 0
    tmp := 0
    for _,v := range costs{
        if tmp < n{
            res += v[0]
            tmp += 1
        }else{
            res += v[1]
        }
    }
    return res
}

/************************************************************
*Areas that can be improved in this code

func(i, j int) bool {
		return costs[i][2] < costs[j][2]  *this way could be costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1] can save some spaces
	}
 
for _,v := range costs{
        if tmp < n{
            res += v[0]
            tmp += 1
        }else{
            res += v[1]
        }
    }
    
can change to 

for i := 0; i < n; i++{
res += costs[i][0] + costs[i+n][i]
}
can save some times
************************************************************/
