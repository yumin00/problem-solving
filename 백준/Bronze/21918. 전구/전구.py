if __name__ == '__main__':
    n, m = map(int, input().split())
    lights = list(map(int, (input().split())))

    for _ in range(m):
        a, b, c, = map(int, input().split())
        if a == 1:
            lights[b-1] = c
        elif a == 2:
            idx = 0
            for i in lights:
                if b-1 <= idx <= c-1:
                    if lights[idx] == 0:
                        lights[idx] = 1
                    else:
                        lights[idx] = 0
                idx += 1
        elif a == 3:
            idx = 0
            for i in lights:
                if b - 1 <= idx <= c - 1:
                    lights[idx] = 0
                idx += 1
        elif a == 4:
            idx = 0
            for i in lights:
                if b - 1 <= idx <= c - 1:
                    lights[idx] = 1
                idx += 1

    idx = 0
    for i in lights:
        if idx == len(lights):
            print(i, end='')

        print(i, end=' ')
        idx += 1