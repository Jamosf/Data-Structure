#include <iostream>

typedef int ElementType;

void swap(int *a, int *b){
    int tmp;
    tmp = *a;
    *a = *b;
    *b = tmp;
}

//ʱ�临�Ӷ����O(n)�����O(n^2)
//�������ڵ��������ݽṹ������
void bubble_sort(ElementType A[], int N){
    int flag;
    for(int p = N-1;p>=0;p--){
        flag = 0;
        for(int i = 0; i < p;i++){
            if(A[i] > A[i+1]){
                swap(&A[i],&A[i+1]);
                flag = 1;
            }
        }
        if(flag == 0)break;
    }
}