import sys
from collections import deque
import heapq
r=sys.stdin.readline
A,B=map(int,r().split())
def BtoA(A,B):
    cnt=1
    while A!=B:
        cnt+=1
        if A>B:
            return -1
        if B%2==0:
            B/=2
        elif B%10==1:
            B-=1
            B/=10
        else:
            return -1
    return cnt
print(BtoA(A,B))