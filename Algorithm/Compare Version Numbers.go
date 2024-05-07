/************************************************************
*Initially, parse the two strings into arrays by splitting each string at each '.' and remove leading zeros until you reach another '.' or a non-zero number.
*Pad the two arrays to the same length with '0's and then compare the two arrays element by element.
time complexity O(n)
space complexity O(n)
************************************************************/
func compareVersion(version1 string, version2 string) int {
    s1 := StringToArray(version1)
    s2 := StringToArray(version2)
    s1Length := len(s1)
    s2Length := len(s2)
    if s1Length > s2Length{
     for i := s2Length; i < s1Length; i++{
        s2 = append(s2,0)
     }
    }else if s2Length > s1Length{
       for i := s1Length; i < s2Length; i++{
        s1 = append(s1,0)
     } 
     s1Length = s2Length
    }
    fmt.Print(s1,s2)
    for i := 0; i < s1Length; i++{
        if s1[i] > s2[i]{
                return 1
            }else if s1[i] < s2[i]{
                return -1
            }
        }
    return 0

}
func StringToArray(s string) []int{
    flag := 0
    tmp := 0
    var s1 []int
    for _,v := range s{
        switch v {
            case '0':
                if flag == 0{
                    continue
                }else{
                    tmp = tmp * 10
                }
            case '.':
                s1 = append(s1,tmp)
                flag = 0
                tmp = 0
            default :
                flag = 1
                tmp = tmp*10 + int(v - '0')              
        }
    }
    s1 = append(s1,tmp)
    return s1
}
