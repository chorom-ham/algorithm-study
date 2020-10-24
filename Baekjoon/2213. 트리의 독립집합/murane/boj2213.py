import sys
r=sys.stdin.readline
rr=sys.stdin.readlines
n=int(r())
V=[-1]+list(map(int,r().split()))
g=[[] for _ in range(n+1)]
for line in rr():
    s,e=map(int,line.split())
    g[s].append(e)
    g[e].append(s)

tree=[[] for _ in range(n+1)]
parents=[0]*(n+1)
def dfsTree(cur,parent):
    for node in g[cur]:#현재 노드에서 갈수있는 노드들을 순회
        if node!=parent:
            tree[cur].append(node)
            parents[node]=cur
            dfsTree(node,cur)
dfsTree(1,0)

#dp[i][0] => i를 루트로하는 서브트리에서 i를 포함하지 않았을때 크기
#dp[i][1] => i를 루트로하는 서브트리에서 i를 포함했을때 크기
dp=[[-1]*2 for _ in range(n+1)]

def inde(i,have):
    ret=dp[i][have]
    if ret!=-1:
        return ret
    if not tree[i]:#리프노드의 경우
        if have==1:
            dp[i][have]=V[i]
            return V[i]
        else:
            dp[i][have]=0 
            return 0
    tmp=0
    for node in tree[i]:#자식노드를 탐색
        if have==1:
            tmp+=inde(node,0)
        else:
            tmp+=max(inde(node,0),inde(node,1))
    if have==1: tmp+=V[i]
    dp[i][have]=tmp
    return tmp
print(max(inde(1,1),inde(1,0)))
ans=[]
def findV(v):
    if dp[v][1]>=dp[v][0] and parents[v] not in ans:
        ans.append(v)
    for node in tree[v]:
        findV(node)
findV(1)
ans.sort()
print(*ans)