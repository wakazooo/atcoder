def validate(N, Y):
    for x in reversed(range(N + 1)):
        if 10000 * x > Y:
            continue
        elif 10000 * x + 5000 * (N - x) < Y:
            break
        elif x == N and Y - 10000 * x == 0:
            return x, 0, 0
        else:
            Y = Y - 10000 * x
        for y in reversed(range(N - x + 1)):
            if 5000 * y > Y or N - x - y < 0:
                continue
            elif 10000 * x + 5000 * y + 1000 * (N - x - y) < Y:
                break
            elif Y - 5000 * y - 1000 * (N - x - y) == 0:
                return x, y, N - x - y
        Y = Y + 10000 * x
    return -1, -1, -1

input = list(map(int, input().split()))

x, y, z = validate(input[0], input[1])
print("{} {} {}".format(x, y, z))
