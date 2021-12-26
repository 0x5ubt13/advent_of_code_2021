#!/usr/bin/python3

def get_input():
    with open('./day_11/11.in', 'r') as f:
        return [[int(row.strip()) for row in column.strip()] for column in f]


def plus_1(grid):
    # First, the energy level of each octopus increases by 1.
    for row in range(len(grid)):
        for col in range(len(grid)):
            grid[row][col] += 1   


def pepinazu(grid, flashed):
    # Then, any octopus with an energy level greater than 9 flashes. 
    # This increases the energy level of all adjacent octopuses by 1, including octopuses that are diagonally adjacent. 
    # If this causes an octopus to have an energy level greater than 9, it also flashes. 
    # This process continues as long as new octopuses keep having their energy level increased beyond 9. 
    # (An octopus can only flash at most once per step.)
    
    for row in range(len(grid)):
        for col in range(len(grid)):
            if grid[row][col] > 9 and str(str(row)+str(col)) not in flashed:
                flashed.append(str(str(row)+str(col)))
                # Primera fila
                if row == 0:
                    if col == 0:
                        grid[row+1][col] += 1
                        grid[row][col+1] += 1
                        grid[row+1][col+1] += 1
                    elif col == 9:
                        grid[row+1][col] += 1
                        grid[row][col-1] += 1
                        grid[row+1][col-1] += 1
                    else:
                        grid[row][col-1] += 1
                        grid[row+1][col-1] += 1
                        grid[row+1][col] += 1
                        grid[row+1][col+1] += 1
                        grid[row][col+1] += 1
                # Ultima fila
                elif row == 9:
                    if col == 0:
                        grid[row-1][col] += 1
                        grid[row][col+1] += 1
                        grid[row-1][col+1] += 1
                    elif col == 9:
                        grid[row][col-1] += 1
                        grid[row-1][col] += 1
                        grid[row-1][col-1] += 1
                    else:
                        grid[row][col-1] += 1
                        grid[row-1][col-1] += 1
                        grid[row-1][col] += 1
                        grid[row-1][col+1] += 1
                        grid[row][col+1] += 1
                # Cualquiera del medio     
                else:
                    if col == 0:
                        grid[row-1][col] += 1
                        grid[row-1][col+1] += 1
                        grid[row][col+1] += 1
                        grid[row+1][col] += 1
                        grid[row+1][col+1] += 1
                    elif col == 9:
                        grid[row-1][col] += 1 # arriba
                        grid[row-1][col-1] += 1 # arriba izda
                        grid[row][col-1] += 1 # izda
                        grid[row+1][col] += 1 # abajo
                        grid[row+1][col-1] += 1 # abajo izda
                    else:
                        grid[row-1][col] += 1 # arriba
                        grid[row-1][col+1] += 1 # arriba dcha
                        grid[row][col+1] += 1 # dcha
                        grid[row+1][col+1] += 1 # abajo dcha
                        grid[row+1][col] += 1 # abajo
                        grid[row+1][col-1] += 1 # abajo izda
                        grid[row][col-1] += 1 # izda
                        grid[row-1][col-1] += 1 # arriba izda


def resetea(grid, step):
    part_2_counter = 0
    for row in range(len(grid)):
        for col in range(len(grid)):
            if grid[row][col] > 9:
                part_2_counter += 1
                if part_2_counter == 100:
                    print(f"Part 2: {step}")
                grid[row][col] = 0


if __name__ == '__main__':
    counter = 0
    grid = get_input() # pilla el input

    for i in range(1, 1001): # pasos (objetivo parte 1: 100)
        #print(f"Paso {i}:")
        plus_1(grid) # suma 1

        flashed = [] # lista de guarros 
        for j in range(15):
            pepinazu(grid, flashed) # si mas de 9, pepinazo

        resetea(grid, i) # si mas de 9 = 0
        for i in flashed:
            counter += 1

        # for line in grid:
        #     print(line)
        # else: print("")

    print(counter)




