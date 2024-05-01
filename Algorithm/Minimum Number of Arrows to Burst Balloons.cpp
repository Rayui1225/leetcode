/************************************************************
*Reframe the problem as determining how many points overlap with a reference point (note that there may be multiple reference points).
*To check if a point overlaps with a reference point, verify if Xstart is greater than the reference point. Consequently, sort the points vector to facilitate a single pass through it.
*As you iterate through the vector, increment the result counter when a point does not overlap with the current reference point, then update the reference point to this new non-overlapping point.
************************************************************/    
class Solution {
public:
    static bool cmp(vector<int> a,vector<int> b){
    if(a[1]!=b[1]){
        return a[1]<b[1]; 
    }
    return a[0]<b[0];
    }
    
    int findMinArrowShots(vector<vector<int>>& points) {
        sort(points.begin(),points.end(),cmp);
        int tmp=points[0][1];
        int res=1;
        for(auto p :points){
            if(p[0]>tmp){
                res++;
                tmp=p[1];
            }
        }
        return res;
    }
};
/************************************************************
somewhere can improve
I can just compare a[0] < a[1] and if p[0] < tmp we just change tmp to  min(tmp,p[1])
************************************************************/  
