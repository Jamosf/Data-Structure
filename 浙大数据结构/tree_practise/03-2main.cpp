#include<iostream>
#include<queue>
#include<stdlib.h>
using namespace std;

#define MaxNode 10

struct Tree{
    int element;
    char right;
    char left;
};

bool visited[MaxNode] = {false};

int getRoot(Tree *t, int nodes){
    int sum = 0;
    for(int i = 0; i < nodes;i++){
        if(t[i].left != '-'){
            sum += t[i].left - '0';
        }
        if(t[i].right != '-'){
            sum += t[i].right - '0';
        }
    }
    return 28-sum;
}

void bfs(Tree *t,int nodes){
    queue<Tree> q;
    Tree tmp;
    bool flag = false;
    int r = getRoot(t, nodes);
    q.push(t[r]);
    visited[r] = true;
    while(q.size() != 0){
        tmp = q.front();
        q.pop();
        if(tmp.right == '-' && tmp.left == '-'){
            if(flag){
                cout << " " << tmp.element;
            }else{
                cout << tmp.element;
            }
            flag = true;
        }

        if(tmp.left != '-' && !visited[tmp.left - '0']){
            q.push(t[tmp.left-'0']);
            visited[tmp.left-'0'] = true;
        }

        if(tmp.right != '-' && !visited[tmp.right - '0']){
            q.push(t[tmp.right-'0']);
            visited[tmp.right-'0'] = true;
        }
    }
}

int main(){
    int nodes;
    char right, left;
    Tree *t = (Tree *)malloc(nodes *sizeof(Tree));
    cin >> nodes;
    for(int i = 0; i < nodes;i++){
        cin >> left >> right;
        t[i].right = right;
        t[i].left = left;
        t[i].element = i;
    }
    bfs(t,nodes);
}