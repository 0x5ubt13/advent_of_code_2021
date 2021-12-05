#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        return [line.strip() for line in f]

def draw_map():
    m = {}
    coords = [(x, y) for x in range(10) for y in range(10)]
    for coord in coords:
        m 

def part_1(m, data):
    for i in data:
        ins = i.split(' -> ')
        initial = ins[0].split(',')
        final = ins[1].split(',')
        print(f"initial: x = {initial[0]}, y = {initial[1]} , final: x = {final[0]}, y = {final[1]}")
        print(f"Initial: {ins[0]}, final: {ins[1]}")
        initial_x, initial_y, final_x, final_y = int(initial[0]), int(initial[1]), int(final[0]), int(final[1])
        
        print(m)

        # if initial_x == final_x or initial_y == final_y:
        #     if initial_x < final_x:
        #         for coord in range(initial_x, final_x):
        #             m[f"x{coord}"] += 1
        #     else:
        #         for coord in range(final_x, initial_x):
        #             m[f"x{coord}"] += 1
        #     if initial_y < final_y:
        #         for coord in range(initial_y, final_y):
        #             m[f"y{coord}"] += 1
        #     else:
        #         for coord in range(final_y, initial_y):
        #             m[f"y{coord}"] += 1
        
        
    # result = 0
    # for v in m.values():
    #     if v >= 2:
    #         result += 1
    
    # print(result)

        


if __name__ == '__main__':
    # data = get_input("./day_5/day_5_input.txt")
    data = get_input("./day_5/day_5_test.txt")
    part_1(draw_map(), data)