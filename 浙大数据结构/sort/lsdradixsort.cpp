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
    /*��ԭʼ������������ʼ����*/
    for(i = 0;i < N;i++){
        tmp = (PtrToNode)malloc(sizeof(struct Node));
        tmp->key = A[i];
        tmp->next = List;
        List = tmp;
    }
    /*���濪ʼ����*/
    p = List;
    for(D = 1;D <= MaxDigit;D++){
        /*�����Ƿ������*/
        while(p){
            Di = GetDigit(p->key,D);
            /*��list���Ƴ�*/
            tmp = p; p = p->next;
            /*����B[Di]��Ͱ��ĩβ*/
            tmp->next = NULL;
            if(B[Di].head == NULL){
                B[Di].head = B[Di].tail = tmp;
            }else{
                B[Di].tail->next = tmp;
                B[Di].tail = tmp;
            }
        }
        /*�������ռ�����*/
        List = NULL;
        for(Di = Radix - 1;Di >= 0;Di--){
            if(B[Di].head){
                B[Di].tail->next = List;
                List = B[Di].head;
                B[Di].head = B[Di].tail = NULL;
            }            
        }
    }
    /*��list���ݵ���A*/
    for(i = 0; i < N;i++){
        tmp = List;
        A[i] = tmp->key;
        List = List->next;
        free(tmp);
    }
}