#include<iostream>

using namespace std;

typedef struct LNode *Ptr2Node;
struct LNode{
    LNode *pre;
    int preAddress;
    int data;
    int nextAddress;
    LNode *next;
};

typedef Ptr2Node List;

void InsertAfter(int pre, int x, int next, List *l){
    List tmp;
    tmp = (List)malloc(sizeof(struct LNode));
    tmp->data = x;
    tmp->nextAddress = next;
    tmp->next = NULL;
    tmp->next = (*l)->next;
    (*l)->next = tmp;
}

void InsertBefore(int pre, int x, int next, List *l){
    List tmp;
    tmp = (List)malloc(sizeof(struct LNode));
    tmp->data = x;
    tmp->nextAddress = next;
    tmp->next = NULL;
    tmp->next = (*l);
    (*l)->pre = tmp;
}

List ReadList(int address,int n){
    List p,rear,tmp;
    p = (List)malloc(sizeof(struct LNode));
    p->pre = NULL;
    p->next = NULL;
    p->nextAddress = address;
    rear = p;
    while(n--){
        List l = (List)malloc(sizeof(struct LNode));
        cin >> l->preAddress >> l->data >> l->nextAddress;
        while(rear){
            if(rear->preAddress == l->nextAddress){
                InsertBefore(l->preAddress, l->data,l->nextAddress,&rear);
            }else if(rear->nextAddress == l->preAddress){
                InsertAfter(l->preAddress, l->data,l->nextAddress,&rear);
            }
            rear = rear->next;
        }
        InsertAfter(l->preAddress, l->data,l->nextAddress,&rear);
    }
    tmp = p;
    p=p->next;
    free(tmp);
    return p;
}

void ReverseList(List *l,int n, int k){
    List rear = *l;
    int loop = n/k;
    while(loop--){
        for(int i = 0; i < k;i++){
            rear->next->next = rear;
        }
    }
}

int main(){
    int address,n,k;
    cin >> address >> n >> k;
    List pp;
    pp = ReadList(address,n);

}