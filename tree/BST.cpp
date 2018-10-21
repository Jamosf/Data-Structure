#include<iostream>

Position Find(ElementType x, BinTree BST){
    if(!BST) return NULL;
    if(x > BST->Data) return Find(x,BST->Right);
    else if(x < BST->Data) return Find(x,BST->Left);
    else
        return BST;
}

/*�����ٶ�ȡ�������ĸ߶�*/
Position IterFind(ElementType x, BinTree BST){
    while(BST){
        if(x > BST->Data){
            BST = BST->Right;
        }else if(x < BST->Data){
            BST = BST->Left;
        }else{
            return BST;
        }
    }
    return NULL;
}

Position FindMin( BinTree BST){
    if(!BST)return NULL;
    else if(!BST->Left)return BST;
    else return FindMin(BST->Left);
}

Position FindMax(BinTree BST){
    while(BST){
        if(BST->Right) BST = BST->Right;
    }
    return BST;
}

BinTree Insert(ElementType x, BinTree BST){
    if(!BST){
        BST = malloc(sizeof(struct TreeNode));
        BST->Data = x;
        BST->Left = BST->Right = NULL;
    }
    else{
        if(x > BST->Data){
            BST->Right = Insert(x,BST->Right);
        }else if(x < BST->Data){
            BST->Left = Insert(x,BST->Left);
        }
    }
    return NULL;
}

Position Delete(BinTree BST,Element X){
    Position tmp;
    if(!BST){
        printf("Ҫɾ����Ԫ��δ�ҵ�");
    }
    else{
        if(X < BST->Data) BST->Left = Delete(BST->Left,X);
        else if(X > BST->Data) BST->Right = Delete(BST->Right, X);
        else{/*BSTҪɾ���Ľڵ�*/
            /*�����ɾ���Ľڵ������������ӽڵ�*/
            if(BST->Left && BST->Right){
                /*��������С���ɾ�����*/
                tmp = FindMin(BST->Right);
                BST->Data = tmp.Data;
                /*����������ɾ����СԪ��*/
                BST->Right = Delete(BST->Right,BST->Data);
            }
            else{
                tmp = BST;
                if(!BST->Left) BST = BST->Right;/*ֻ�����ӻ���û�ж��ӽڵ�*/ 
                else if(!BST->Right) BST = BST->Left;/*ֻ���Һ���*/
                free(tmp);
            }
        }
    }
    return BST;
}