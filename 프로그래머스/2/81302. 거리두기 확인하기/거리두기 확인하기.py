def is_valid(place):
    people = [(i, j) for i in range(5) for j in range(5) if place[i][j] == 'P']


    for i in range(len(people)):
        for j in range(i + 1, len(people)):
            r1, c1 = people[i]
            r2, c2 = people[j]
            distance = abs(r1 - r2) + abs(c1 - c2)

            if distance <= 2:
                
                if r1 == r2:
                    if place[r1][min(c1, c2) + 1] != 'X':
                        return 0
                elif c1 == c2:
                    if place[min(r1, r2) + 1][c1] != 'X':
                        return 0
                else:
                    if not (place[r1][c2] == 'X' and place[r2][c1] == 'X'):
                        return 0
    return 1

def solution(places):
    return [is_valid(place) for place in places]