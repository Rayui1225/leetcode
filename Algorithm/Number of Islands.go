/************************************************************
* Each time you reach a point whose value is '1', add one to the result.
* This is because when you reach that point, the DFS function will execute and change all reachable points from the origin to 'X'.
************************************************************/
func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])
    res := 0
    for i := 0; i < m; i++{
        for j :=0; j < n;j++{
            if grid[i][j] == '1'{
                dfs(&grid,i,j,m,n)
                res +=1
            }
        }
    }
    return res
}
func dfs(grid *[][]byte, x int, y int, m int, n int){
     directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
     (*grid)[x][y] = 'X'
     for _,v := range directions{
        nx := x+v[0]
        ny := y+v[1]
        if nx >= 0 && nx < m && ny >= 0 && ny < n && (*grid)[nx][ny] == '1'{
            dfs(grid,nx,ny,m,n)
        }
     }
}
/************************************************************
something can be improved in this code
Just about coding way someone use (but if more than 4 directions it will be lengthy)
dfs(grid,x+1,y) dfs(grid,x-1,y) ....
and set filter condition in first line of function dfs 
************************************************************
