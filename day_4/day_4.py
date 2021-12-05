#!/usr/bin/python3

from typing import Match


def get_input(filename):
    with open(filename, 'r') as f:
        return [line.strip() for line in f]




def solve(data):
    drawn_numbers = data[0].split(',')
    playing_boards = data[1:]
    parsed_boards = {}

    # Making the boards
    x, y = -1, -1
    for line in range(len(playing_boards)):
        # print(f"Playing boards line [{line}]:", playing_boards[line])
        # print(f"value of y: {y}")
        # new line:
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

    # print(parsed_boards.items())
    # Drawing numbers
    playing = True
    part2 = True
    while playing == True or part2 == True:
        for drawn in range(len(drawn_numbers)):
            x, y = 0, 0
            # x = board, y = column, z = row index
            # print("Drawn:", drawn_numbers[drawn])
            for board, board_values_dict in parsed_boards.items():
                for columns, rows in board_values_dict.items():
                    # print(f"board {board}, column {columns}, row {rows}")
                    # print(x)
                    # print(y)
                    z = 0
                    for number in rows:
                        if drawn_numbers[drawn] == rows[z]:
                            rows[z] = 100
                            if rows.count(100) == 5:
                                result = 0
                                last = int(number)
                                for col, row in board_values_dict.items():
                                    for num in row:
                                        if int(num) != 100:
                                            result += int(num)
                                print("Part 1: ", result * last)
                                playing = False

                                # Part 2
                                del parsed_boards[board]
                                if len(parsed_boards.keys()) == 1:
                                    if rows.count(100) == 5:
                                        result = 0
                                        last = int(number)
                                        for col, row in board_values_dict.items():
                                            for num in row:
                                                if int(num) != 100:
                                                    result += int(num)
                                            print("Part 2: ", result * last)
                                            part2 = False
                                # print('LINEA! Last number = ' + number)
                                break
                            if int(board_values_dict[0][z]) + int(board_values_dict[1][z]) + int(board_values_dict[2][z]) + int(board_values_dict[3][z]) + int(board_values_dict[4][z]) == 500:
                                result = 0
                                last = int(number)
                                for col, row in board_values_dict.items():
                                    for num in row:
                                        if int(num) != 100:
                                            result += int(num)
                                print("Part 1: ", result * last)
                                playing = False
                                
                                # Part 2
                                del parsed_boards[board]
                                if len(parsed_boards.keys()) == 1:
                                    if int(board_values_dict[0][z]) + int(board_values_dict[1][z]) + int(board_values_dict[2][z]) + int(board_values_dict[3][z]) + int(board_values_dict[4][z]) == 500:
                                        result = 0
                                        last = int(number)
                                        for col, row in board_values_dict.items():
                                            for num in row:
                                                if int(num) != 100:
                                                    result += int(num)
                                            print("Part 2: ", result * last)
                                            part2 = False

                                # print('LINEA! Last number = ' + number)
                                break
                            if playing == False:
                                break
                        z += 1
                    if playing == False:
                        break
                # if drawn_numbers[drawn] in rows:
                #     board_values_dict[columns].remove(drawn_numbers[drawn])
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
    

if __name__ == '__main__':
    data = get_input('./day_4/day_4_input.txt')
    # test = get_input('./day_4/day_4_test_input.txt')
    solve(data)