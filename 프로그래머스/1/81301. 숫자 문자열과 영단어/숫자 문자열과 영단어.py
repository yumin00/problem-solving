def solution(s):
    l = len(s)
    idx = 0
    result = []
    while idx < l:
        hi = s[idx]
        if s[idx].isdigit():
            result.append(int(s[idx]))
            idx += 1
        else:
            if s[idx] == "z":
                result.append(0)
                idx += 4
            elif s[idx] == "o":
                result.append(1)
                idx += 3
            elif s[idx] == "t" and s[idx+1] == "w":
                result.append(2)
                idx += 3
            elif s[idx] == "t" and s[idx+1] == "h":
                result.append(3)
                idx += 5
            elif s[idx] == "f" and s[idx+1] == "o":
                result.append(4)
                idx += 4
            elif s[idx] == "f" and s[idx+1] == "i":
                result.append(5)
                idx += 4
            elif s[idx] == "s" and s[idx+1] == "i":
                result.append(6)
                idx += 3
            elif s[idx] == "s" and s[idx++1] == "e":
                result.append(7)
                idx += 5
            elif s[idx] == "e":
                result.append(8)
                idx += 5
            else:
                result.append(9)
                idx += 4

    return int(''.join(map(str, result)))