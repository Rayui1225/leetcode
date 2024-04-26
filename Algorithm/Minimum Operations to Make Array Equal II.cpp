/************************************************************
* We need to compare all elements in nums1 and nums2 by index
* Calculate the difference (diff) for each corresponding element at index i in nums1 and nums2.
* If K equals zero, check if all differences are zero. If so, return 0; otherwise, return -1 (indicates failure).
* If K is not zero, check if each diff modulo k is zero. If so, add diff divided by k to the result (res); otherwise, return -1 (indicates failure).
* Finally return res 
************************************************************/

class Solution {
public:
    long long minOperations(vector<int>& nums1, vector<int>& nums2, long long k) {
        int len=nums1.size();
        vector<long long> diff(len,0);
        long long sum=0;
        for(int i=0;i<len;i++){
          long long tmp=nums1[i]-nums2[i];
          if(k!=0){
          if(tmp%k!=0){
            return -1;
          } 
          }else{
            if(tmp!=0){
             return -1;
            }
          }
          sum+=tmp;
          diff[i]=tmp;
        }
        if(sum!=0){
            return -1;
        }
        if(k==0){
            return 0;
        }
        long long res=0;
        for(int i=0;i<len;i++){
            if(diff[i]>0){
                res+=diff[i]/k;
            }
        }
        return res;
    }
};

/************************************************************
Something can improve
if(k==0 && sum !=0){return -1} else{...} to replace original if condition at line 17
I did't need to store each diff in each corresponding index at diff vector I can store abs(tmp) to a variable (for examle diff) then res = diff/k/2  
************************************************************/
