/*
4.电梯问题
从A向上或者向下搜索，终止条件是判断，是否到达B.
*/
#include<iostream>
using namespace std;

#define MAX 500

int a[MAX];
int result = 0,n;
int A,B;

void judge(int total,int x){   
    if(x+a[x] == B || x-a[x] == B){
        if(total < result || result == 0){
            result = total;
        }
    }    
}

void dfs(int total,int x){
    judge(total,x);
    if(x + a[x] <= B){
        dfs(total+1,x + a[x]);
    }
    if(x-a[x] >= 1){
        dfs(total+1,x-a[x]);
    }
}

int main(){
    cout << "the proram is running..." << endl;
    cin >> n >> A >> B;
    for(int i=1; i<=n;i++){
        cin >> a[i];
    }
    dfs(1,A);
    if(result) cout << "res:" << result << endl;
    else cout << "res:" << -1 << endl;
    return 0;
}