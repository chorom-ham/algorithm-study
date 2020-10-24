import sys
r=sys.stdin.readline
N=int(r())
matrix=[]
for _ in range(N):
    matrix.append(list(map(int,r().split())))
res=0
dp=[[0]*N for _ in range(N)]
for i in range(1,N):
    for j in range(N-i):
        tmp=[]
        for k in range(j,i+j):
            tmp.append(dp[j][k]+dp[k+1][i+j]+matrix[j][0]*matrix[k][1]*matrix[i+j][1])
        dp[j][i+j]=min(tmp)

print(dp[0][N-1])