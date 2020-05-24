#include <stdio.h>
#include <stdlib.h>
#include <queue>

/*图的邻接矩阵表示法（c语言实现）*/
#define MaxVertexNum 100 /*最大顶点数设为100*/
#define INFINITY 65535   /*设为双字节无符号整数的最大值65535*/
typedef int VertexType;  /*顶点类型设为字符型*/
typedef int EdgeType;     /*边的权值设为整型*/
enum GraphType {DG,UG,DN,UN};
/*有向图，无向图，有向网图，无向网图*/

typedef struct{
    VertexType vertices[MaxVertexNum]; /*顶点表*/ 
    EdgeType Edges[MaxVertexNum][MaxVertexNum]; /*邻接矩阵，边表*/
    int n, e;
    enum GraphType GType;
}MGraph;/*MGraph是以邻接矩阵存储的图类型*/

void CreateMGraph(MGraph *G){
    int i, j, k, w;
    G->GType = UN;
    printf("请输入顶点数和边数（输入格式为：顶点数,边数）：\n");
    scanf("%d,%d", &(G->n), &(G->e));
    printf("请输入顶点信息（输入格式为：顶点号<CR>）：\n");
    for(i = 0;i < G->n;i++){
        scanf("%d",&(G->vertices[i]));
    }
    for(i = 0;i < G->n;i++){
        for(j = 0;j < G->n;j++){
            G->Edges[i][j] = INFINITY;  /*初始化邻接矩阵*/
        }
    }
    printf("请输入每条边对应的两个顶点的序号和权值，输入格式为：i,j,w:\n");
    for(k = 0;k < G->e;k++){
        scanf("%d,%d,%d",&i,&j,&w);
        G->Edges[i][j] = w;
        G->Edges[j][i] = w;            /*无向图的邻接矩阵是对称的*/
    }
}

int firstAdjV(MGraph G, int v){
    for(int j = 0;j < G.n; j++){
        if(G.Edges[v][j] < INFINITY){
            return j;
        }
    }
    return 0;
}

int nextAdjV(MGraph G, int v, int w){
    for(int j = w + 1; j < G.n; j++){
         if(G.Edges[v][j] < INFINITY){
             return j;
         }
    }
    return 0;
}

//广度优先搜索
void BFS(MGraph G){
    /*按广度优先遍历图G，使用辅助队列Q和访问标志数组visited*/
    std::queue<int> Q;
    bool visited[MaxVertexNum] = {false};
    VertexType u,v,w;
    for(u=0;u<G.n;u++){
        visited[u] = false;
    }
    for(u = 0;u < G.n;u++){
        if(!visited[u]){
            visited[u] = true;
            printf( "visit vertex: %d\n", G.vertices[u] );
            /*相当于访问顶点U*/
            Q.push(u);
            while(!Q.empty()){
                int v = Q.front();
                Q.pop();
                for(w = firstAdjV(G,v);w;w = nextAdjV(G,v,w)){
                    if(!visited[w]){
                        visited[w] = true;
                        printf( "visit vertex: %d\n", G.vertices[w] );
                        /*相当于访问顶点w*/
                        Q.push(w);
                    }
                }
            }
        }
    }
}

int main(){
    MGraph *g = (MGraph *)malloc(sizeof(MGraph));
    CreateMGraph(g);
    BFS(*g);
    system("pause");
    return 0;
}
