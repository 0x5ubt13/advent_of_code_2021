#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        data = [i for i in f]
    return [int(i) for i in data[0].split(',')]


def solve(data):
    total_fuel = {}
    for i in range(1, max(data)+1):
        total_fuel[i] = 0
    total_fuel_2 = total_fuel.copy()

    for i in range(1, max(data) + 1):
        for position in data:
            if i > position:
                total_fuel[i] += i - position
                for j in range(i - position):
                    total_fuel_2[i] += j + 1
            elif i < position:
                total_fuel[i] += position - i
                for j in range(position - i):
                    total_fuel_2[i] += j + 1
            else:
                total_fuel[i] += 0

    print(f"Part 1: {total_fuel[min(total_fuel, key=total_fuel.get)]}")
    print(f"Part 2: {total_fuel_2[min(total_fuel_2, key=total_fuel_2.get)]}")

if __name__ == '__main__':
    data = get_input("./day_7/day_7_input.txt")
    solve(data)

