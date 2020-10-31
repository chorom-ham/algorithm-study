# 17472. 다리만들기 2

## 풀이 방법

다른 사람 java풀이를 참고했다...다시 풀어볼것...

각 섬마다 번호를 지정하고 다리를 만든 후

최소스패닝트리(MST)를 생성한다

```go
//  최소스패닝트리(Kruskal 알고리즘)
//  간선들이 오름차순으로 정렬 되어 있어야 한다


//  두 정점 연결
func union(x, y int) bool {
    xRoot := find(x)    // 정점 x의 부모 노드 반환
    yRoot := find(y)    // 정점 y의 부모 노드 반환

    // 둘의 부모가 같지 않다면 합치는 것 가능
	if xRoot != yRoot {
		parents[yRoot] = xRoot //   x나 y 중 편한 것을 선택
		return true
	}
	return false
}
// 해당 정점의 연결된 정점 찾기
func find(x int) int {
    //  처음 parents 들은 -1로 초기화
    //  0보다 작은 경우 부모 노드인 것이므로 x 반환
	if parents[x] < 0 {
		return x
    }

    // 아닌 경우 계속해서 부모를 찾기 위해 거슬러 올라감.
	parents[x] = find(parents[x])
	return parents[x]
}
```

## 어려웠던 부분 혹은 새로 알게 된 내용

여러 알고리즘이 복합적으로 사용되어 푸는 방법을 찾는 것 조차 어려웠다

최소스패닝트리(MST)문제는 이론은 알았지만 실제 구현 문제를

한번도 풀어보지 않아서 1197번 문제를 풀어보면서 방법을 익혔다

최소스패닝트리는 크게 2가지 알고리즘으로 만들 수 있는데

간선이 적을 땐 Kruskal 알고리즘, 정점이 적을 땐 Prim 알고리즘을 사용한다.
