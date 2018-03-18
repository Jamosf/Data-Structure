/*
2.N皇后问题
N*N的国际棋牌，摆上N个皇后，互不攻击。
把每一列皇后的位置存于数组中，数组的值就是皇后在一列中的位置，只需要比较每两个皇后是否在同一行或者对角线上即可！
*/
#include<iostream>
using namespace std;

#define N 8

int a[N+1];
int v[N+1]={0};
int sum = 0;

void judge(){   
    for(int i=1;i <= N;i++){
        for(int j=1;j < i;j++){
            if(abs(j-i) == abs(a[j]-a[i])){
                return;
            }
        }
    }
    sum++;
}

void dfs(int x){
    if(x > N){
        judge();
    }
    for(int i=1;i <= N;i++){
        if(v[i] == 0){
            v[i] = 1;
            a[x] = i;
            dfs(x+1);
            v[i] = 0;
        }
    }
}

int main(){
    cout << "the proram is running..." << endl;
    dfs(1);
    cout << "res:" << sum << endl;
    return 0;
}