import sys
from collections import deque

n = int(sys.stdin.readline().rstrip())

q = deque()

for i in range(1, n + 1):
    q.append(i)

while len(q) > 1:
    q.popleft()
    q.append(q.popleft())


for i in q:
    print(i)
    break