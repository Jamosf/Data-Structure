#include<iostream>
using namespace std;

#define MaxNode 10
#define null -1

struct Tree{
    char element;
    int right;
    int left;
};

int BuildTree(Tree *t,int n){
    char lc, rc;
    int root = -1;
    int check[MaxNode] = {0};
    for(int i = 0; i < n; i++){
        cin >> t[i].element >> lc >> rc;
        if(lc != '-'){
            t[i].left = lc - '0';
            check[t[i].left] = 1;
        }else{
            t[i].left = -1;
        }
        if(rc != '-'){
            t[i].right = rc - '0';
             check[t[i].right] = 1;
        }else{
            t[i].right = -1;
        }
    }
    for(int i = 0; i < n;i++){
        if(!check[i]){
            root = i;
        }
    }
    return root;
}

int lsomorphic(int root1,int root2,Tree *t1, Tree *t2){
    if(root1 == -1 && root2 == -1){
        return 1;
    }
    if(root1 != -1 && root2 == -1){
        return 0;
    }
    if(root1 == -1 && root2 != -1){
        return 0;
    }
    if(t1[root1].element != t2[root2].element){
        return 0;
    }
    if(t1[root1].left == -1 && t2[root2].left == -1){
        return lsomorphic(t1[root1].right,t2[root2].right,t1,t2);
    }
    if(t1[root1].left != -1 && t2[root2].left != -1 && t1[t1[root1].left].element == t2[t2[root2].left].element){
        return lsomorphic(t1[root1].left,t2[root2].left,t1,t2) && lsomorphic(t1[root1].right,t2[root2].right,t1,t2);
    }else{
        return lsomorphic(t1[root1].left,t2[root2].right,t1,t2) && lsomorphic(t1[root1].right,t2[root2].left,t1,t2);
    }
}

int main(){
    int m, n;
    cin >> m;
    Tree *t1 = new Tree[m];
    int root1 = BuildTree(t1,m);
    cin >> n;
    Tree *t2 = new Tree[n];
    int root2 = BuildTree(t2,n);
    if(lsomorphic(root1,root2,t1,t2) == 1){
        cout << "Yes";
    }else{
        cout << "No";
    }
    return 0;
}