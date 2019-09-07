//题目大意：在【l,r】区间内，不包含某个数字（0~9）的总个数。最多10位数。
#include <iostream>
#define MAXLEN 11
using namespace std;

int dp[MAXLEN];
int a[MAXLEN];

int dfs(int pos, bool lead, bool limit, int num){
    if(pos == -1){
        return 1;
    }
    if(!lead && !limit && dp[pos] != -1) return dp[pos];
    int n = limit ? a[pos] : 9;
    int ans = 0;
    for(int i = 0; i <= n; i++){
        if(num == 0){
            if(i == 0 && !lead) continue;
        }else{
            if(i == num) continue;
        }
        ans += dfs(pos -1, i == 0 && lead, i == a[pos] && limit, num);
    }
    if(!limit && !lead) dp[pos] = ans;
    return ans;
}

int solve(int x, int n){
    int pos = 0;
    memset(dp,-1,sizeof(int) * MAXLEN);
    while(x) {
        a[pos++] = x%10;
        x /= 10;
    }
    return dfs(pos-1, true, true, n);
}

bool hasNum(int x, int n){
    while(x){
        if(x%10 == n){
            return true;
        }
        x /= 10;
    }
    return false;
}

int solve2(int l, int r, int n){
    int ans = 0;
    for(int i = l; i <= r; i++){
        if(hasNum(i,n)){
            ans++;
        }
    }
    return r - l + 1 - ans;
}

int main(){
    int l, r, n;
    cin >> l >> r >> n;
    cout << solve(r,n) - solve(l-1,n) << endl;
    cout << solve2(l,r,n) << endl;
}