import sys
from string import digits
r=sys.stdin.readline
S=r().strip()

def recur(S):
    tmp=0
    idx=0
    while True:
        if idx==len(S):
            return tmp,idx
        if S[idx] in digits:
            tmp+=1
        elif S[idx]=="(":
            par,i=recur(S[idx+1:])
            if idx==0:
                tmp+=par
            elif S[idx-1] in digits:
                tmp-=1
                tmp+=(int(S[idx-1])*par)
            else:
                tmp+=par
            idx+=i
        elif S[idx]==")":
            return tmp,idx+1
        idx+=1

print(recur(S)[0])

