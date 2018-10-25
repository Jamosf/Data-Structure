#include<iostream>
using namespace std;
#define MaxNum 10001

int unionSet[MaxNum];


// int find(int x){
//     if(x > MaxNum)return -1;
//     for(;unionSet[x] > 0; x=unionSet[x]);    
//     return x; 
// }

int find(int x){
    if(unionSet[x] < 0) return x;
    else
        return unionSet[x] = find(unionSet[x]);
}

void unionfunc(int root1,int root2){
    if(root1 != root2){
        if(unionSet[root1] < unionSet[root2]){
            unionSet[root1] = unionSet[root1] + unionSet[root2];
            unionSet[root2] = root1;            
        }
        else{
            unionSet[root2] = unionSet[root1] + unionSet[root2];
            unionSet[root1] = root2;
        }
    }
}

int main(){
    int n;
    cin >> n;
    for(int i= 1; i <= n;i++){
        unionSet[i] = -1;
    }
    char state;
    do{
        cin>>state;
        int num1,num2;
        switch(state){            
            case 'C':{
                cin >> num1 >> num2;
                if(find(num1) != find(num2)){
                    cout << "no" << endl;
                }else{
                    cout << "yes" << endl;
                }
                break;
            }
            case 'I':{
                cin >> num1 >> num2;
                unionfunc(find(num1),find(num2));
                break;
            }
            case 'S':{
                int count = 0;
                for(int i =0;i <= n;i++){
                    if(unionSet[i] == -1) count++;
                }
                if(count == 1) cout << "The network is connected." << endl;
                else cout << "There are "<< count <<" components." <<endl;
                break;
            }
        }        
    }while(state != 'S');
    return 0;
}
