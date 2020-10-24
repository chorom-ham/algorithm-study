import sys
r=sys.stdin.readline
def LCS(A,B):
    dp=[[0]*(len(B)+1) for _ in range(len(A)+1)]
    s_dp=[[0]*(len(B)+1) for _ in range(len(A)+1)]
    for i in range(1,len(A)+1):
        for j in range(1,len(B)+1):
            if A[i-1]==B[j-1]:
                dp[i][j]=dp[i-1][j-1]+1
                s_dp[i][j]=3#대각선
            else:
                dp[i][j]=max(dp[i-1][j],dp[i][j-1])
                if dp[i-1][j]==dp[i][j]:
                    s_dp[i][j]=1#위쪽
                else:
                    s_dp[i][j]=2#왼쪽
    ans=""
    a,b=len(A),len(B)
    while dp[a][b]!=0:
        if s_dp[a][b]==1:
            a-=1
        elif s_dp[a][b]==2:
            b-=1
        else:
            ans=A[a-1]+ans
            a-=1
            b-=1
            
    return dp[-1][-1],ans

res=LCS(r().strip(),r().strip())
print(res[0])
print(res[1])