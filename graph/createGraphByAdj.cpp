#include <stdio.h>
#include <stdlib.h>

#define MaxVertexNum 100 /*��󶥵���Ϊ100*/
enum GraphType {DG,UG,DN,UN};

typedef struct node{    /*�߱�ڵ�*/
    int Adjv;           /*�ڽӵ���*/
    struct node *Next;  /*ָ����һ���ڽӵ��ָ����*/
    /*��Ҫ��ʾ���ϵ�Ȩֵ��Ϣ����Ӧ����һ��������weight*/
}EdgeNode;

typedef int VertexType;    /*�������ַ���ʾ*/

typedef struct Vnode{
    VertexType Vertex;      /*������*/
    EdgeNode *FirstEdge;    /*�߱�ͷָ��*/
}VertexNode;

typedef VertexNode AdjList[MaxVertexNum];   /*�ڽӱ�����*/

typedef struct{
    AdjList adjlist;         /*�ڽӱ�*/
    int n, e;               /*�������ͱ���*/
    enum GraphType GType; 
}ALGraph;           /*ALGraph�����ڽӱ�ʽ�����ͼ����*/

void CreateALGraph(ALGraph *G){
    int i, j, k;
    EdgeNode *edge;
    G->GType = DG;
    printf( "�����붥�����ͱ���(�����ʽΪ:������,����)�� \n" );
    scanf( "%d,%d", &(G->n), &(G->e) ); /* ���붥�����ͱ��� */
    printf( "�����붥����Ϣ(�����ʽΪ:�����<CR>)�� \n" );
    for(i = 0; i < G->n; i++){
        scanf( "%d", &(G->adjlist[i].Vertex) ); /* ���붥����Ϣ */
        G->adjlist[i].FirstEdge = NULL; /* ����ı߱�ͷָ����Ϊ�� */
    }
    printf( "������ߵ���Ϣ(�����ʽΪ: i, j <CR>)�� \n" );
    for(k = 0; k < G->e; k++){
        scanf( "\n%d,%d", &i, &j); /* �����<vi,vj>�Ķ����Ӧ���*/
        edge = (EdgeNode*)malloc(sizeof(EdgeNode));
        edge->Adjv = j;
        edge->Next = G->adjlist[i].FirstEdge;
        /* ���±߱��� edge ���뵽���� vi �ı߱�ͷ�� */
        G->adjlist[i].FirstEdge = edge;
        /* ��������ͼ����Ҫ����һ����㣬������ʾ��< vj, vi> */
    }
}

bool visited[] = {false};

//�����������
void DFS(ALGraph *G,int i){
    /*��viΪ��������ڽӱ�洢��ͼG����DFS����*/
    EdgeNode *w;
    printf( "visit vertex: %d\n", G->adjlist[i].Vertex );
    /*�൱�ڷ��ʶ���vi*/
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


