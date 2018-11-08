#include <iostream>

typedef int ElementType;


//最好情况：O(N)
//最坏情况：O(N^2)
void Insertion_sort(ElementType A[], int N){
    int tmp, pos;
    for(int p = 1;p < N;p++){
        tmp = A[p];
        for(int pos = p;pos > 0 && A[pos-1] > tmp;pos--){
            A[pos] = A[pos-1];
        }
        A[pos] = tmp;
    }
}