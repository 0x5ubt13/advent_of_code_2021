#!/usr/bin/python3D

def get_input(filename):
    with open(filename, "r") as f:
        return [[int(point) for point in line.strip()] for line in f]

def solve(matrix):
    low_points_risk = 0
    for row in range(len(matrix)):

        for column in range(len(matrix[row])):

            if column == 0: #First of the row
                
                if row == 0: # First row, first column
                    num, down, right = matrix[row][column], matrix[row+1][column], matrix[row][column+1]
                    print(num, down, right)
                    if down > num and right > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1
                elif row == len(matrix) -1: # Bottom row (last row), first number
                    num, up, right = matrix[row][column], matrix[row-1][column], matrix[row][column+1]
                    print(num, up, right)
                    if up > num and right > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1
                else:
                    num, up, down, right = matrix[row][column], matrix[row-1][column], matrix[row+1][column], matrix[row][column+1]
                    print(num, up, down, right)
                    if up > num and down > num and right > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1
                    

            elif column == (len(matrix[row])-1): # Last of the row
                
                if row == 0: # first row, Last column
                    num, down, left = matrix[row][column], matrix[row+1][column], matrix[row][column-1]
                    print(num, down, left)
                    if left > num and down > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1
                
                elif row == len(matrix) -1: # Bottom row (last row), last number
                    num, up, left = matrix[row][column], matrix[row-1][column], matrix[row][column-1]
                    print(num, up, left)
                    if up > num and right > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1

                else: # Any middle
                    num, up, down, left = matrix[row][column], matrix[row-1][column], matrix[row+1][column], matrix[row][column-1]
                    print(num, up, down, left)
                    if up > num and down > num and left > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1


            else: # Any other middle number

                if row == 0: # First row, no up numbers
                    num, down, left, right = matrix[row][column], matrix[row+1][column], matrix[row][column-1], matrix[row][column+1]
                    print(num, down, left, right)
                    if left > num and down > num and right > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1

                elif row == len(matrix) -1: # Bottom row
                    num, up, left, right = matrix[row][column], matrix[row-1][column], matrix[row][column-1], matrix[row][column+1]
                    print(num, up, left, right)
                    if up > num and left > num and right > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1

                else:
                    num, up, down, left, right = matrix[row][column], matrix[row-1][column], matrix[row+1][column], matrix[row][column-1], matrix[row][column+1]
                    print(num, up, down, left, right)
                    if up > num and down > num and left > num and right > num:
                        print(f"match at matrix[{row}][{column}] -> {num}", )
                        low_points_risk += num + 1

    # Part 2:
    # boilerplate: num, up, down, left, right = matrix[row][column], matrix[row-1][column], matrix[row+1][column], matrix[row][column-1], matrix[row][column+1]
    for row in range(len(matrix)):
        for col in range(len(matrix[row])):
            print("part2:")
            num = matrix[row][col]
            down, right = matrix[row+1][col-1], matrix[row][col+1]
            print(nu, down, right)
            if matrix[row][col] != 9: # Main part 2 loop, 9s are discarded:
                num = matrix[row][col]
                basin, basins, ignore = 0, [], []
                if column == 0: #First of the row
                    if row == 0: # First row, first column, first number of the matrix [0][0]
                        down, left, right = matrix[row-1][column], matrix[row+1][column], matrix[row][column-1], matrix[row][column+1]
                        print(down)
                        # Check adjacent numbers
                        if down == num+1 or down == num-1:
                            
                            basins.append(set(row, col))
                            basin = True
                            while basin:
                                basin+=1


                    
                #     elif row == len(matrix) -1: # Bottom row (last row), first number
                    
                #     else: # First num of any middle row

                # elif column == (len(matrix[row])-1): # Last of the row
                #     if row == 0: # first row, Last column

                #     elif row == len(matrix) -1: # Bottom row (last row), last number

                #     else: # Any middle

                
                # else: # Any other middle number
                #     if row == 0: # First row, no numbers above
                    
                #     elif row == len(matrix) -1: # Bottom row, no numbers below

                #     else: # The majority of middle numbers



    print("Part 1:", low_points_risk)


if __name__ == '__main__':
    solve(get_input("./day_9/9.test"))
