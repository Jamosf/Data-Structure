#include<iostream>
#include<string>
using namespace std;

#define Max 5
int A[Max] = {0};

struct Person{
    int score[Max];
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
    Person *pl = (Person*)malloc(sizeof(Person)*p);
    string name;
    int number;
    int score;
    for(int i = 0; i < s;i++){
        cin >> name >> number >> score;
        if(pl[str2int(name)].score[number] < score){
            pl[str2int(name)].score[number] = score;
        }
    }
    for(int i = 0; i < p;i++){
        for(int j = 0;j < n;j++){
            pl[i].total += pl[i].score[j];
        }
    }
    
}