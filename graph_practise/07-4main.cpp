//
// 浙江大学数据结构课程作业：https://pintia.cn/problem-sets/1010070491934568448/problems/1050401612970643456
//
#include<iostream>
#include<queue>
using namespace std;
#define MaxNum 101
int edge[MaxNum][MaxNum], N, E;
bool visited[MaxNum] = {false};
int dist[MaxNum][MaxNum];

void floyd(){
    for(int i = 1; i  <= N; i++){
        for(int j = 1; j <= N; j++){
            dist[i][j] = edge[i][j];
        }
    }
    for(int k = 1; k <= N; k++){
        for(int i = 1; i <= N; i++){
            for(int j = 1; j <= N; j++){
                if(dist[i][j] > dist[i][k] + dist[k][j]){
                    dist[i][j] = dist[i][k] + dist[k][j];
                }
            }
        }
    }
}

bool isConnectedGraph(){
    queue<int> q;
    visited[1] = true;
    q.push(1);
    while (q.size()!= 0){
        int num = q.front();
        q.pop();
        for(int j = 1;j <= N;j++){
            if(edge[num][j] != 0 && edge[num][j] != 65535 && !visited[j]){
                visited[j] = true;
                q.push(j);
            }
        }
    }
    for(int i = 1; i <= N; i++){
        if(!visited[i]){
            return false;
        }
    }
    return true;
}

int main(){
    cin >> N >> E;
    int node1,node2,weight;
    for(int i = 1; i <= N;i++){
        for(int j = 1; j <= N;j++){
            edge[i][j] = 65535;
            if(i == j){
                edge[i][j] = 0;
            }
        }
    }
    for(int i = 0; i < E;i++){
        cin >> node1 >> node2 >> weight;
        edge[node1][node2] = weight;
        edge[node2][node1] = weight;
    }
    if(!isConnectedGraph()){
        cout << 0;
        return 0;
    }
    floyd();
    int result[MaxNum] = {65535};
    for(int i = 1; i <= N; i++){
        int max = 0;
        for(int j = 1; j <= N;j++){
            if(dist[i][j] > max){
                max = dist[i][j];
            }
        }
        result[i] = max;
    }
    int min = 65535;
    int location = 1;
    for(int i = 1; i <= N; i++){
        if(result[i] < min){
            min = result[i];
            location = i;
        }
    }
    cout << location << " " << min;
    system("pause");
    return 0;
}

