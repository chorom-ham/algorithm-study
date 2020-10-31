import sys
from collections import deque
r=sys.stdin.readline
N,M=map(int,r().split())
maps=[]
for _ in range(N):
    maps.append(list(map(int,r().split())))
d=[(1,0),(-1,0),(0,1),(0,-1)]

def bfs(x,y,cur):
    q=deque([(x,y)])
    maps[x][y]=cur
    while q:
        x,y=q.popleft()
        for i in range(4):
            Nx,Ny=x+d[i][0],y+d[i][1]
            if not 0<=Nx<N or not 0<=Ny<M or maps[Nx][Ny]!=1:
                continue
            maps[Nx][Ny]=cur
            q.append((Nx,Ny))
cur=2
for i in range(N):
    for j in range(M):
        if maps[i][j]!=1: continue
        bfs(i,j,cur)
        cur+=1

distDict=dict()
for i in range(N):
    for j in range(M):
        curCountry=maps[i][j]
        if curCountry!=0:#이게 나라냐..
            for k in range(4):
                connect=False
                dest=0
                dist=0
                Ni,Nj=i+d[k][0],j+d[k][1]
                while True:
                    if not 0<=Ni<N or not 0<=Nj<M:
                        break
                    elif maps[Ni][Nj]==0:
                        dist+=1
                    elif maps[Ni][Nj]==curCountry:
                        pass
                    else:
                        connect=True
                        dest=maps[Ni][Nj]
                        break
                    Ni+=d[k][0]
                    Nj+=d[k][1]
                if connect:
                    pair=(min(curCountry,dest),max(curCountry,dest))
                    if pair in distDict:
                        distDict[pair]=min(dist,distDict[pair])
                    else:
                        distDict[pair]=dist

parents=[i for i in range(cur)]
def union(x,y):
    x,y=find(x),find(y)
    if x!=y:
        parents[x]=y
def find(x):
    if x==parents[x]:
        return x
    else:
        y=find(parents[x])
        parents[x]=y
        return y
print(1)