#include<iostream>

typedef int ElementType;

void swap(int *a, int *b){
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}

void InsertionSort(ElementType A[], int n){
    return;
}

//找基准值pivot
ElementType Media3(ElementType A[], int left, int right){
    int center = (left + right)/2;
    if(A[left] > A[center]) swap(&A[left], &A[center]);
    if(A[left] > A[right]) swap(&A[left],&A[right]);
    if(A[center] > A[right]) swap(&A[center],&A[right]);
    swap(&A[center],&A[right-1]);
    return A[right-1];
}

void Qsort(ElementType A[], int left, int right){
    int pivot, cutoff, low, high;
    if(cutoff <= right-left){/*如果序列元素较多，进入快排 */
        pivot = Media3(A, left, right);/*选基准*/
        low = left;
        high = right - 1;
        while(1){/*将序列中比基准小的移到基准左边，大的移到右边*/
            while(A[++low] < pivot);
            while(A[--high] > pivot);
            if(low < high) swap(&A[low], &A[high]);
            else break;
        }
        swap(&A[low], &A[right-1]); /*将基准换到正确的位置*/
        Qsort(A,left,low-1);/*递归解决左边*/
        Qsort(A,low+1,right);/*递归解决右边*/
    }
    else InsertionSort(A+left,right-left+1);/*元素太少，用简单排序*/
}

void QuickSort(ElementType A[], int N){
    Qsort(A,0,N-1);
}