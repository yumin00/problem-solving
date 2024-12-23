import sys
from collections import deque

k, n = sys.stdin.readline().rstrip().split()

people = deque()
result = []

for i in range(1, int(k)+1):
    people.append(i)

while people:
    for _ in range(int(n)-1):
        people.append(people.popleft())

    result.append(people.popleft())


print(str(result).replace('[', '<').replace(']', '>'))