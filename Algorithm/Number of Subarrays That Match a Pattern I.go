/************************************************************
* We can create an array to store the results of converting elements from the nums array to a pattern based on specific conditions.
* Once an element is converted, it will not be used again.
* Thus, we can directly modify the original nums array into the desired format.
* Then, simply compare the converted elements with the pattern to calculate how many patterns match in this array.
time complexity O(n*m) (convert O(n) compare the converted elements O(n*m) )
space complexity O(1) (just use n,m,res)
************************************************************/

func countMatchingSubarrays(nums []int, pattern []int) int {
    n :=len(nums)
    for i := 0; i < n-1; i++{
        if nums[i] < nums[i+1]{
            nums[i] = 1
        }else if nums[i] > nums[i+1]{
            nums[i] = -1
        }else{
            nums[i] = 0
        }
    }
    m := len(pattern)
    res := 0
    for i := 0; i < n-m; i++{
        if nums[i] != pattern[0]{
            continue
        }else{
            flag := 1
            for j := 1; j < m; j++{
                if nums[i+j] != pattern[j]{
                    flag = 0
                    break
                }
            }
            if  flag == 1{
                 res += 1
            }
        } 
    } 
    return res
}
