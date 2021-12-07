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
        x1, y1, x2, y2 = int(initial[0]), int(initial[1]), int(final[0]), int(final[1])
        
        if x1 == x2 or y1 == y2:
            if x1 == x2:
                if y1 < y2:
                    for coord in range(y1, y2 +1):
                        diagram[(x1, coord)] += 1
                elif y1 > y2:
                    for coord in range(y1, y2 -1, -1):
                        diagram[(x1, coord)] += 1
            elif y1 == y2:
                if x1 < x2:
                    for coord in range(x1, x2 +1):
                        diagram[(coord, y1)] += 1
                elif x1 > x2:
                    for coord in range(x1, x2 -1, -1):
                        diagram[(coord, y1)] += 1
                        # print(f"diagram[(x: {coord}, y: {y1})] += 1")
        # Part 2
        else: #x1 != x2 and y1 != y2: # diagonal
            if x1 > x2 and y1 > y2: # Si son todos pabajo
                for coord in range(x1, x2 -1, -1):
                    if x1 > y1: # si x es mayor que y
                        diagram[(coord, coord - (x1 - y1))] += 1
                    elif x1 < y1: # si x es menor que y
                        diagram[(coord - (y1 - x1), coord)] += 1
                        print(f"initial: x = {initial[0]}, y = {initial[1]} , final: x = {final[0]}, y = {final[1]}")
                        print(f"Initial: {ins[0]}, final: {ins[1]}")
                        print(f"diagram[(x: {coord - (y1 - x1)}, y: {coord})] += 1")

                        
            elif x1 < x2 and y1 < y2: # Si son todos para arriba
                for coord in range(x1, x2):
                    if x1 < y1:
                        diagram[(coord, coord)] += 1
                        
                        # print(f"initial: x = {initial[0]}, y = {initial[1]} , final: x = {final[0]}, y = {final[1]}")
                        # print(f"Initial: {ins[0]}, final: {ins[1]}")
                        # print(f"diagram[(x: {coord - (y1 - x1)}, y: {coord})] += 1")
                pass 
            elif x1 < x2 and y1 > y2: # Si es x para arriba e y para abajo
                pass
            elif x1 > x2 and y1 < y2: # Si es x para abajo e y para arriba
                pass
                # print(f"initial: x = {initial[0]}, y = {initial[1]} , final: x = {final[0]}, y = {final[1]}")
                # print(f"Initial: {ins[0]}, final: {ins[1]}")
            elif x1 > x2 and y1 < y2:
                pass
                # print(f"diagram[(x: {x1}, y: {coord})] += 1")
            else:
                pass

                    



            #     for coord in range(y1, y2 +1):
            #         diagram[(x1, coord)] += 1
            # elif y1 > y2:
            #     for coord in range(y1, y2 -1, -1):
            #         diagram[(x1, coord)] += 1

            # elif y1 == y2:
            #     if x1 < x2:
            #         for coord in range(x1, x2 +1):
            #             diagram[(coord, y1)] += 1

            #     elif x1 > x2:
            #         for coord in range(x1, x2 -1, -1):
            #             diagram[(coord, y1)] += 1

    # result = 0
    # for v in diagram.values():
    #     if v >= 2:
    #         result += 1

    # print(result)


if __name__ == '__main__':
    data = get_input("./day_5/day_5_input.txt")
    # data = get_input("./day_5/day_5_test.txt")
    part_1(draw_map(), data)