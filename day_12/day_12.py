#!/usr/bin/python3


def get_input(filename):
    with open(filename, 'r') as f:
        return [line.strip().split('-') for line in f]

def possibilities(data):
    possibilities = []
    for i in data:
        for j in i:
            if j not in possibilities:
                possibilities.append(j)
    return possibilities


def part_1(data):
    paths = []
    x = 1
    for i in data:
        print(f"Path {x}")
        print(f"Pos {i}")
        for j in i:
            print(j)
        x += 1


if __name__ == '__main__':
    data = get_input("./day_12/12.test")
    part_1(data)
    