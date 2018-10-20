//浙江大学数据结构课程作业：https://pintia.cn/problem-sets/1010070491934568448/problems/1048028714800685056
#include <iostream>
#include <queue>
#include <stack>
#define MaxVertexNum 15

using namespace std;

typedef int VertexType;
typedef int EdgeType;

struct MGraph{
    int n,e;
    VertexType vertexs[MaxVertexNum];
    EdgeType edges[MaxVertexNum][MaxVertexNum];
};

bool visited[MaxVertexNum] = {false};

void bfs(MGraph graph){
    for(int i = 0; i < graph.n; i++){
        if(!visited[i]){
            queue<int> q;
            visited[i] = true;
            q.push(i);
            cout<<"{ ";
            while (q.size()!= 0){
                int num = q.front();
                q.pop();
                cout<<num<<" ";
                for(int j = 0;j < graph.n;j++){
                    if(graph.edges[num][j] == 1 && !visited[j]){
                        visited[j] = true;
                        q.push(j);
                    }
                }
            }
            cout << "}" << endl;
        }
    }
}

void dfs(MGraph graph){
    for(int i = 0; i < graph.n; i++){
        if(!visited[i]){
            stack<int> s;
            visited[i] = true;
            s.push(i);
            cout<<"{ ";
            while (s.size()!= 0){
                int num = s.top();
                s.pop();
                cout<<num<<" ";
                for(int j = graph.n - 1;j >= 0;j--){
                    if(graph.edges[num][j] == 1 && !visited[j]){
                        visited[j] = true;
                        s.push(j);
                    }
                }
            }
            cout << "}" << endl;
        }
    }
}

void dfs_recursion(MGraph graph, int i){
    visited[i] = true;
    cout << i << " ";
    for(int j = 0; j < graph.n; j++){
        if(graph.edges[i][j] == 1 && !visited[j]){
            dfs_recursion(graph,j);
        }
    }

}

int main() {
    int n,e;
    MGraph g;

    cin >> n >> e;

    if(n > 10 || e < 0){
        return 0;
    }
    for(int i = 0; i < n;i++){
        for(int j = 0 ; j < n; j++){
            g.edges[i][j] = 0;
        }
    }
    g.n =n;
    g.e = e;
    for(int i = 0; i < n; i++){
        g.vertexs[i] = i;
    }
    for(int k = 0; k < e; k++){
        int i,j;
        cin >> i >> j;
        g.edges[i][j] = 1;
        g.edges[j][i] = 1;
    }

    for(int i = 0; i < g.n;i++){
        if(!visited[i]){
            cout << "{ ";
            dfs_recursion(g, i);
            cout << "}" << endl;
        }
    }
    for(int i=0;i < MaxVertexNum;i++){
        visited[i] = false;
    }
    bfs(g);
//    for(int i=0;i < MaxVertexNum;i++){
//        visited[i] = false;
//    }
//  dfs(g);
    system("pause");
    return 0;
}