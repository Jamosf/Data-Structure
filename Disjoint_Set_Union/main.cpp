#include <stdio.h>
#define MaxNum 100
typedef int ElementType;


typedef struct{
    ElementType element;
    int parent;
}SetType;


/*���鼯�Ĳ��Ҳ���*/
int Find(SetType S[],ElementType X){
    int i;
    for(i=0;(i < MaxNum) && (S[i].element != X);i++) ;
    if(i >= MaxNum) return -1;
    for(;S[i].parent >= 0;i = S[i].parent) ;
    return i;
}

/*���鼯�ĺϲ�����*/
void Union(SetType S[], ElementType X1,ElementType X2){
    int Root1, Root2;
    Root1 = Find(S,X1);
    Root2 = Find(S,X2);
    if(Root1 != Root2) S[Root2].parent = Root1;
}

void Union( SetType S, SetName Root1, SetName Root2 )
{ /* ����Ĭ��Root1��Root2�ǲ�ͬ���ϵĸ���� */
    /* ��֤С���ϲ���󼯺� */
    if ( S[Root2] < S[Root1] ) { /* �������2�Ƚϴ� */
        S[Root2] += S[Root1];     /* ����1���뼯��2  */
        S[Root1] = Root2;
    }
    else {                         /* �������1�Ƚϴ� */
        S[Root1] += S[Root2];     /* ����2���뼯��1  */
        S[Root2] = Root1;
    }
}
 
SetName Find( SetType S, ElementType X )
{ /* Ĭ�ϼ���Ԫ��ȫ����ʼ��Ϊ-1 */
    if ( S[X] < 0 ) /* �ҵ����ϵĸ� */
        return X;
    else
        return S[X] = Find( S, S[X] ); /* ·��ѹ�� */
}