#include<iostream>

#define MaxNum 1000
#define MinNum -10001
using namespace std;

struct Heap{
    int data[MaxNum];
    int size;
    int capacity;
};

Heap h;
void createHeap(){
    h.capacity = MaxNum;
    h.data[0] = MinNum;
}

void insert(int x){
    int i;
    h.size++;
    if(h.size > h.capacity){
        return;
    }
    for(i = h.size;h.data[i/2] > x;i/=2){
        h.data[i] = h.data[i/2];
    }
    h.data[i] = x;
}


int main(){
    int n,m;
    cin >> n >> m;
    createHeap();
    for(int i = 0;i < n;i++){
        int num;
        cin >> num;
        insert(num);
    }
    for(int i = 0; i < m;i++){
        int num;
        cin >> num;
        cout << h.data[num] ;
        while(num > 1){
            num /=2;
            cout << " "<< h.data[num];
        }
        cout << endl;
    }
    //system("pause");
    return 0;
}