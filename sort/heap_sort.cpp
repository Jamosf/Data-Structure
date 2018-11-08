#include <iostream>

typedef int ElementType;


//ѡ�������Ƕ������˼��
void Selection_Sort(ElementType A[], int N){
    for(int i = 0; i < N;i++){
        MinPosition = ScanForMin(A,i,N-1);
        Swap(A[i],A[MinPosition]);
    }
}


//������1:nlogn������ռ�n
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
{ /* �ı����4.24��PercDown( MaxHeap H, int p )    */
  /* ��N��Ԫ�ص���������A[p]Ϊ�����Ӷѵ���Ϊ���� */
    int Parent, Child;
    ElementType X;
 
    X = A[p]; /* ȡ��������ŵ�ֵ */
    for( Parent=p; (Parent*2+1)<N; Parent=Child ) {
        Child = Parent * 2 + 1;
        if( (Child!=N-1) && (A[Child]<A[Child+1]) )
            Child++;  /* Childָ�������ӽ��Ľϴ��� */
        if( X >= A[Child] ) break; /* �ҵ��˺���λ�� */
        else  /* ����X */
            A[Parent] = A[Child];
    }
    A[Parent] = X;
}

//������2
void heap_sort(ElementType A[], int N){
    for(int i=N/2;i>=0;i--){
        PercDown(A,i,N);
    }
    for(int i=N-1;i>0;i--){
        swap(&A[0],&A[i]);
        PercDown(A,0,i);
    }
}