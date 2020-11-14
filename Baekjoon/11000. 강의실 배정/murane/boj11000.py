import sys
from collections import defaultdict
r=sys.stdin.readline
start=defaultdict(int)
end=defaultdict(int)
N=int(r())
for _ in range(N):
    s,e=map(int,r().split())
    start[s]+=1
    end[e]+=1
ans=0
cur=0
time=set()
time.update(list(start.keys()))
time.update(list(end.keys()))
for t in sorted(list(time)):
    cur+=start[t]
    cur-=end[t]
    ans=max(ans,cur)
print(ans)
