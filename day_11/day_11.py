#!/usr/bin/python3

with open('./day_11/11.test', 'r') as f:
    grid = [[int(row.strip()) for row in column.strip()] for column in f]

glowing = False
for step in range(1, 5):
    print('Step', step)

    for row in range(0, len(grid)):
        for col in range(len(grid)):
            if grid[col][col] == 0:
                glowing = True

    while glowing == True:
        for row in range(0, len(grid)):
            for col in range(len(grid)):
                check = 0

                # Calculate adjacent to 0's
                num = grid[row][col]
                if num == 0:
                    glowing = True
                    try: 
                        if grid[row-1][col] % 9 == 0: 
                            grid[row-1][col] = 0 
                            check += 1
                        else: grid[row-1][col]+= 1
                    except IndexError: pass                    
                    try:
                        if grid[row-1][col-1] % 9 == 0:
                            grid[row-1][col-1] = 0 
                            check += 1
                        else: grid[row-1][col-1] += 1
                    except IndexError: pass                    
                    try: 
                        if grid[row][col-1] % 9 == 0: 
                            grid[row][col-1] = 0
                            check += 1
                        else: grid[row][col-1] += 1
                    except IndexError: pass                    
                    try: 
                        if grid[row+1][col-1] % 9 == 0:
                            grid[row+1][col-1] = 0
                            check += 1
                        else: grid[row+1][col-1] += 1
                    except IndexError: pass                    
                    try: 
                        if grid[row+1][col] % 9 == 0:
                            grid[row+1][col] = 0
                            check += 1
                        else: grid[row+1][col] += 1
                    except IndexError: pass                    
                    try: 
                        if grid[row+1][col+1] % 9 == 0:
                            grid[row+1][col+1] = 0
                            check += 1
                        else: grid[row+1][col+1] += 1
                    except IndexError: pass                    
                    try: 
                        if grid[row][col+1] % 9 == 0: 
                            grid[row][col+1] = 0
                            check += 1
                        else: grid[row][col+1] += 1
                    except IndexError: pass
                    try: 
                        if grid[row-1][col+1] % 9 == 0:
                            grid[row-1][col+1] = 0
                            check += 1
                        else: grid[row-1][col+1] += 1
                    except IndexError: pass

                    if check == 0:
                        glowing = False

    # Start next step: add 1 to every number and convert 9's in 0's
    for row in range(0, len(grid)):
        for col in range(len(grid)):
            if grid[row][col] % 9 != 0 or grid[row][col] == 0:
               grid[row][col] += 1
            else:
                grid[row][col] = 0
    
    for line in grid: 
        print(line)
    
    
        




