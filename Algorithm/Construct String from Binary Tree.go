/************************************************************
* Perform Depth First Search (DFS) while adhering to the following rules:
* If the node is a leaf, return null.
* If the left child is null, return the string formed by concatenating the node value, empty parentheses for the missing left child, and the result of the DFS on the right child.
* If the right child is null, return the string formed by concatenating the node value and the result of the DFS on the left child (since no parentheses are needed for the right child when it is absent).
time complexity O(n)
space complexity O(1)
************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strconv"

func tree2str(root *TreeNode) string {
    if root == nil {
        return ""
    }
    
    stringL := tree2str(root.Left)
    stringR := tree2str(root.Right)
    if stringL == "" && stringR == ""{
        return strconv.Itoa(root.Val) 
    }else if stringR == ""{
        return strconv.Itoa(root.Val) + "(" + stringL + ")"    
    }else if stringL == ""{
        return strconv.Itoa(root.Val) + "()"+"(" + stringR + ")"    
    }else{
        return strconv.Itoa(root.Val) + "(" + stringL + ")" + "(" + stringR + ")"  
    }   
}
/************************************************************
I can just check current node left and right child , if that node is not null then preform DFS
************************************************************/
