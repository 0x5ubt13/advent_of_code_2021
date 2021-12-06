#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        return [line for line in f]


def solve(data, days):
    school = {}
    for i in range(10):
        school[i] = 0
    for fish in data[0].split(','):
        school[int(fish)] += 1

    for _ in range(days):
        for fish in range(10):
            # New cycle
            if fish == 0:
                school[9] += school[fish]
                school[7] += school[fish]
                school[fish] = 0
            else:
                school[fish - 1] += school[fish]
                school[fish] = 0
            
    print(sum(school.values()))

    
if __name__ == '__main__':
    data = get_input("./day_6/day_6_input.txt")
    part_1 = solve(data, 80)
    part_2 = solve(data, 256)
