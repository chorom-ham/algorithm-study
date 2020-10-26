#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int compression(string s, int unit){
    string temp = s.substr(0, unit);
    
    string str=""; 
    int same = 1;
    
    for(int i = unit; i< s.size(); i+=unit){
        if(temp == s.substr(i,unit)){
            same++;
        }else {
            if(same != 1) str += to_string(same);
            str += temp;  
            temp = s.substr(i,unit);
            same = 1;
        }
    }
    if(same != 1) str += to_string(same); 
    str += temp;  

    return str.length();
}

int solution(string s) {
    int answer=s.size();

    for(int unit = 1; unit <= s.size()/2 ; unit++){
        answer = min(answer, compression(s, unit));
    }
    
    return answer;
}