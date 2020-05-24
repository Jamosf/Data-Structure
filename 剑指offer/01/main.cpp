/*
剑指offer第一题
题目描述：
在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/
#include <stdio.h>


class Solution {
public:
    bool Find(int target, vector<vector<int> > array) {
        for(int i = 0; i < array.size();i++){
            if (array[i][i] == target){
                return true;
            }
            if (array[i][i] < target && (i+1 < array.size()) && array[i+1][i+1] > target){
                for (int p = i; p <= i+1; p++){
                    for (int q = 0;q < array.size();q++){
                        if (array[p][q] == target){
                            return true;
                        }
                    }
                }
                return false;
            }
        }
        return false;
    }
};