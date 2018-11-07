#include<iostream>
#include<string>
using namespace std;

#define Max 5
int A[Max] = {0};

struct Person{
    string namesb;
    int score[Max];
    int correctCount;
    int total;
};

int str2int(string s){
    string s1;
    for(int i = 0; i < s.length();i++){
        if(s[i] != '0'){
            s1 = s.substr(i,s.length()-i);
            break;
        }
    }
    return atoi(s1.c_str());
}

int main(){
    int p, n, s;
    cin >> p >> n >> s;
    for(int i = 0; i < n;i++){
        cin >> A[i];
    }
    /*构建一个桶，用于存储person类*/
    Person *pl = new Person[p];
    for(int i = 0;i < p;i++){
        for(int j = 0; j < n;j++){
            pl[i].score[j] = -2;
        }
        pl[i].total = 0;
        pl[i].correctCount = 0;
    }
    string name;
    int number;
    int score;
    for(int i = 0; i < s;i++){
        cin >> name >> number >> score;
        pl[str2int(name)-1].namesb = name;
        if(pl[str2int(name)-1].score[number-1] < score){
            pl[str2int(name)-1].score[number-1] = score;
        }
        if(score == -1) pl[str2int(name)-1].score[number-1] = 0;
        if(score == A[number-1]){
            pl[str2int(name)-1].correctCount++;
        }
    }
    bool flag;
    for(int i = 0; i < p;i++){
        flag = false;
        for(int j = 0;j < n;j++){
            if(pl[i].score[j] > 0 || pl[i].score[j] == -1){
                pl[i].total += pl[i].score[j];
                flag = true;
            }
        }
        if(!flag)pl[i].total = -1;
    }
    Person tmp;
    for(int i = 0; i < p-1;i++){
        for(int j = 0;j < p-1-i;j++){
            if(pl[j].total < pl[j+1].total){
                tmp = pl[j];
                pl[j] = pl[j+1];
                pl[j+1] = tmp;
            }else if(pl[j].total == pl[j+1].total){
                if(pl[j].correctCount < pl[j+1].correctCount){
                    tmp = pl[j];
                    pl[j] = pl[j+1];
                    pl[j+1] = tmp;
                }
            }
        }
    }
    int rank = 0;
    for(int i = 0; i < p;i++){
        if(pl[i].total >= 0){
            if(i == 0 || (i >= 1 && pl[i].total != pl[i-1].total)){
                rank += 1;
            }
            cout << rank << " " << pl[i].namesb << " " << pl[i].total;
            for(int j = 0; j < n;j++){
                if(pl[i].score[j] >= 0 || pl[i].score[j] == -1){
                    cout << " " << pl[i].score[j];
                }else{
                    cout << " -";
                }
            }
            cout << endl;
        }
    }
    delete[]pl;
    system("pause");
}