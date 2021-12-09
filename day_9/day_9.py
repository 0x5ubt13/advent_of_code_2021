#!/usr/bin/python3

def get_input(filename):
    with open(filename, "r") as f:
        return [[int(point) for point in line.strip()] for line in f]

def solve(matrix):
    low_points_risk = 0
    for row in range(len(matrix)):
        for column in range(len(matrix[row])):
            try:
                num, up, down, left, right = matrix[row][column], matrix[row-1][column], matrix[row+1][column], matrix[row][column-1], matrix[row][column+1]
            except:
                pass
            if column == 0: #First of the row
                if row == 0: # First row, first column
                    if down < num and right < num:
                        low_points_risk += num + 1
                elif row == len(matrix) -1:
                    print(len(matrix), len(matrix) -1 ) 

            elif column == (len(matrix[row]) - 1): # Last of the row
                if row == 0: # first row, Last column
                    if left < num and down < num:
                        low_points_risk += num + 1

            else: # Any other middle number
                if row == 0: # First row, no up numbers
                    if left < num and down < num and right < num:
                        low_points_risk += num + 1
                



            #     print(f"first row: [{row}][{column}]", matrix[row][column])
            # if row == 1:
            #     print(f"second row: [row][column]", matrix[row][column])
            # if i - 1 >= 0 and j - 1 >= 0:
            #     print(data[i][j])


if __name__ == '__main__':
    solve(get_input("./day_9/9.test"))
