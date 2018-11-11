#include <iostream>

typedef int ElementType;

//�鲢�������ȶ����㷨��ȱ������Ҫ����Ŀռ䣬����������

//�ϲ��㷨
void merge(ElementType A[], ElementType TmpA[], int L, int R, int RightEnd){
    int LeftEnd = R -1;
    int Tmp = L;
    int length = RightEnd - L + 1;
    while(L <= LeftEnd && R <= RightEnd){
        if(A[L] < A[R]){
            TmpA[Tmp++] = A[L++];
        }else{
            TmpA[Tmp++] = A[R++];
        }
    }
    while(L <= LeftEnd){
        TmpA[Tmp++] = A[L++];
    }
    while(R <= RightEnd){
        TmpA[Tmp++] = A[R++];
    }
    for(int i = 0; i < length;i++,RightEnd--){
        A[RightEnd] = TmpA[RightEnd];
    }
}

//�ݹ��ʵ�ֹ鲢����,ʱ�临�Ӷ�nlogn
void Msort(ElementType A[], ElementType TmpA[], int L, int RightEnd){
    int center;
    if(L < RightEnd){
        center = (L+RightEnd)/2;
        Msort(A,TmpA,L,center);
        Msort(A,TmpA,center+1,RightEnd);
        merge(A,TmpA,L,center+1,RightEnd);
    }
}

//����ӿ�
void Merge_sort(ElementType A[], int N){
    ElementType *TmpA;
    TmpA = (ElementType*)malloc(N * sizeof(ElementType));
    if(TmpA != NULL){
        Msort(A,TmpA,0,N-1);
        free(TmpA);
    }
}

//�ϲ��㷨
void Merge1(ElementType A[], ElementType TmpA[], int L, int R, int RightEnd){
    int LeftEnd = R -1;
    int Tmp = L;
    while(L <= LeftEnd && R <= RightEnd){
        if(A[L] < A[R]){
            TmpA[Tmp++] = A[L++];
        }else{
            TmpA[Tmp++] = A[R++];
        }
    }
    while(L <= LeftEnd){
        TmpA[Tmp++] = A[L++];
    }
    while(R <= RightEnd){
        TmpA[Tmp++] = A[R++];
    }
}

//�ǵݹ�ʵ�ֹ鲢����
void merge_pass(ElementType A[], ElementType TmpA[], int N, int length){
    int i;
    for(i = 0; i <= N-2*length;i += 2*length){
        Merge1(A,TmpA,i,i+length,i+2*length-1);
    }
    if(i+length < N){
        Merge1(A,TmpA,i,i+length,N-1);
    }
    else{
        for(int j = i;j < N;j++)TmpA[j] = A[j];
    }
}

//����ӿ�
void Merge_sort1(ElementType A[], int N){
    int length = 1; //��ʼ�������еĳ���
    ElementType *TmpA;
    TmpA = (ElementType*)malloc(N * sizeof(ElementType));
    if(TmpA != NULL){
        while(length < N){
            merge_pass(A,TmpA,N,length);
            length *= 2;
            merge_pass(TmpA,A,N,length);
            length *= 2;
        }
        free(TmpA);
    }
}