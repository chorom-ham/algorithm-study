import sys
from collections import defaultdict
r=sys.stdin.readline
def BellmanFord():
    possible=False
    INF=sys.maxsize
    #모든 정점에 대해 최단거리를 구한다
    dist=[INF]*(N+1)
    visit=[False]*(N+1)
    for v in range(1,N+1):
        #사이클을 확인한 그룹은 방문하지 않는다
        if visit[v]:continue
        dist[v]=0
        visit[v]=True
        flg=False
        for i in range(N):
            for S in range(1,N+1):
                for E,D in doro[S]:
                    if dist[S]!=INF and dist[E] > dist[S]+D:
                        visit[E]=True
                        dist[E] = dist[S]+D
                        if i==N-1:flg=True
        if flg:
            possible=True
            break
    return possible
for _ in range(int(r())):
    N,M,W=map(int,r().split())
    doro=defaultdict(set)
    for _ in range(M):
        S,E,T=map(int,r().split())
        doro[S].add((E,T))
        doro[E].add((S,T))
    #웜홀은 음수가중치의 단방향 그래프로 생성한다.
    for _ in range(W):
        S,E,T=map(int,r().split())
        doro[S].add((E,-T))
    if BellmanFord():print("YES")
    else:print("NO")
