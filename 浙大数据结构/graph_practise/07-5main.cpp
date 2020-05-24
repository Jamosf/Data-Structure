#include <iostream>
#include <queue>
#include <stack>
using namespace std;
#define MaxCrocodiles 101
#define MaxInt 65535

bool visited[MaxCrocodiles] = {false};
int dist[MaxCrocodiles] = {MaxInt};
int path[MaxCrocodiles] = {MaxInt};
bool flag = false;

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

int firstJump(int pos, Corcodiles cor){
    int i;
    for(i = pos; i != 0; i = path[i]){}
    return cor[i].x * cor[i].x + cor[i].y + cor[i].y;
}

int bfs(Graph *G, Corcodiles cor, int step){
    queue<int> s;
    int i;
    visited[0] = true;
    s.push(0);
    dist[0] = 0;
    int min = MaxInt;
    int firstDis = MaxInt;
    int pos = 0;
    while(s.size() != 0){
        i = s.front();
        s.pop();
        if(isSafe(cor[i].x, cor[i].y, step) && dist[i] < min || (dist[i] == min && firstJump(i,cor) < firstDis)){
            min = dist[i];
            firstDis = firstJump(i,cor);
            flag = true;
            pos = i;
        }
        for(int j = 0; j < G->n;j++){
            if(!visited[j] && G->edges[i][j] == 1){
                visited[j] = true;
                dist[j] = dist[i] + 1;
                path[j] = i;
                s.push(j);
            }
        }
    }
    return pos;
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
    int pos = bfs(&G,cor, n);
    stack<Corcodile> s;
    Corcodile c;
    if(flag){
        cout << dist[pos] + 1 << endl;
        for(int i = pos; i != 0; i = path[i]){
            c.x = cor[i].x;
            c.y = cor[i].y;
            s.push(c);
        }
        while(s.size() != 0){
            cout << s.top().x << " " << s.top().y <<endl;
            s.pop();
        }
    }else{
        cout << "0";
    }
}