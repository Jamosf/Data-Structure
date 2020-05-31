#include <stdio.h>
#include <stdlib.h>

int max(int a,int b){
	return a > b ? a : b;
}

int rob(int* nums, int numsSize) {
	if(numsSize == 1){
		return nums[0];
	}
	if(numsSize == 2){
		return max(nums[0],nums[1]);
	}
	int *maxNums = (int *)malloc(sizeof(int)*numsSize);	
	maxNums[0] = nums[0];
	maxNums[1] = max(nums[0],nums[1]); 
	for(int i=2;i<numsSize;i++){
		maxNums[i] = max(maxNums[i-1],nums[i]+maxNums[i-2]);
	}
	return maxNums[numsSize-1];
}

int main(){
	int a[5] = {2,7,9,3,1};
	printf("the result is: %d",rob(a,sizeof(a)/sizeof(int)));
	return 0;
}
