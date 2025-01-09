from collections import deque

def get_document_print(m, document_list):
    queue = deque(enumerate(document_list))
    count = 0
    max_value = max(document_list)

    while queue:
        current = queue.popleft()

        if current[1] == max_value:
            count += 1
            if current[0] == m:
                print(count)
                return
            else:
                max_value = max(item[1] for item in queue)
        else:
            queue.append(current)


test_case = int(input())

for _ in range(test_case):
    n, m = map(int, input().split())
    document_list = list(map(int, input().split()))
    get_document_print(m, document_list)
