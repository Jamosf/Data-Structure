#include <stdio.h>
#include <stdlib.h>

/*ͼ���ڽӾ����ʾ����c����ʵ�֣�*/
#define MaxVertexNum 2 /*��󶥵�����Ϊ100*/
#define INFINITY 65535   /*��Ϊ˫�ֽ��޷������������ֵ65535*/
typedef char VertexType;  /*����������Ϊ�ַ���*/
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
    printf("�����붥�����ͱ����������ʽΪ���������� ��������\n");
    scanf("%d,%d", &(G->n), &(G->e));
    printf("�����붥����Ϣ�������ʽΪ�������<CR>����\n");
    for(i = 0;i < G->n;i++){
        scanf("%c",&(G->vertices[i]));
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

int main(){
    MGraph *g = (MGraph *)malloc(sizeof(MGraph));
    CreateMGraph(g);
    system("pause");
    return 0;
}
