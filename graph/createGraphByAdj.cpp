#include <stdio.h>
#include <stdlib.h>

#define MaxVertexNum 100 /*��󶥵���Ϊ100*/
enum GraphType {DG,UG,DN,UN};

typedef struct node{    /*�߱�ڵ�*/
    int Adjv;           /*�ڽӵ���*/
    struct node *Next;  /*ָ����һ���ڽӵ��ָ����*/
    /*��Ҫ��ʾ���ϵ�Ȩֵ��Ϣ����Ӧ����һ��������weight*/
}EdgeNode;

typedef char VertexType;    /*�������ַ���ʾ*/

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
        scanf( " %c", &(G->adjlist[i].Vertex) ); /* ���붥����Ϣ */
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



