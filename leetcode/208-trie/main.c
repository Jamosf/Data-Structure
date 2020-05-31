#include <stdio.h>
#include <stdbool.h>
#define Max 26

typedef struct trie{
	struct trie *next[Max];
	int num; //标识子节点的个数
	bool nEndFlag; //标识是否为一个单词
} Trie;

/** Initialize your data structure here. */

Trie* trieCreate() {
	Trie *t = (Trie*)malloc(sizeof(Trie));
	for (int i = 0; i < Max; i++) {
		t->next[i] = NULL;
	}
	t->num = 0;
	t->nEndFlag = 0;
	return t;
}

/** Inserts a word into the trie. */
void trieInsert(Trie* obj, char* word) {
	int len = strlen(word);
	Trie * p = obj;
	for (int i = 0; i < len; i++) {
		char c = word[i] - 'a';
		if (p->next[c] == NULL) {
			p->next[c] = trieCreate();
			p->num++;
		}
		p = p->next[c];
	}
	p->nEndFlag = 1;
}

/** Returns if the word is in the trie. */
bool trieSearch(Trie* obj, char* word) {
	int len = strlen(word);
	Trie* p = obj;
	for (int i = 0; p && i < len; i++) {
		char c = word[i] - 'a';
		p = p->next[c];
	}
	return p != NULL && p->nEndFlag;
}

/** Returns if there is any word in the trie that starts with the given prefix. */
bool trieStartsWith(Trie* obj, char* prefix) {
	int len = strlen(prefix);
	Trie* p = obj;
	for (int i = 0; p && i < len; i++) {
		char c = prefix[i] - 'a';
		p = p->next[c];
	}
	return p != NULL;
}

void trieFree(Trie* obj) {
	Trie* p = obj;
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

/**
 * Your Trie struct will be instantiated and called as such:
 * Trie* obj = trieCreate();
 * trieInsert(obj, word);
 
 * bool param_2 = trieSearch(obj, word);
 
 * bool param_3 = trieStartsWith(obj, prefix);
 
 * trieFree(obj);
*/