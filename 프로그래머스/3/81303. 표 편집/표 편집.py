def solution(n, k, cmd):
    table = ['O'] * n
    deleted = []
    current = k

    prev = {i: i - 1 for i in range(n)}
    next = {i: i + 1 for i in range(n)}
    next[n - 1] = -1

    for command in cmd:
        if command[0] == 'U':
            X = int(command[2:])
            for _ in range(X):
                current = prev[current]
        elif command[0] == 'D':
            X = int(command[2:])
            for _ in range(X):
                current = next[current]
        elif command[0] == 'C':
            deleted.append(current)
            table[current] = 'X'

            if prev[current] != -1:
                next[prev[current]] = next[current]
            if next[current] != -1:
                prev[next[current]] = prev[current]

            current = next[current] if next[current] != -1 else prev[current]
        elif command[0] == 'Z':
            restore = deleted.pop()
            table[restore] = 'O'

            if prev[restore] != -1:
                next[prev[restore]] = restore
            if next[restore] != -1:
                prev[next[restore]] = restore

    return ''.join(table)
