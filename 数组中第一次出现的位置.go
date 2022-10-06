package main

func main() {
	
}

func firstPostsion(nums []int, k){
	l, r := 0, len(nums)-1
    mid := l +(r-l)/2
    for l <r{
        tmp := nums[mid]
        if nums[mid]>tmp{
            l = mid
        }else if nums[mid] < tmp{
            r = mid -1
        }else{
            return mid
        }
    }

    return -1
}