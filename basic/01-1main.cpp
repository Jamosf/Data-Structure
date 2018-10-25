#include<iostream>
#define MaxNum 100000
using namespace std;

int maxSubSquence(int a[],int n,int &second){
    int thisSum = 0;
    int max = 0;
    for(int i = 0; i < n;i++){
        thisSum += a[i];
        if(thisSum > max) {
            max = thisSum;
            second = i;
        }
        if(thisSum < 0) {
            thisSum = 0;
        }
    }
    return max;
}

bool isAllNegtive(int a[],int n){
    for(int i = 0; i < n;i++){
        if(a[i] >=0){
            return false;
        }
    }
    return true;
}

int main(){
    int n;
    int a[MaxNum];
    cin >> n;
    for(int i = 0; i < n; i++){
        cin >> a[i];
    }
    int first = 0;
    int second = 0;
    int sum = maxSubSquence(a,n,second);
    int temp = sum;
    for(int i = second;i >= 0;i--){
        temp -= a[i];
        if(temp == 0){
            first = i;
        }
    }
    if(sum == 0 && isAllNegtive(a,n)){
        first = 0;
        second = n-1;
    }else if(sum == 0 && !isAllNegtive(a,n)){
        cout << sum <<" "<< 0 <<" "<< 0;
        return 0;
    }
    cout << sum <<" "<< a[first] <<" "<< a[second];
    //system("pause");
    return 0;
}