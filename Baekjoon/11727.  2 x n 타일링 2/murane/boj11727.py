import sys
r=sys.stdin.readline
N=int(r())
res=0
dp=[0,1,3]+[-1]*(1000-1)
for i in range(3,N+1):
    dp[i]=dp[i-2]*2+dp[i-1]
print(dp[N]%10007)