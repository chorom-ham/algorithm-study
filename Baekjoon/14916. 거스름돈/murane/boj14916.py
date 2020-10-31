import sys
r=sys.stdin.readline
n=int(r())
limit=n//5
for i in range(limit,-1,-1):
    m=n-(i*5)
    if m%2==0:
        print(i+m//2)
        exit(0)
print(-1)