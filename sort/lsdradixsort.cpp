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

void LSDRadixSort(ElementType A[], int N){
    /*�������򣬴�λ����*/
    int D, Di, i;
    Bucket B;
    PtrToNode tmp, p, List = NULL;
    /*��ʼ��ÿ��ͰΪ������*/
    for(i = 0;i < Radix;i++){
        B[i].head = B[i].tail = NULL;
    }
}