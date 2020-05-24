#include <stdio.h>
#include <stdlib.h>
#include <queue>

/*ͼ���ڽӾ����ʾ����c����ʵ�֣�*/
#define MaxVertexNum 100 /*��󶥵�����Ϊ100*/
#define INFINITY 65535   /*��Ϊ˫�ֽ��޷������������ֵ65535*/
typedef int VertexType;  /*����������Ϊ�ַ���*/
typedef int EdgeType;     /*�ߵ�Ȩֵ��Ϊ����*/
enum GraphType {DG,UG,DN,UN};
/*����ͼ������ͼ��������ͼ��������ͼ*/

typedef struct{
    VertexType vertices[MaxVertexNum]; /*�����*/ 
    EdgeType Edges[MaxVertexNum][MaxVertexNum]; /*�ڽӾ��󣬱߱�*/
    int n, e;
    enum GraphType GType;
}MGraph;/*MGraph�����ڽӾ���洢��ͼ����*/

void CreateMGraph(MGraph *G){
    int i, j, k, w;
    G->GType = UN;
    printf("�����붥�����ͱ����������ʽΪ��������,��������\n");
    scanf("%d,%d", &(G->n), &(G->e));
    printf("�����붥����Ϣ�������ʽΪ�������<CR>����\n");
    for(i = 0;i < G->n;i++){
        scanf("%d",&(G->vertices[i]));
    }
    for(i = 0;i < G->n;i++){
        for(j = 0;j < G->n;j++){
            G->Edges[i][j] = INFINITY;  /*��ʼ���ڽӾ���*/
        }
    }
    printf("������ÿ���߶�Ӧ�������������ź�Ȩֵ�������ʽΪ��i,j,w:\n");
    for(k = 0;k < G->e;k++){
        scanf("%d,%d,%d",&i,&j,&w);
        G->Edges[i][j] = w;
        G->Edges[j][i] = w;            /*����ͼ���ڽӾ����ǶԳƵ�*/
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

//�����������
void BFS(MGraph G){
    /*��������ȱ���ͼG��ʹ�ø�������Q�ͷ��ʱ�־����visited*/
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
            /*�൱�ڷ��ʶ���U*/
            Q.push(u);
            while(!Q.empty()){
                int v = Q.front();
                Q.pop();
                for(w = firstAdjV(G,v);w;w = nextAdjV(G,v,w)){
                    if(!visited[w]){
                        visited[w] = true;
                        printf( "visit vertex: %d\n", G.vertices[w] );
                        /*�൱�ڷ��ʶ���w*/
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
