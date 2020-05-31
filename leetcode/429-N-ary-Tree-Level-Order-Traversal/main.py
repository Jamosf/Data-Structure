# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children

class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children

class Solution(object):
    def traverse(self,root,level,result):
        if not root:
            return
        if len(result) == level:
            result.append([])
        result[level].append(root.val)
        for child in root.children:
            self.traverse(child,level+1,result)
    
    def levelOrder(self, root):
        if not root:
            return []
        result = []
        self.traverse(root,0,result)
        return result