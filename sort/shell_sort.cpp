#include <iostream>

typedef int ElementType;


//shell的增量序列，取中间数
//时间复杂度：n^2
void shell_sort(ElementType A[], int N){
    int tmp, pos;
    for(int D = N/2;D > 0;D/=2){
        for(int p = D;p < N;p++){
            tmp = A[p];
            for(int pos = p;pos > 0 && A[pos-1] > tmp;pos--){
                A[pos] = A[pos-1];
            }
            A[pos] = tmp;
        }
    }
}


//sedgewick增量序列
void shell_sort(ElementType A[], int N){
    int Si, tmp, pos;
    int sedgewick[] = {929, 505, 209, 109, 41, 19, 5, 1, 0};
    for(Si = 0;sedgewick[Si]>=N;Si++);
    for(int D = sedgewick[Si];D > 0;D = sedgewick[++Si]){
        for(int p = D;p < N;p++){
            tmp = A[p];
            for(int pos = p;pos > 0 && A[pos-1] > tmp;pos--){
                A[pos] = A[pos-1];
            }
            A[pos] = tmp;
        }
    }
}