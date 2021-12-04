#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        return [line.strip() for line in f]


def solve(data):
    drawn_numbers = data[0:1]
    playing_boards = data[1:]
    parsed_boards = {}
    x = -1
    for line in range(len(playing_boards)):
        print(f"Playing boards line [{line}]:", playing_boards[line])
        # new line:
        if playing_boards[line] == '':
            x += 1
            parsed_boards[x] = {}
        else:
            y = 0    
            parsed_boards[x][y] = []
            for num in playing_boards[line].split(" "):
                print(num)
                y += 1


if __name__ == '__main__':
    data = get_input('./day_4/day_4_input.txt')
    solve(data)