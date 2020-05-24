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
    /*将原始序列逆序存入初始序列*/
    for(i = 0;i < N;i++){
        tmp = (PtrToNode)malloc(sizeof(struct Node));
        tmp->key = A[i];
        tmp->next = List;
        List = tmp;
    }
    /*下面开始排序*/
    p = List;
    for(D = 1;D <= MaxDigit;D++){
        /*下面是分配过程*/
        while(p){
            Di = GetDigit(p->key,D);
            /*从list中移除*/
            tmp = p; p = p->next;
            /*插入B[Di]号桶的末尾*/
            tmp->next = NULL;
            if(B[Di].head == NULL){
                B[Di].head = B[Di].tail = tmp;
            }else{
                B[Di].tail->next = tmp;
                B[Di].tail = tmp;
            }
        }
        /*下面是收集过程*/
        List = NULL;
        for(Di = Radix - 1;Di >= 0;Di--){
            if(B[Di].head){
                B[Di].tail->next = List;
                List = B[Di].head;
                B[Di].head = B[Di].tail = NULL;
            }            
        }
    }
    /*将list内容倒入A*/
    for(i = 0; i < N;i++){
        tmp = List;
        A[i] = tmp->key;
        List = List->next;
        free(tmp);
    }
}