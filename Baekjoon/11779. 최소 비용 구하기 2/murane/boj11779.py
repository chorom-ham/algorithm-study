import sys
import heapq
r=sys.stdin.readline
n=int(r())
m=int(r())
city=[[] for _ in range(n+1)]
INF=sys.maxsize
for _ in range(m):
    s,e,cost=map(int,r().split())
    city[s].append((e,cost))
start,end=map(int,r().split())
heap=[]
heapq.heappush(heap,(0,start))
dist=[(INF,"")]*(n+1)
dist[start]=(0,str(start))
while heap:
    d,v=heapq.heappop(heap)#start에서 v정점까지의 거리 d
    for dest,cost in city[v]:#v정점에서 갈수 있는 노드를 순회
        if dist[dest][0] > d+cost:#기존의 dest까지의 거리와 비교
            dist[dest]=(d+cost,dist[v][1]+","+str(dest))
            heapq.heappush(heap,(d+cost,dest))
            
print(dist[end][0])
route=dist[end][1].split(",")
print(len(route))
print(*route)