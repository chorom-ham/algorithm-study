import sys
from collections import deque
r=sys.stdin.readline
M,N=map(int,r().split())
tomato = [list(map(int,r().split())) for _ in range(N)]
d=[(1,0),(-1,0),(0,1),(0,-1)]
def ck():
    for i in range(N):
        for j in range(M):
            if tomato[i][j]==0:
                return False
    return True

def bfs(coords:list):#익은 토마토에서 출발
    q=deque(coords)
    minimum_day=0
    while q:
        x,y=q.popleft()
        for i in range(4):
            Nx,Ny=x+d[i][0],y+d[i][1]
            if not 0<=Nx<N or not 0<=Ny<M: continue
            if tomato[Nx][Ny]!=0: continue
            q.append((Nx,Ny))
            tomato[Nx][Ny]=tomato[x][y]+1
    for line in tomato:
        minimum_day=max(minimum_day,max(line))
    return minimum_day
    
coords=[]
for i in range(N):
    for j in range(M):
        if tomato[i][j]==1:
            coords.append((i,j))
if ck():
    print(0)
    exit(0)
ans=bfs(coords)
if ck():
    print(ans-1)
else:
    print(-1)