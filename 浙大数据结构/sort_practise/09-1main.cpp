#include <iostream>
#define MaxNum 100000
using namespace std;
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

void Insertion_Sort(ElementType A[], int N){
    int tmp, pos;
    for(int p = 1;p < N;p++){
        tmp = A[p];
        for(pos = p;pos > 0 && A[pos-1] > tmp;pos--){
            A[pos] = A[pos-1];
        }
        A[pos] = tmp;
    }
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

//sedgewick��������
void shell_sort(ElementType A[], int N){
    int Si, tmp, pos;
    int sedgewick[] = {929, 505, 209, 109, 41, 19, 5, 1, 0};
    for(Si = 0;sedgewick[Si]>=N;Si++);
    for(int D = sedgewick[Si];D > 0;D = sedgewick[++Si]){
        for(int p = D;p < N;p++){
            tmp = A[p];
            for(pos = p;pos > 0 && A[pos-1] > tmp;pos--){
                A[pos] = A[pos-1];
            }
            A[pos] = tmp;
        }
    }
}

void shell_sort1(ElementType A[], int N){
    int tmp, pos;
    for(int D = N/2;D > 0;D/=2){
        for(int p = D;p < N;p++){
            tmp = A[p];
            for(pos = p;pos > 0 && A[pos-1] > tmp;pos--){
                A[pos] = A[pos-1];
            }
            A[pos] = tmp;
        }
    }
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
    cutoff = 100;
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
    else Insertion_Sort(A+left,right-left+1);/*Ԫ��̫�٣��ü�����*/
}

void QuickSort(ElementType A[], int N){
    Qsort(A,0,N-1);
}

int main(){
    int n;
    cin >> n;
    int A[MaxNum];
    for(int i = 0;i < n;i++){
        cin >> A[i];
    }
    //bubble_sort(A, n);
    //Insertion_Sort(A, n);
    //heap_sort(A, n);
    //Merge_sort(A, n);
    //shell_sort(A, n);
    QuickSort(A, n);
    cout << A[0];
    for(int i = 1;i < n;i++){
        cout <<" "<< A[i];
    }
    //system("pause");
}