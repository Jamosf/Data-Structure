#include<iostream>
#include<cstdio>
#include<cstdlib>

using namespace std;

typedef int ElementType;

//�ж����������Ƿ�һ��
bool JudgeTheSame(int origin[],int changed[],int len)
{
    for(int i=0;i<len;i++)
    {
        if(origin[i]!=changed[i])
            return false;
    }
    return true;
}

//����һ�β������򣬲��������������飬���鳤�ȣ���ǰ���������������Ԫ��Nums-1��
void InsertSort(ElementType origin[],int N,int Nums)
{
    int i;
    ElementType temp=origin[Nums]; //ȡ��δ�������еĵ�һ��Ԫ��
    for(i=Nums;i>0&&origin[i-1]>temp;i--)
    {
        origin[i]=origin[i-1];   //��С��������δ�ҵ�����λ�ã�Ԫ����������ƶ�
    }
    origin[i]=temp;
}

void Insert_Sort(int origin[],int N,int changed[])
{
    for(int i=1;i<N;i++) //�ӵڶ���Ԫ�ؿ�ʼ��������
    {
        InsertSort(origin,N,i);
        if(JudgeTheSame(origin,changed,N)) //һ�ֲ���������жϽ��
        {
            cout<<"Insertion Sort"<<endl;
            InsertSort(origin,N,i+1);
            for(int j=0;j<N-1;j++)
                cout<<origin[j]<<" ";
            cout<<origin[N-1]<<endl;
            return;
        }
    }
}

/*L=��ߵ���ʼλ�ã�R=�ұ���ʼλ�ã�RightEnd=�ұ��յ�λ��*/
void Merge(ElementType A[],ElementType TempA[],int L,int R,int RightEnd)
{
    /* �������A[L]~A[R-1]��A[R]~A[RightEnd]�鲢��һ���������� */
    int LeftEnd=R-1;
    int temp=L; //�������е���ʼλ��
    int NumElements=RightEnd-L+1;

    while(L<=LeftEnd&&R<=RightEnd)
    {
        if(A[L]<=A[R])
            TempA[temp++]=A[L++]; /*ע���±��ʹ��temp*/
        else
            TempA[temp++]=A[R++];
    }
    while(L<=LeftEnd)
        TempA[temp++]=A[L++];
    while(R<=RightEnd)
        TempA[temp++]=A[R++];

    //for(int i=0;i<NumElements;i++,RightEnd--) //����L��R�������±��Ѿ��ı䣬RightEndû�б�
    //    A[RightEnd]=TempA[RightEnd]; //�������ݵ�ԭʼ������
}

/*length = ��ǰ�������еĳ��ȣ������鲢������������*/
void Merge_pass(ElementType A[],ElementType TempA[],int N,int length)
{
    int i,j;
    for(i=0;i<=N-2*length;i+=2*length)
        Merge(A,TempA,i,i+length,i+2*length-1);
    if(i+length<N)  //�鲢���2������
        Merge(A,TempA,i,i+length,N-1);
    else  //���ֻʣһ������
    {
        for(j=i;j<N;j++)
            TempA[j]=A[j];
    }
}

void Merge_Sort(ElementType A[],int N,int changed[])
{
    int length=1; //��ʼ�������г���
    ElementType *TempA;
    TempA=(ElementType*)malloc(N*sizeof(ElementType)); //��ǰ����ÿռ�

    if(TempA!=NULL)
    {
        while(length<N)
        {
            Merge_pass(A,TempA,N,length);
            if(JudgeTheSame(TempA,changed,N)) //�鲢�����Ľ��
            {
                cout<<"Merge Sort\n";
                length*=2;
                Merge_pass(TempA,A,N,length); //�ٹ鲢һ��,������˳��,��Լ�ռ䣬��������A,TempA
                for(int i=0;i<N-1;i++)
                    cout<<A[i]<<" ";
                cout<<A[N-1]<<endl;
                return;
            }

            length*=2;
            Merge_pass(TempA,A,N,length);
            if(JudgeTheSame(A,changed,N)) //�鲢�����Ľ��
            {
                cout<<"Merge Sort\n";
                length*=2;
                Merge_pass(A,TempA,N,length); //�ٹ鲢һ��,������˳��
                for(int i=0;i<N-1;i++)
                    cout << TempA[i] << " ";
                cout << TempA[N - 1] << endl;
                return;
            }

            length*=2;
        }
        free(TempA);
    }
    else
    {
        cout<<"�ռ䲻��"<<endl;
    }
}

void swap(int *a, int *b){
    int tmp;
    tmp = *a;
    *a = *b;
    *b = tmp;
}

void PercDown( ElementType A[], int p, int N )
{ /* �ı����4.24��PercDown( MaxHeap H, int p )    */
  /* ��N��Ԫ�ص���������A[p]Ϊ�����Ӷѵ���Ϊ���� */
    int Parent, Child;
    ElementType X;
 
    X = A[p]; /* ȡ��������ŵ�ֵ */
    for( Parent=p; (Parent*2+1)<N; Parent=Child ) {
        Child = Parent * 2 + 1;
        if( (Child!=N-1) && (A[Child]<A[Child+1]) )
            Child++;  /* Childָ�������ӽ��Ľϴ��� */
        if( X >= A[Child] ) break; /* �ҵ��˺���λ�� */
        else  /* ����X */
            A[Parent] = A[Child];
    }
    A[Parent] = X;
}

//������2
void heap_sort(ElementType A[], int N, ElementType changed[]){
    for(int i=N/2;i>=0;i--){
        PercDown(A,i,N);
    }
    for(int i=N-1;i>0;i--){
        swap(&A[0],&A[i]);
        PercDown(A,0,i);
        if(JudgeTheSame(A,changed,N)){
            cout<<"Heap Sort\n";
            if(i > 0){
                swap(&A[0],&A[i-1]);
                PercDown(A,0,i-1);
            }
            cout << A[0];
            for(int j = 1; j < N;j++){
                cout <<" "<< A[j];
            }
        }
    }
}

int main()
{
    int N;
    int origin[105],origin_copy[105],changed[105];
    cin>>N;
    for(int i=0;i<N;i++)
    {
        cin>>origin[i];
        origin_copy[i]=origin[i];
    }

    for(int i=0;i<N;i++)  //�м�������
        cin>>changed[i];

    Insert_Sort(origin,N,changed);
    heap_sort(origin_copy,N,changed);

    return 0;
}