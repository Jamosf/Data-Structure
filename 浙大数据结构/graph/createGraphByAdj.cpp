#include <stdio.h>
#include <stdlib.h>

#define MaxVertexNum 100 /*最大顶点数为100*/
enum GraphType {DG,UG,DN,UN};

typedef struct node{    /*边表节点*/
    int Adjv;           /*邻接点域*/
    struct node *Next;  /*指向下一个邻接点的指针域*/
    /*若要表示边上的权值信息，则应增加一个数据域weight*/
}EdgeNode;

typedef int VertexType;    /*顶点用字符表示*/

typedef struct Vnode{
    VertexType Vertex;      /*顶点域*/
    EdgeNode *FirstEdge;    /*边表头指针*/
}VertexNode;

typedef VertexNode AdjList[MaxVertexNum];   /*邻接表类型*/

typedef struct{
    AdjList adjlist;         /*邻接表*/
    int n, e;               /*顶点数和边数*/
    enum GraphType GType; 
}ALGraph;           /*ALGraph是以邻接表方式储存的图类型*/

void CreateALGraph(ALGraph *G){
    int i, j, k;
    EdgeNode *edge;
    G->GType = DG;
    printf( "请输入顶点数和边数(输入格式为:顶点数,边数)： \n" );
    scanf( "%d,%d", &(G->n), &(G->e) ); /* 读入顶点数和边数 */
    printf( "请输入顶点信息(输入格式为:顶点号<CR>)： \n" );
    for(i = 0; i < G->n; i++){
        scanf( "%d", &(G->adjlist[i].Vertex) ); /* 读入顶点信息 */
        G->adjlist[i].FirstEdge = NULL; /* 顶点的边表头指针设为空 */
    }
    printf( "请输入边的信息(输入格式为: i, j <CR>)： \n" );
    for(k = 0; k < G->e; k++){
        scanf( "\n%d,%d", &i, &j); /* 读入边<vi,vj>的顶点对应序号*/
        edge = (EdgeNode*)malloc(sizeof(EdgeNode));
        edge->Adjv = j;
        edge->Next = G->adjlist[i].FirstEdge;
        /* 将新边表结点 edge 插入到顶点 vi 的边表头部 */
        G->adjlist[i].FirstEdge = edge;
        /* 若是无向图，还要生成一个结点，用来表示边< vj, vi> */
    }
}

bool visited[] = {false};

//深度优先搜索
void DFS(ALGraph *G,int i){
    /*以vi为出发点对邻接表存储的图G进行DFS搜索*/
    EdgeNode *w;
    printf( "visit vertex: %d\n", G->adjlist[i].Vertex );
    /*相当于访问顶点vi*/
    visited[i] = true;
    for(w = G->adjlist[i].FirstEdge;w;w=w->Next){
        if(!visited[w->Adjv]){
            DFS(G,w->Adjv);
        }
    }
}

int main(int argc, char const *argv[])
{
    ALGraph *G = (ALGraph*)malloc(sizeof(ALGraph));
    CreateALGraph(G);
    DFS(G,0);
    system("pause");
    return 0;
}


