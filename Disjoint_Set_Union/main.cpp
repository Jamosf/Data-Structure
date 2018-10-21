#include <stdio.h>
#define MaxNum 100
typedef int ElementType;


typedef struct{
    ElementType element;
    int parent;
}SetType;


/*并查集的查找操作*/
int Find(SetType S[],ElementType X){
    int i;
    for(i=0;(i < MaxNum) && (S[i].element != X);i++) ;
    if(i >= MaxNum) return -1;
    for(;S[i].parent >= 0;i = S[i].parent) ;
    return i;
}

/*并查集的合并操作*/
void Union(SetType S[], ElementType X1,ElementType X2){
    int Root1, Root2;
    Root1 = Find(S,X1);
    Root2 = Find(S,X2);
    if(Root1 != Root2) S[Root2].parent = Root1;
}

void Union( SetType S, SetName Root1, SetName Root2 )
{ /* 这里默认Root1和Root2是不同集合的根结点 */
    /* 保证小集合并入大集合 */
    if ( S[Root2] < S[Root1] ) { /* 如果集合2比较大 */
        S[Root2] += S[Root1];     /* 集合1并入集合2  */
        S[Root1] = Root2;
    }
    else {                         /* 如果集合1比较大 */
        S[Root1] += S[Root2];     /* 集合2并入集合1  */
        S[Root2] = Root1;
    }
}
 
SetName Find( SetType S, ElementType X )
{ /* 默认集合元素全部初始化为-1 */
    if ( S[X] < 0 ) /* 找到集合的根 */
        return X;
    else
        return S[X] = Find( S, S[X] ); /* 路径压缩 */
}