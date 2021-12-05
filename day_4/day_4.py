#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        return [line.strip() for line in f]


def make_boards(data):
    drawn_numbers = data[0].split(',')
    playing_boards = data[1:]
    parsed_boards = {}

    # Making the boards
    x, y = -1, -1
    for line in range(len(playing_boards)):
        if len(playing_boards[line]) == 0:
            x += 1
            parsed_boards[x] = {}
            for i in range(5):
                parsed_boards[x][i] = []
        else:   
            for num in playing_boards[line].split(" "):
                if num.isdigit():
                    parsed_boards[x][y].append(num)
        y += 1
        if y % 5 == 0:
            y = 0 

    return drawn_numbers, parsed_boards

def part_1(drawn_numbers, parsed_boards):
    # Drawing numbers
    playing = True
    while playing == True:
        for drawn in range(len(drawn_numbers)):
            x, y = 0, 0
            for board, board_values_dict in parsed_boards.items():
                for columns, rows in board_values_dict.items():
                    z = 0
                    for number in rows:
                        if drawn_numbers[drawn] == rows[z]:
                            rows[z] = 100
                            if rows.count(100) == 5 or int(board_values_dict[0][z]) + int(board_values_dict[1][z]) + int(board_values_dict[2][z]) + int(board_values_dict[3][z]) + int(board_values_dict[4][z]) == 500:
                                result = 0
                                last = int(number)
                                for col, row in board_values_dict.items():
                                    for num in row:
                                        if int(num) != 100:
                                            result += int(num)
                                print("Part 1: ", result * last)
                                print(f"Board {board} complete")
                                playing = False
                                break
                            
                            if playing == False:
                                break
                        z += 1
                    if playing == False:
                        break
                    y += 1
                    if y % 5 == 0:
                        y = 0
                if playing == False:
                    break
                x += 1
                if x % 5 == 0:
                    x = 0
            if playing == False:
                break
        break

def part_2(drawn_numbers, parsed_boards):
    # Drawing numbers
    def solve(drawn_numbers, parsed_boards):
        playing = True
        while playing == True:
            d = -1
            for drawn in range(len(drawn_numbers)):
                d += 1
                x, y = 0, 0
                for board, board_values_dict in parsed_boards.items():
                    for columns, rows in board_values_dict.items():
                        z = 0
                        for number in rows:
                            if drawn_numbers[drawn] == rows[z]:
                                rows[z] = 100
                                if rows.count(100) == 5 or int(board_values_dict[0][z]) + int(board_values_dict[1][z]) + int(board_values_dict[2][z]) + int(board_values_dict[3][z]) + int(board_values_dict[4][z]) == 500:
                                    result = 0
                                    last = int(number)
                                    for col, row in board_values_dict.items():
                                        for num in row:
                                            if int(num) != 100:
                                                result += int(num)
                                    part2 = result * last
                                    # print(f"Board {board} complete")
                                    playing = False
                                    break
                                
                                if playing == False:
                                    break
                            z += 1
                        if playing == False:
                            break
                        y += 1
                        if y % 5 == 0:
                            y = 0
                    if playing == False:
                        break
                    x += 1
                    if x % 5 == 0:
                        x = 0
                if playing == False:
                    break
            break
            
        del parsed_boards[board]    
        return drawn_numbers[d:], parsed_boards, part2
    
    # print(f"drawn numbers: {len(drawn_numbers)}, boards: {len(parsed_boards)}")
    while len(parsed_boards) > 0:
        drawn_numbers, parsed_boards, solution = solve(drawn_numbers, parsed_boards)
        # print(f"drawn numbers: {len(drawn_numbers)}, boards: {len(parsed_boards)}")
    print("Part 2: ", solution)
    

if __name__ == '__main__':
    data = get_input('./day_4/day_4_input.txt')
    d, l = make_boards(data)
    part_1(d, l)
    d, l = make_boards(data)
    part_2(d, l)

