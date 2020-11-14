import sys
r=sys.stdin.readline
N,X=map(int,r().split())
if not N<=X<=N*26:#출력할 수 있는지 확인하고
    print("!")
else:
    a,z=0,0
    remain=""
    while N>0:#뒤에 Z를 최대한 땡기고 나머지를 앞에 붙인후 A로 앞을 채움
        tmp=0
        if N==1:
            remain=chr(X+64)
        elif X/(N-1)<26:
            a+=1
            tmp=1
        else:
            z+=1
            tmp=26
        X-=tmp
        N-=1
    print("A"*a+remain+"Z"*z)