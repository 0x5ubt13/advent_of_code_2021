#!/usr/bin/python3

def get_input():
    with open("./day_1/day_1_input.txt", "r") as f:
        return [int(line.strip()) for line in f]

def part_1(depth):
    return sum([1 for i in range(len(depth)) if depth[i] > depth[i-1]])

def part_2(depth):
    return [(depth[i] + depth[i+1] + depth[i+2]) for i in range(len(depth) - 2)]

if __name__ == "__main__":
    print(f'The solution to part 1 is {part_1(get_input())}. The solution to part 2 is {part_1(part_2(get_input()))}.')