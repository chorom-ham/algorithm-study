def solution(n):
    div=0
    x=len(n)/2
    a1=[]
    print(x)
    while(int(x)>=div):
        div+=1
        a=0
        string=''
        count=1
        count2=div
        while(True):
            if n[a:a+div]==n[a+count2:a+div+count2]:
                count+=1
                count2+=div
            else:
                if count==1:
                    string+=n[a:a+div]
                    a=a+count2
                    count2=div
                else:
                    string+=str(count)+n[a:a+div]
                    count=1
                    a=a+count2
                    count2=div
            if a>=len(n):
                break
        a1.append(len(string))
                    
    return min(a1)


solution("aabbbacccc")