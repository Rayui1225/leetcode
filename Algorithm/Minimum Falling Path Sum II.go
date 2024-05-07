/************************************************************
* Initially, identify the smallest and second smallest numbers. Note that each time the smallest number is updated, the second smallest number should be changed to the previously identified smallest number.
* Iterate through each row using an index, and check whether the current index in the last row corresponds to the smallest number.
* If the current index in the last row is not the smallest number, update this element by adding it to the previously stored smallest number. If it is the smallest number, add the second smallest number instead.
* Finally, return the smallest number from the last row after all updates.
time complexity O(n^2)
space complexity O(n)
************************************************************/
func minFallingPathSum(grid [][]int) int {
    smallest := 100
    s_smallest := 100
    n := len(grid)
    for _,v := range grid[0]{
        if v < smallest{
            s_smallest = smallest 
            smallest = v
        }else if v < s_smallest{
            s_smallest = v
        }
    }
    dp := make([][]int, 2)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    dp[0] = grid[0]
    for i := 1; i < n; i++{
        nsmallest := 20000
        ns_smallest := 20000
        for j := 0; j < n; j++{
            if dp[(i-1)%2][j] != smallest || smallest == s_smallest{
               dp[i%2][j] = grid[i][j] + smallest
               if dp[i%2][j] < nsmallest{
                  ns_smallest = nsmallest
                  nsmallest = dp[i%2][j]
               }else if dp[i%2][j] < ns_smallest{
                  ns_smallest = dp[i%2][j]
               }
            }else{
               dp[i%2][j] = grid[i][j] + s_smallest
               if dp[i%2][j] < nsmallest{
                  ns_smallest = nsmallest
                  nsmallest = dp[i%2][j]
               }else if dp[i%2][j] < ns_smallest{
                  ns_smallest = dp[i%2][j]
               }
            }
        } 
        smallest = nsmallest
        s_smallest = ns_smallest   
    } 
    return smallest
}
/************************************************************
I don't need to create space to store , can just use grid to replace dp
************************************************************/
