package main
import "fmt"

func main() {
	nums := []int{1,2,3,1}
	containsDuplicate(nums)
}

func containsDuplicate(nums []int) bool {
	// need to allocate mem
    mymap := make(map[int]int)
    
    for _, val := range nums {
        if(mymap[val] >=1) {
            mymap[val] = mymap[val]+1
			return true;
            
        } else {
            mymap[val] = 1
        }
    }
    
	return false;
    //fmt.Println(mymap);
}
