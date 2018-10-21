#include<iostream>

Position Find(ElementType x, BinTree BST){
    if(!BST) return NULL;
    if(x > BST->Data) return Find(x,BST->Right);
    else if(x < BST->Data) return Find(x,BST->Left);
    else
        return BST;
}

/*查找速度取决于树的高度*/
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
        printf("要删除的元素未找到");
    }
    else{
        if(X < BST->Data) BST->Left = Delete(BST->Left,X);
        else if(X > BST->Data) BST->Right = Delete(BST->Right, X);
        else{/*BST要删除的节点*/
            /*如果被删除的节点有左右两个子节点*/
            if(BST->Left && BST->Right){
                /*右子树最小填充删除结点*/
                tmp = FindMin(BST->Right);
                BST->Data = tmp.Data;
                /*从右子树中删除最小元素*/
                BST->Right = Delete(BST->Right,BST->Data);
            }
            else{
                tmp = BST;
                if(!BST->Left) BST = BST->Right;/*只有左孩子或者没有儿子节点*/ 
                else if(!BST->Right) BST = BST->Left;/*只有右孩子*/
                free(tmp);
            }
        }
    }
    return BST;
}