def solution(s):
    if len(s)==1:
        return 1
    length=[]
    for i in range(1,len(s)//2+1):
        tmpstr=s
        curstr=''
        cnt=1
        tmpLength=0
        while tmpstr:
            #처음경우 갱신
            if curstr=='':
                curstr=tmpstr[:i]
            else:
            #처음이 아닌경우
                #패턴이 일치하는경우
                #카운트업, 문자열 갱신
                if curstr==tmpstr[:i]:
                    cnt+=1
                else:
                #일치하지 않는경우
                #지금까지의 카운트+패턴의 길이 저장, 초기화, 문자열 갱신
                    if cnt==1:
                        tmpLength+=(len(curstr))
                    else:
                        tmpLength+=(len(curstr)+len(str(cnt)))
                    curstr=tmpstr[:i]
                    cnt=1
            tmpstr=tmpstr[i:]
            if not tmpstr:
                if cnt==1:
                        tmpLength+=(len(curstr))
                else:
                    tmpLength+=(len(curstr)+len(str(cnt)))
        length.append(tmpLength)
    
    return min(length)