import sys

n = int(sys.stdin.readline().rstrip())

def print_answer(ps):
    q = []
    for v in ps:
        if v == "(":
            q.append(v)
        else:
            if len(q) == 0:
                print("NO")
                return
            else:
                q.pop()

    if len(q) > 0:
        print("NO")
    else:
        print("YES")

    return

for _ in range(n):
    ps = sys.stdin.readline().rstrip()
    print_answer(ps)