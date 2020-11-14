import sys
r=sys.stdin.readline
def verify(num):
    while num>=100:
        num/=100
    if num==25:
        return True
    else:
        return False
for _ in range(int(r())):
    money=int(r())
    ans=0
    coins=[10**x for x in range(16)]+[25*(100**x) for x in range(7)]
    coins=list(filter(lambda x: x<=money,coins))
    coins.sort()
    #25*100**x의 경우 하위의 2개 동전과 비교하여 최대치를 구한다
    while coins:
        if money==0: break
        l=len(coins)-1
        cur=coins.pop()
        if verify(cur):
            q=money//cur
            cnt=9999
            for quarter_cnt in range(q+1):
                tmpcnt=0
                tmp=money
                tmpcnt+=quarter_cnt
                tmp-=cur*quarter_cnt
                tmpcnt+=tmp//coins[-1]
                tmp%=coins[-1]
                tmpcnt+=tmp//coins[-2]
                tmp%=coins[-2]
                cnt=min(cnt,tmpcnt)
            money=tmp
            ans+=cnt
            coins.pop()
            coins.pop()
        else:
            ans+=money//cur
            money=(money%cur)
    print(ans)

