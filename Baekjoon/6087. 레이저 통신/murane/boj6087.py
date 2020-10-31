import sys
from collections import deque
r=sys.stdin.readline
W,H=map(int,r().split())
Map=[]
for _ in range(H):
    line=list(r().strip())
    Map.append(line)
coords=[]
for i in range(H):
    for j in range(W):
        if Map[i][j]=="C": coords.append((i,j))
di=[(1,0),(0,1),(-1,0),(0,-1)]
visited=[[10001 for _ in range(W)] for _ in range(H)]
ans=[]
def ck(x,y):
    if not 0<=x<H or not 0<=y<W:
        return False
    if Map[x][y]=="*":
        return False
    return True
def bfs():
    q=deque()
    s=coords[0]    #시작점, 끝점
    q.append((s[0],s[1]))
    visited[s[0]][s[1]]=0
    while q:
        x,y=q.popleft()
        for i in range(4):
            Nx,Ny=x+di[i][0],y+di[i][1]
            while True:
                if ck(Nx,Ny) and visited[Nx][Ny]>=visited[x][y]+1:
                    q.append((Nx,Ny))
                    visited[Nx][Ny]=visited[x][y]+1
                else: break
                Nx,Ny=Nx+di[i][0],Ny+di[i][1]
bfs()
e=coords[1]
print(visited[e[0]][e[1]]-1)