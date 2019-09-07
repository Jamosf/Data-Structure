#include <stdio.h>
#include <stdbool.h>
#define Max 26

typedef struct wordDic{
    struct wordDic *next[Max];
    int num;
    bool isWord;
} WordDictionary;

/** Initialize your data structure here. */

WordDictionary* wordDictionaryCreate() {
    WordDictionary *w = (WordDictionary *)malloc(sizeof(WordDictionary));
    for(int i = 0; i < Max; i++){
        w->next[i] = NULL;
    }
    w->isWord = false;
    return w;
}

/** Adds a word into the data structure. */
void wordDictionaryAddWord(WordDictionary* obj, char * word) {
    WordDictionary *p = obj;
    int len = strlen(word);
    for(int i = 0; i < len; i++){
        char c = word[i] - 'a';
        if(p->next[c] == NULL){
            p->next[c] = wordDictionaryCreate();
            p->num++;
        }
        p = p->next[c];
    }
    p->isWord = true;
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
bool wordDictionarySearch(WordDictionary* obj, char * word) {
    WordDictionary *p = obj;
    int len = strlen(word);
    if(p == NULL){
        return false;
    }
    if (len == 0 && p != NULL && p->isWord){
        return true;
    }
    char *newWord = "";
    for(int i = 1; i < len; i++){
        newWord[i-1] = word[i];
    }
    if(word[0] == '.') {
        for(int j = 0; j < Max; j++){
            if(p->next[j] != NULL){
                wordDictionarySearch(p->next[j], newWord);
            }            
        }
        
    }else{
        wordDictionarySearch(p->next[word[0]-'a'], newWord);
    }
    return false;
}

void wordDictionaryFree(WordDictionary* obj) {
    WordDictionary* p = obj;
	while (obj != NULL) {
		for (int i = 0; i < Max; i++) {
			if (obj != NULL) {
				p = obj;
				obj = obj->next[i];
				free(p);
			}	
		}
	}
}

int main(){

    WordDictionary* obj = wordDictionaryCreate();
    wordDictionaryAddWord(obj, "char");
 
    bool param_2 = wordDictionarySearch(obj, "char");

    printf("%B\n",param_2);
 
    wordDictionaryFree(obj);
}

/**
 * Your WordDictionary struct will be instantiated and called as such:
 * WordDictionary* obj = wordDictionaryCreate();
 * wordDictionaryAddWord(obj, word);
 
 * bool param_2 = wordDictionarySearch(obj, word);
 
 * wordDictionaryFree(obj);
*/