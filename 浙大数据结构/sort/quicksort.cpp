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

//�һ�׼ֵpivot
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
    if(cutoff <= right-left){/*�������Ԫ�ؽ϶࣬������� */
        pivot = Media3(A, left, right);/*ѡ��׼*/
        low = left;
        high = right - 1;
        while(1){/*�������бȻ�׼С���Ƶ���׼��ߣ�����Ƶ��ұ�*/
            while(A[++low] < pivot);
            while(A[--high] > pivot);
            if(low < high) swap(&A[low], &A[high]);
            else break;
        }
        swap(&A[low], &A[right-1]); /*����׼������ȷ��λ��*/
        Qsort(A,left,low-1);/*�ݹ������*/
        Qsort(A,low+1,right);/*�ݹ����ұ�*/
    }
    else InsertionSort(A+left,right-left+1);/*Ԫ��̫�٣��ü�����*/
}

void QuickSort(ElementType A[], int N){
    Qsort(A,0,N-1);
}