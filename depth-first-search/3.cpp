/*
2.带分数问题
直接用深度优先搜索去尝试，注意约束条件就可以。x+y/z=p,即x*z+y=p*z,(p-x)*z=y
*/
#include<iostream>
using namespace std;

#define N 10

int a[N];
int v[N]={0};
int result = 0,p;

int sum(int x,int y)  
{  
    int s = 0;  
      
    if(y==0)  
        return 0;  
          
    for(int i=x;i<=y;i++)  
    {  
        s = s*10+a[i];  
    }  
    return s;  
}  

void judge(){   
    int x=0,y=0,z=0;
    for(int i=0;i < N-2;i++){
        for(int j=i+1;j < N-1;j++){
            x = sum(0,i);
            y = sum(i+1,j);
            z = sum(j+1,N-1);
            if((p-x)*z == y){
                result++;
            }
        }
    }
}

void dfs(int x){
    if(x > N-1){
        judge();
    }
    for(int i=1;i < N;i++){
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
    cin >> p;
    dfs(1);
    cout << "res:" << result << endl;
    return 0;
}