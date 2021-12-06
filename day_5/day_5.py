#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        return [line.strip() for line in f]

def draw_map():
    diagram = {}
    coords = [(x, y) for x in range(1000) for y in range(1000)]
    for coord in coords:
        diagram[coord] = 0
    
    return diagram

def part_1(diagram, data):
    for coords in data:
        ins = coords.split(' -> ')
        initial = ins[0].split(',')
        final = ins[1].split(',')
        # print(f"initial: x = {initial[0]}, y = {initial[1]} , final: x = {final[0]}, y = {final[1]}")
        # print(f"Initial: {ins[0]}, final: {ins[1]}")
        initial_x, initial_y, final_x, final_y = int(initial[0]), int(initial[1]), int(final[0]), int(final[1])
        
        if initial_x == final_x or initial_y == final_y:
            if initial_x == final_x:
                if initial_y < final_y:
                    for coord in range(initial_y, final_y +1):
                        diagram[(initial_x, coord)] += 1
                        # print(f"diagram[(x: {initial_x}, y: {coord})] += 1")
                elif initial_y > final_y:
                    for coord in range(initial_y, final_y -1, -1):
                        diagram[(initial_x, coord)] += 1
                        # print(f"diagram[(x: {initial_x}, y: {coord})] += 1")
            elif initial_y == final_y:
                if initial_x < final_x:
                    for coord in range(initial_x, final_x +1):
                        diagram[(coord, initial_y)] += 1
                        # print(f"diagram[(x: {coord}, y: {initial_y})] += 1")
                elif initial_x > final_x:
                    for coord in range(initial_x, final_x -1, -1):
                        diagram[(coord, initial_y)] += 1
                        # print(f"diagram[(x: {coord}, y: {initial_y})] += 1")
        # Part 2
        elif initial_x != final_x and initial_y != final_y:
            if initial_y < final_y:
                if initial_x < final_x:
                    for coord in range(initial_x, final_x +1):
                        if initial_x < initial_y:
                            diagram[(coord, (initial_y - initial_x) + coord)] += 1
                            # print(f"diagram[(x: {coord}, y: {(initial_y - initial_x) + coord})] += 1")
                        elif initial_x > initial_y:
                            diagram[(coord, (initial_x - initial_y))] += 1
                            # print(f"diagram[(x: {coord}, y: {(initial_x - initial_y) + coord})] += 1")
                elif initial_x > final_x:
                    print(f"initial: x = {initial[0]}, y = {initial[1]} , final: x = {final[0]}, y = {final[1]}")
                    print(f"Initial: {ins[0]}, final: {ins[1]}")
                    for coord in range(initial_y, final_y +1):
                        if initial_x < initial_y:    
                            diagram[(initial_x, coord)] += 1
                            print(f"diagram[(x: {coord}, y: {coord})] += 1")
                            # print(f"diagram[(x: {initial_x}, y: {coord})] += 1")

                    



            #     for coord in range(initial_y, final_y +1):
            #         diagram[(initial_x, coord)] += 1
            # elif initial_y > final_y:
            #     for coord in range(initial_y, final_y -1, -1):
            #         diagram[(initial_x, coord)] += 1

            # elif initial_y == final_y:
            #     if initial_x < final_x:
            #         for coord in range(initial_x, final_x +1):
            #             diagram[(coord, initial_y)] += 1

            #     elif initial_x > final_x:
            #         for coord in range(initial_x, final_x -1, -1):
            #             diagram[(coord, initial_y)] += 1

    # result = 0
    # for v in diagram.values():
    #     if v >= 2:
    #         result += 1

    # print(result)


if __name__ == '__main__':
    data = get_input("./day_5/day_5_input.txt")
    # data = get_input("./day_5/day_5_test.txt")
    part_1(draw_map(), data)