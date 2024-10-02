import heapq
import sys

if __name__ == '__main__':
    input = sys.stdin.readline

    n = int(input())
    heap = []

    init_number = list(map(int, input().split())) # 첫 번째 라인 먼저 받기

    for num in init_number:
        heapq.heappush(heap, num)

    for i in range(n-1):
        numbers = list(map(int, input().split()))
        for num in numbers:
            if heap[0] < num:
                heapq.heappush(heap, num)
                heapq.heappop(heap)

    
    print(heap[0])