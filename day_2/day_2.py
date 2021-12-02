#!/usr/bin/python3

def solve():
    with open("./day_2/day_2_input.txt", "r") as f:
        directions = [{line.strip().split(" ")[0]:int(line.strip().split(" ")[1])} for line in f]

    horizontal, aim, depth_1, depth_2 = 0, 0, 0, 0
    for line in directions:
        for direction, instruction in line.items():
            if direction.startswith("f"):
                horizontal += instruction
                depth_2 += aim * instruction
            elif direction.startswith("d"):
                depth_1 += instruction
                aim += instruction
            elif direction.startswith("u"):
                depth_1 -= instruction
                aim -= instruction
    print(f'Part 1 = {horizontal * depth_1} | Part 2 = {horizontal * depth_2}')

if __name__ == "__main__":
    solve()