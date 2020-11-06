import sys
from collections import defaultdict,deque
r=sys.stdin.readline
N,M=map(int,r().split())
students=defaultdict(set)
for _ in range(M):
    a,b=map(int,r().split())
    students[a].add(b)
    #defaultdict(set)을 이용하여 인접그래프 구성
def topologySort(graph):
    ans=[]
    q=deque()
    ck=[0]*(N+1)
    for S in graph.keys():
        for E in graph[S]:
            ck[E]+=1
    #S->E인 간선을 모두 순회하여 E에 진입차수를 ++해줌
    for i in range(1,N+1):
        if ck[i]==0:
            q.append(i)
    #ck리스트에 진입차수가 0인 시작점을 넣음
    while q:
        cur=q.popleft()
        ans.append(cur)
        tmp=graph[cur]
        graph.pop(cur)
        ck[cur]=-1
        #pop되면 답이 될수있음
        #tmp는 pop되고 다음 후보 노드들을 저장
        #그래프에서 cur을 제거하고 진입차수는 -1로 만들어줌
        for node in tmp:
            ck[node]-=1
            if ck[node]==0:
                q.append(node)
        #후보노드들의 진입차수를 -- 해주고
        #진입차수가 0이되면 시작할 수 있으므로 q에 넣음
    return ans
print(*topologySort(students))
