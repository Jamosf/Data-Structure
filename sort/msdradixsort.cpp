#include<iostream>

/*��������  ��λ����*/

/*����Ԫ�������maxDigit���ؼ��֣�����ȫ��ͬ����Radix*/
#define MaxDigit 4
#define Radix 10

typedef int ElementType;
/*ͰԪ�ؽ��*/
typedef struct Node *PtrToNode;
struct Node{
    int key;
    PtrToNode next;
};

/*Ͱͷ���*/
struct HeadNode{
    PtrToNode head, tail;
};

typedef struct HeadNode Bucket[Radix];

int GetDigit(int X, int D){
    /*Ĭ�ϴ�λD=1����λD<=MaxDigit*/
    int d, i;
    for(i = 1; i <= D; i++){
        d = X % Radix;
        X /= Radix;
    }
    return d;
}

void MSD(ElementType A[], int l, int r, int d){
/*���ĵݹ麯������A[l]....A[r]�ĵ�dλ��������*/
    int Di, i, j;
    Bucket B;
    PtrToNode tmp, p, list = NULL;
    if(d == 0)return;
    /*��ʼ��ÿ��ͰΪ������*/
    for(i = 0;i < Radix;i++){
        B[i].head = B[i].tail = NULL;
    }
    /*��ԭʼ������������ʼ����*/
    for(i = l;i < r;i++){
        tmp = (PtrToNode)malloc(sizeof(struct Node));
        tmp->key = A[i];
        tmp->next = list;
        list = tmp;
    }
    /*�����Ƿ������*/
    
}