import sys
r=sys.stdin.readline
N=int(r())
g=[]
INF=sys.maxsize
for _ in range(N):
    g.append(list(map(int,r().split())))
dp=[[None]*(1<<N) for _ in range(N)]

def tsp(cur,visited):
    ret=dp[cur][visited]
    if ret!=None:
        return ret
    if visited==(1<<N)-1:
        return g[cur][0] or INF
    
    tmp=INF
    for i in range(N):
        if g[cur][i]!=0 and visited&(1<<i)==0:
            tmp=min(tmp,tsp(i,visited|(1<<i))+ g[cur][i]) 
    dp[cur][visited]=tmp
    return tmp
print(tsp(0,1<<0))


