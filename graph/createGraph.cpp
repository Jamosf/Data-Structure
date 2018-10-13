#include <stdio.h>
#include <stdlib.h>

/*图的邻接矩阵表示法（c语言实现）*/
#define MaxVertexNum 2 /*最大顶点数设为100*/
#define INFINITY 65535   /*设为双字节无符号整数的最大值65535*/
typedef char VertexType;  /*顶点类型设为字符型*/
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
    printf("请输入顶点数和边数（输入格式为：顶点数， 边数）：\n");
    scanf("%d,%d", &(G->n), &(G->e));
    printf("请输入顶点信息（输入格式为：顶点号<CR>）：\n");
    for(i = 0;i < G->n;i++){
        scanf("%c",&(G->vertices[i]));
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

int main(){
    MGraph *g = (MGraph *)malloc(sizeof(MGraph));
    CreateMGraph(g);
    system("pause");
    return 0;
}
