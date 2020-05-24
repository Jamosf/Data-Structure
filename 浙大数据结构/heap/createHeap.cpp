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
        printf("最大堆已满");
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
        printf("最大堆已为空");
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

/*----------- 建造最大堆 -----------*/
void PercDown( MaxHeap H, int p )
{ /* 下滤：将H中以H->Data[p]为根的子堆调整为最大堆 */
    int Parent, Child;
    ElementType X;
 
    X = H->Elements[p]; /* 取出根结点存放的值 */
    for( Parent=p; Parent*2<=H->Size; Parent=Child ) {
        Child = Parent * 2;
        if( (Child!=H->Size) && (H->Elements[Child]<H->Elements[Child+1]) )
            Child++;  /* Child指向左右子结点的较大者 */
        if( X >= H->Elements[Child] ) break; /* 找到了合适位置 */
        else  /* 下滤X */
            H->Elements[Parent] = H->Elements[Child];
    }
    H->Elements[Parent] = X;
}
 
void BuildHeap( MaxHeap H )
{ /* 调整H->Data[]中的元素，使满足最大堆的有序性  */
  /* 这里假设所有H->Size个元素已经存在H->Data[]中 */
 
    int i;
 
    /* 从最后一个结点的父节点开始，到根结点1 */
    for( i = H->Size/2; i>0; i-- )
        PercDown( H, i );
}
