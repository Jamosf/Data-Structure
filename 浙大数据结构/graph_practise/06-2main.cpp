#include <iostream>
#include <stack>
using namespace std;
#define MaxCrocodiles 100

bool visited[MaxCrocodiles] = {false};

struct Graph{
    int n, e;
    int vertex[MaxCrocodiles];
    int edges[MaxCrocodiles][MaxCrocodiles];
};

typedef struct Corcodile{
    float x;
    float y;
}Corcodiles[MaxCrocodiles];

bool isSafe(float x, float y, int step){
    if(x + step >= 50 || x - step <= -50 || y + step >= 50 || y - step <= -50){
        return true;
    }
    return false;
}

void dfs(Graph *G, Corcodiles cor, int step){
    stack<int> s;
    int i;
    visited[0] = true;
    s.push(0);
    while(s.size() != 0){
        i = s.top();
        s.pop();
        if(isSafe(cor[i].x, cor[i].y, step)){
            cout << "Yes";
            return;
        }
        for(int j = 0; j < G->n;j++){
            if(!visited[j] && G->edges[i][j] == 1){
                visited[j] = true;
                s.push(j);
            }
        }
    }
    cout << "No";
}

int main(){
    int m;
    float n;
    cin >> m >> n;
    Corcodiles cor;
    Graph G;
    G.n = m + 1;
    for(int i = 0;i <= m;i++){
        G.vertex[i] = i;
    }
    for(int i = 0;i <= m;i++){
        for(int j = 0;j <=m;j++){
            G.edges[i][j] = 0;
        }
    }
    for(int i = 1; i <= m;i++){
        cin >> cor[i].x >> cor[i].y;
        if((cor[i].x)*(cor[i].x) + (cor[i].y)*(cor[i].y) <= (n + 7.5)*(n+7.5)){
            G.edges[0][i] = 1;
            G.edges[i][0] = 1;
        }
    }
    for(int i = 1; i <= m;i++){
        for(int j = i+1; j <= m;j++){
            if((cor[i].x - cor[j].x)*(cor[i].x - cor[j].x) + (cor[i].y - cor[j].y)*(cor[i].y - cor[j].y) <= n*n){
                G.edges[i][j] = 1;
                G.edges[j][i] = 1;
            }
        }
    }
    dfs(&G,cor, n);
}