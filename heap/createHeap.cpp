#include<stdio.h>
#include<stdlib.h>

typedef int ElementType;
typedef struct HeapStruct *MaxHeap;
struct HeapStruct{
    ElementType *Elements;
    int Size;
    int Capacity;
};

MaxHeap Create(int MaxSize){
    MaxHeap h = (MaxHeap)malloc(sizeof(struct HeapStruct));
    h->Elements = (ElementType*)malloc((MaxSize+1)*sizeof(ElementType));
    h->Size = 0;
    h->Capacity = MaxSize;
    h->Elements[0] = 65535;
}

bool IsEmpty(MaxHeap H){
    return H->Size == 0;
}

bool IsFull(MaxHeap H){
    if(IsEmpty(H))return true;
    return H->Size == H->Capacity;
}

void Insert(MaxHeap H,ElementType item){
    int i;
    if(IsFull(H)){
        printf("��������");
        return;
    }
    i = ++H->Size;
    for(;H->Elements[i/2] < item;i/=2){
        H->Elements[i] = H->Elements[i/2];
    }
    H->Elements[i] = item;
}

ElementType DeleteMax(MaxHeap H){
    int parent, child;
    ElementType maxItem,temp;
    if(IsEmpty(H)){
        printf("������Ϊ��");
        return;
    }
    maxItem = H->Elements[1];
    temp = H->Elements[H->Size--];
    for(parent=1;parent*2 <= H->Size;parent=child){
        child = parent*2;
        if((child!=H->Size)&&(H->Elements[child] < H->Elements[child+1])) child++;
        if(temp >= H->Elements[child])break;
        else
            H->Elements[parent] = H->Elements[child];
    }
    H->Elements[parent] = temp;
    return maxItem;
}

/*----------- �������� -----------*/
void PercDown( MaxHeap H, int p )
{ /* ���ˣ���H����H->Data[p]Ϊ�����Ӷѵ���Ϊ���� */
    int Parent, Child;
    ElementType X;
 
    X = H->Elements[p]; /* ȡ��������ŵ�ֵ */
    for( Parent=p; Parent*2<=H->Size; Parent=Child ) {
        Child = Parent * 2;
        if( (Child!=H->Size) && (H->Elements[Child]<H->Elements[Child+1]) )
            Child++;  /* Childָ�������ӽ��Ľϴ��� */
        if( X >= H->Elements[Child] ) break; /* �ҵ��˺���λ�� */
        else  /* ����X */
            H->Elements[Parent] = H->Elements[Child];
    }
    H->Elements[Parent] = X;
}
 
void BuildHeap( MaxHeap H )
{ /* ����H->Data[]�е�Ԫ�أ�ʹ�������ѵ�������  */
  /* �����������H->Size��Ԫ���Ѿ�����H->Data[]�� */
 
    int i;
 
    /* �����һ�����ĸ��ڵ㿪ʼ���������1 */
    for( i = H->Size/2; i>0; i-- )
        PercDown( H, i );
}
