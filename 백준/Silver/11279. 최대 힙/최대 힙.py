import heapq
from sys import stdin

if __name__ == '__main__':
    n = int(stdin.readline())
    arr = []

    for i in range(n):
        num = int(stdin.readline())
        if num == 0:
            if len(arr) == 0:
                print(0)
            else:
                print(-(heapq.heappop(arr)))
        else:
            heapq.heappush(arr, -num)