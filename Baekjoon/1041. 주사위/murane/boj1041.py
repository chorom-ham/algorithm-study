import sys
r=sys.stdin.readline
N=int(r())
dice=list(map(int,r().split()))
#3면 4개
#2면 (N-2)*4 + (N-1)*4 = (2N-3)*4
#1면 (N-2)**2개 + 4 * (N-2)(N-1)
ans=0
one=min(dice)
ans+=(one*((N-2)**2 + 4*(N-2)*(N-1)))
two=[
    dice[0]+dice[1],dice[0]+dice[3],dice[0]+dice[2],dice[0]+dice[4],
    dice[5]+dice[1],dice[5]+dice[3],dice[5]+dice[2],dice[5]+dice[4],
    dice[1]+dice[2],dice[2]+dice[4],dice[3]+dice[4],dice[1]+dice[3]
    ]
ans+=(min(two)*(2*N-3)*4)
three=[
    dice[0]+dice[3]+dice[4],dice[0]+dice[1]+dice[3],dice[0]+dice[1]+dice[2],dice[0]+dice[2]+dice[4],
    dice[3]+dice[4]+dice[5],dice[1]+dice[3]+dice[5],dice[1]+dice[2]+dice[5],dice[2]+dice[4]+dice[5]
]
ans+=(min(three)*4)
if N==1:
    print(sum(dice)-max(dice))
else:
    print(ans)