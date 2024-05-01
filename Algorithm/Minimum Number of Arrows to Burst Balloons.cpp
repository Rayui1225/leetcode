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
