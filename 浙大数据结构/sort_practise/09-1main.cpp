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

//时间复杂度最好O(n)，最坏是O(n^2)
//可以用于单链表数据结构的排序
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

//合并算法
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

//递归的实现归并排序,时间复杂度nlogn
void Msort(ElementType A[], ElementType TmpA[], int L, int RightEnd){
    int center;
    if(L < RightEnd){
        center = (L+RightEnd)/2;
        Msort(A,TmpA,L,center);
        Msort(A,TmpA,center+1,RightEnd);
        merge(A,TmpA,L,center+1,RightEnd);
    }
}

//对外接口
void Merge_sort(ElementType A[], int N){
    ElementType *TmpA;
    TmpA = (ElementType*)malloc(N * sizeof(ElementType));
    if(TmpA != NULL){
        Msort(A,TmpA,0,N-1);
        free(TmpA);
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
    cutoff = 100;
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
    else Insertion_Sort(A+left,right-left+1);/*元素太少，用简单排序*/
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