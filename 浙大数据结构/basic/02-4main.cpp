/*
������ջ��ջ�Ķ������ǿ���֪��������ϵ�� 
(1)���ڸճ�ջ����x��1~x��δ��ջ�ĸ���Ӧ��С��ջ�Ĵ�СM�� 
(2)��Ӧ�ճ�ջ����x�������ǰһ����ջ����yС����ôx~y֮���������Ӧ��ȫ����ջ��
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