#include <iostream>

typedef int ElementType;


//选择排序是堆排序的思想
void Selection_Sort(ElementType A[], int N){
    for(int i = 0; i < N;i++){
        MinPosition = ScanForMin(A,i,N-1);
        Swap(A[i],A[MinPosition]);
    }
}


//堆排序1:nlogn，额外空间n
void heap_sort(ElementType A[], int N){
    BuildHeap(A);
    ElementType *tmp = (ElementType*)malloc(sizeof(ElementType)*N);
    for(int i = 0; i <N;i++){
        tmp[i] = DeleteMin(A);
    }
    for(int i=0;i<N;i++){
        A[i] = tmp[i];
    }
}

void swap(int *a, int *b){
    int tmp;
    tmp = *a;
    *a = *b;
    *b = tmp;
}

void PercDown( ElementType A[], int p, int N )
{ /* 改编代码4.24的PercDown( MaxHeap H, int p )    */
  /* 将N个元素的数组中以A[p]为根的子堆调整为最大堆 */
    int Parent, Child;
    ElementType X;
 
    X = A[p]; /* 取出根结点存放的值 */
    for( Parent=p; (Parent*2+1)<N; Parent=Child ) {
        Child = Parent * 2 + 1;
        if( (Child!=N-1) && (A[Child]<A[Child+1]) )
            Child++;  /* Child指向左右子结点的较大者 */
        if( X >= A[Child] ) break; /* 找到了合适位置 */
        else  /* 下滤X */
            A[Parent] = A[Child];
    }
    A[Parent] = X;
}

//堆排序2
void heap_sort(ElementType A[], int N){
    for(int i=N/2;i>=0;i--){
        PercDown(A,i,N);
    }
    for(int i=N-1;i>0;i--){
        swap(&A[0],&A[i]);
        PercDown(A,0,i);
    }
}