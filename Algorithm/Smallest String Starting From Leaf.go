/************************************************************
* I initially used the smaller string from both the left and right subtrees, but this approach was incorrect.
* For example, "abd" is greater than "ab", but when combined with the root value (e.g., "z"), "abdz" is less than "abz"
* Therefore, we need to combine the characters of all nodes along the path up to the leaf and then return that string.
* The return value will follow these rules:
* 1. If this node is a leaf, return the combined characters of all nodes along the path up to this node.
* 2. If this node is not a leaf and one of its child nodes is null, return the smallest lexicographical string that can be formed from the other subtree.
* 3. If neither child node is null, return the smaller lexicographical string compared to the smallest lexicographical string that can be formed from the two subtrees.
************************************************************/

func smallestFromLeaf(root *TreeNode) string {
   return dfs(root,"")
}
func dfs(root *TreeNode, prestring string) string{
    if root == nil{
        return ""
    }
    now_string := string('a' + root.Val) + prestring
    left := dfs(root.Left,now_string)
    right := dfs(root.Right,now_string)
    if right == "" && left == ""{
        return now_string
    }else if right == ""{
        return left
    }else if left == ""{
        return right
    }else{
        if left > right{
            return  right
        }else{
            return left
        }
    }
}
/************************************************************
something can be improved in this code
* Initially I didn't use 'string('a' + root.Val)' to change  number to character , I use a slice to change that (e.g., s[0] = 'a') 
but it will spent more times and spaces.

* We can detect whether this node is a leaf to decide if we should perform a depth-first search (DFS) or simply store the string. Finally, we sort these strings and return the first one.
************************************************************/
