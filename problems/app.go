package problems

import (
	"fmt"
)



func Start() error{

	return rotate()

}

func sortedAndRotatedCheck() error{
	nums:=[]int{6,10,6}
	variations:=0
	for i:=0; i<len(nums);i++{
		if(nums[i-1]>nums[i]){
			variations++
		}
	}

	if(nums[len(nums)-1]>nums[0]){
		variations++
	}
	fmt.Println(variations)
	return nil
}


func handleModulo(num int, cap int,n int)int{
	if(num-cap<0){
		return (num-cap)%n+n
	}
	return (num-cap)%n
}

func rotate() error{
	nums:=[]int{1,2,3,4,5,6,7}
	nums1:=make([]int,0)
	nums1=append(nums1, nums...)

	fmt.Println(nums1)

	for i:=0;i<len(nums);i++{
		
		nums[(i+3)%len(nums)]=nums1[i]
	}
	fmt.Println(nums)
	for i:=0;i<3;i++{
		temp:=nums[i]
		nums[i]=nums[3-i]
		nums[3-i]=temp
	}

	for i:=3;i<len(nums);i++{
		temp:=nums[i]
		nums[i]=nums[len(nums)-i-1]
		nums[len(nums)-i-1]=temp
	}

	fmt.Println(nums)
	return nil
}