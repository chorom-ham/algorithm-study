import sys
from collections import deque
r=sys.stdin.readline
N,M=map(int,r().split())
board=[list(r().strip()) for _ in range(N)]
R,B=(),()
for i in range(N):
    for j in range(M):
        if board[i][j]=="R":
            R=(i,j)
        if board[i][j]=="B":
            B=(i,j)
dx=[1,-1,0,0]
dy=[0,0,1,-1]
cnt=0
q=deque([R,B,0])
def move():
    
while q:

if cnt==0:
    print("-1")
else:
    print(cnt)
def check(pos,d):
    for d in range(4):
        while True:
            x,y=pos[0]+dx[d],pos[1]+dy[d]
            if 0<=x<N and 0<=y<M:
                if board[x][y]=="O":
                    flg=False
                    while True:
                        tmp=pos[0]-dx[d],pos[1]-dy[d]
                        if 0<=tmp[0]<N and 0<=tmp[1]<M:
                            if board[tmp[0]][tmp[1]]=="#":
                                break
                            if board[tmp[0]][tmp[1]]=="B":
                                flg=True
                        else:
                            break
                    if not flg:
                        return True
                    else:
                        return False 
                if board[x][y]=="B" or board[x][y]=="#":
                    break
                pos=(pos[0]+dx[d],pos[1]+dy[d])
    return False