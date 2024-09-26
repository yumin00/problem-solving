if __name__ == '__main__':
    n = int(input())
    queue = []
    answers = []

    for i in range(n):
        command = input().strip()
        if command.startswith('push'):
            _, value = command.split()
            queue.append(int(value))
        elif command == 'pop':
            if len(queue) == 0:
                answers.append(-1)
            else:
                answers.append(queue.pop(0))
        elif command == 'size':
            answers.append(len(queue))
        elif command == 'empty':
            answers.append(1 if len(queue) == 0 else 0)
        elif command == 'front':
            if len(queue) == 0:
                answers.append(-1)
            else:
                answers.append(queue[0])
        elif command == 'back':
            if len(queue) == 0:
                answers.append(-1)
            else:
                answers.append(queue[-1])

    for answer in answers:
        print(answer)
