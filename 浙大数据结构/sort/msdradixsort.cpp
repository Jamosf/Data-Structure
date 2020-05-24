#include<iostream>

/*基数排序  主位优先*/

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

void MSD(ElementType A[], int l, int r, int d){
/*核心递归函数，对A[l]....A[r]的第d位进行排序*/
    int Di, i, j;
    Bucket B;
    PtrToNode tmp, p, list = NULL;
    if(d == 0)return;
    /*初始化每个桶为空链表*/
    for(i = 0;i < Radix;i++){
        B[i].head = B[i].tail = NULL;
    }
    /*将原始序列逆序存入初始序列*/
    for(i = l;i < r;i++){
        tmp = (PtrToNode)malloc(sizeof(struct Node));
        tmp->key = A[i];
        tmp->next = list;
        list = tmp;
    }
    /*下面是分配过程*/
    
}