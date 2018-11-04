#include<iostream>

/*基数排序  次位优先*/

/*假设元素最多有maxDigit个关键字，基数全是同样的Radix*/
#define MaxDigit 4
#define Radix 10

typedef int ElementType;
/*桶元素结点*/
typedef struct Node *PtrToNode;
struct Node{
    int key;
    PtrToNode next;
};

/*桶头结点*/
struct HeadNode{
    PtrToNode head, tail;
};

typedef struct HeadNode Bucket[Radix];

int GetDigit(int X, int D){
    /*默认次位D=1，主位D<=MaxDigit*/
    int d, i;
    for(i = 1; i <= D; i++){
        d = X % Radix;
        X /= Radix;
    }
    return d;
}

void LSDRadixSort(ElementType A[], int N){
    /*基数排序，次位优先*/
    int D, Di, i;
    Bucket B;
    PtrToNode tmp, p, List = NULL;
    /*初始化每个桶为空链表*/
    for(i = 0;i < Radix;i++){
        B[i].head = B[i].tail = NULL;
    }
}