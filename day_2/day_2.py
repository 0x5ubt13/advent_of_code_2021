#!/usr/bin/python3

def get_input():
    with open("./day_2/day_2_input.txt", "r") as f:
        return [line.strip() for line in f]


def part_1(instructions):
    horizontal = 0
    depth = 0
    for line in instructions:
        instruction = line.split(" ")
        if line[0].startswith("f"):
            horizontal += int(instruction[1])
        elif line[0].startswith("d"):
            depth += int(instruction[1])
        elif line[0].startswith("u"):
            depth -= int(instruction[1])
    return horizontal * depth


def part_2(instructions):
    aim = 0
    horizontal = 0
    depth = 0
    for line in instructions:
        instruction = line.split(" ")
        if line[0].startswith("f"):
            horizontal += int(instruction[1])
            depth += aim * int(instruction[1])
        elif line[0].startswith("d"):
            aim += int(instruction[1])
        elif line[0].startswith("u"):
            aim -= int(instruction[1])
    
    return horizontal * depth


if __name__ == "__main__":
    print(f"Part 1 = {part_1(get_input())} | Part 2 = {part_2(get_input())}")