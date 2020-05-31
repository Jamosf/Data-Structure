"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution(object):
    def dfs(self,root,result):
		if not root:
			return
		for r in root.children:
			self.dfs(r,result)
		result.append(root.val)
    
    def postorder(self, root):
	    if not root :
	        return []
	    result = []
	    self.dfs(root,result)
	    return result