/*
对于入栈出栈的队列我们可以知道两个关系： 
(1)对于刚出栈的数x，1~x内未出栈的个数应该小于栈的大小M。 
(2)对应刚出栈的数x，如果比前一个出栈的数y小，那么x~y之间的所有数应该全部出栈。
*/
#include<iostream>
#include<stdlib.h>

using namespace std;
#define Max 1000
bool used[Max];
int a[Max];

bool judge2(int x,int y){
    for(int i = x+1;i < y;i++){
        if(!used[i]) return false;
    }
    return true;
}

bool judge(int n,int m){
    int cur, count;
    for(int i = 0; i < n;i++){
        cur = a[i];
        count = 0;

        if(i >= 1 && a[i] < a[i-1] && !judge2(a[i],a[i-1])){
            return false;
        }
        for(int j=1; j <= a[i];j++){
            if(!used[j])count++;
            if(count > m)return false;
        }
        used[a[i]] = true;
    }
    return true;
}

int main(){
    int m,n,k;
    cin >> m >> n >> k;
    while(k--){
        for(int p = 0; p < Max;p++){
            used[p] = false;
        }
        for(int j = 0; j < n;j++){
            cin >> a[j];
        }
        if(!judge(n,m))cout << "NO" << endl;
        else{
            cout << "YES" << endl;
        }
    }
}