#include<iostream>
using namespace std;

#define Max 51
int A[Max] = {0};

int main(){
    int n;
    cin >> n;
    for(int i = 0; i < n;i++){
        int k;
        cin >> k;
        A[k]++;
    }
    for(int i = 0; i < Max;i++){
        if(A[i] != 0){
            cout<<i<<":"<<A[i]<<endl;
        }
    }
}