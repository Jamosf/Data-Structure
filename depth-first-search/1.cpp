#include<iostream>
using namespace std;

int a[13]={0},v[13]={0};

void judge(){
    int i,b[6];
    b[0] = a[1]+a[3]+a[6]+a[8];
    b[1] = a[1]+a[4]+a[7]+a[11];
    b[2] = a[2]+a[3]+a[4]+a[5];
    b[3] = a[2]+a[6]+a[9]+a[12];
    b[4] = a[5]+a[7]+a[10]+a[12];
    b[5] = a[8]+a[9]+a[10]+a[11];
    for (i=1;i < 6;i++){
        if(b[i-1]!=b[i]){
            return;
        }
    }
    cout << "res:" << a[6] << endl;
}

void dfs(int x){
    int i,b[6];
    if (x==1||x==2||x==12){
        dfs(x+1);
        return;
    }

    if(x ==13){
        judge();
    }

    for(i=1;i<13;i++){
        if(v[i]==0){
            v[i] = 1;
            a[x] = i;
            dfs(x+1);
            v[i] = 0;
        }
    }
    

}

int main(){
    int i;
    a[1] = 1;
    a[2] = 8;
    a[12] = 3;
    v[1] = 1;
    v[8] = 1;
    v[3] = 1;
    dfs(1);
    return 0;
}
