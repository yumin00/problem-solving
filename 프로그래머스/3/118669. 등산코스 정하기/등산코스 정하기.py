import heapq

def solution(n, paths, gates, summits):
    graph = [[] for _ in range(n+1)]
    
    for s, e, d in paths:
        graph[s].append((e, d))
        graph[e].append((s, d))
    
    INF = int(1e9)
    distance = [INF] * (n+1)
    queue = []
    for i in gates:
        distance[i] = 0
        heapq.heappush(queue, (0, i))
        
    while queue:
        d, s = heapq.heappop(queue)
        
        if distance[s] < d or s in summits:
            continue
        
        
        for i in graph[s]:
            node = i[0]
            dis = i[1]
            
            if distance[node] > max(d, dis):
                distance[node] = max(d, dis)
                heapq.heappush(queue, (distance[node], node))

    result = [0, INF]
    
    for node in sorted(summits):
        if distance[node] < result[1]:
            result[0] = node
            result[1] = distance[node]        
    
    
    
    
    
    return result