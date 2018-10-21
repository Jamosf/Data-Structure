#include<stdio.h>

#define MaxTree 10
#define ElementType char
#define Tree int
#define Null -1

struct TreeNode{
    ElementType Element;
    Tree Left;
    Tree Right;
}T1[MaxTree], T2[MaxTree];

Tree BuildTree(struct TreeNode T[]){
    scanf("%d\n",&N);
    if(N){
        for(int i=0; i<N;i++){
            check[i] = 0;
        }
        for(int i=0;i<N;i++){
            scanf("%c %c %c\n",&T[i].Element,&cl,&cr);
            if(cl != '-'){
                T[i].left = cl-'0';
                check[T[i].left] = 1;
            }
            else{
                T[i].left = Null;
            }
            if(cr != '-'){
                T[i].right = cr-'0';
                check[T[i].right] = 1;
            }
            else{
                T[i].right = Null;
            }
        }
        for(int i=0;i<N;i++){
            if(!check[i])break;
        }
        Root = i;
    }
    return Root;
}


int main(){
    Tree R1,R2;

}