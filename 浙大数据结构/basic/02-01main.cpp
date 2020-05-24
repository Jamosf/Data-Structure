#include <stdio.h>
#include <stdlib.h>

typedef int ElementType;
typedef struct Node *PtrToNode;
struct Node {
    ElementType Data;
    PtrToNode   Next;
};
typedef PtrToNode List;

List Read(); /* 细节在此不表 */
void Print( List L ); /* 细节在此不表；空链表将输出NULL */

List Merge( List L1, List L2 );

int main()
{
    List L1, L2, L;
    L1 = Read();
    L2 = Read();
    L = Merge(L1, L2);
    Print(L);
    Print(L1);
    Print(L2);
    system("pause");
    return 0;
}

List Read(){
    int n;
    scanf("%d",&n);
    List L = (List)malloc(sizeof(struct Node));
    scanf("%d",&L->Data);
    List tail = L;
    List s;
    for(int i=0; i < n-1;i++){
        s = (List)malloc(sizeof(Node));
        scanf("%d",&s->Data);
        tail->Next = s;
        tail = s;
    }
    tail->Next = NULL;
    List H = (List)malloc(sizeof(struct Node));
    H->Next = L;
    return H;
}

void Print( List L ){
    if(L == NULL || L->Next == NULL){
        printf("NULL");
    }
    L = L->Next;
    if(L != NULL){
        printf("%d",L->Data);
        L = L->Next;
    }
    while(L != NULL){
        printf(" %d",L->Data);
        L = L->Next;
    }
    printf("\n");
}

/* 你的代码将被嵌在这里 */
List Merge( List L1, List L2 ){
    List L = (List)malloc(sizeof(struct Node));
    List tail = NULL;
    List p = L1->Next;
    List q = L2->Next;
    tail = L;    
    while(p != NULL && q != NULL){
        if(p->Data > q->Data){
            tail->Next = q;
            q = q->Next;
        }else{
            tail->Next = p;
            p = p->Next;
        }
        tail = tail->Next;
    }
    if(p) tail->Next = p;
    if(q) tail->Next = q;
    
    L1->Next = NULL;
    L2->Next = NULL;
    return L;
}