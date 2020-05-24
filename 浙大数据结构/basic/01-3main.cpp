#include <stdio.h>
#include <stdlib.h>

#define MAXSIZE 10
#define NotFound 0
typedef int ElementType;

typedef int Position;
typedef struct LNode *List;
struct LNode {
    ElementType Data[MAXSIZE];
    Position Last; /* �������Ա������һ��Ԫ�ص�λ�� */
};

List ReadInput(); /* ����ʵ�֣�ϸ�ڲ���Ԫ�ش��±�1��ʼ�洢 */
Position BinarySearch( List L, ElementType X );

int main()
{
    List L;
    ElementType X;
    Position P;

    L = ReadInput();
    scanf("%d", &X);
    P = BinarySearch( L, X );
    printf("%d\n", P);
    system("pause");
    return 0;
}

/* ��Ĵ��뽫��Ƕ������ */
List ReadInput(){
    int n;
    scanf("%d",&n);
    List l = (LNode*)malloc(sizeof(LNode));
    for(int i = 1; i <= n;i++){
        scanf("%d",&l->Data[i]);
    }
    l->Last = n;
    return l;
}

Position BinarySearch( List L, ElementType X ){
    int left = 1;
    int right = L->Last;
    int mid;
    while(left <= right){
        mid = (left + right)/2;
        if(L->Data[mid] == X){
            return mid;
        }else if(L->Data[mid] > X){
            right = mid - 1;
        }else{
            left = mid + 1;
        }
    }
    return NotFound;
}

