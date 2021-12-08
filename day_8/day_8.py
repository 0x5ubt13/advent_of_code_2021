#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        return [line.strip().split('|') for line in f]

def parse_signals_one(sample):
    if len(sample) == 2:
        return 1 
    elif len(sample) == 3:
        return 7
    elif len(sample) == 4:
        return 4
    elif len(sample) == 7:
        return 8

def parse_signals(sample):
    number = ''
    '''
acedgfb: 8
cdfbe: 5
gcdfa: 2
fbcad: 3
dab: 7
cefabd: 9
cdfgeb: 6
eafb: 4
cagedb: 0
ab: 1

 dddd                d                  
e    a          e          a
e    a
 ffff                f 
g    b          g           b
g    b  
 cccc                c
    '''
    if len(sample) == 2:
        number += '1' 
    elif len(sample) == 3:
        number += '7'
    elif len(sample) == 4:
        number += '4'
    elif len(sample) == 5:
        if 'e' in sample:
            number += '5'
        
        elif 'f' in sample:
                number += '3'
        elif 'g' in sample:
                number += '2'
    elif len(sample) == 6:
        if 'g' in sample:
            number += '6'
        else: 
            if 'f' in sample:
                number += '9'
            else: number += '0'
    elif len(sample) == 7:
        number += '8'
    
    return number


def solve(data):
    one, two = 0, 0
    for line in data:
        two_pool = ''
        print(line[1].split(' '))
        for signals in line[1].split(' '):
            print(signals)
            if parse_signals_one(signals):
                one += 1
            
            two_pool += parse_signals(signals)
        print(two_pool)
        two += int(''.join(two_pool))
        

    print(f"Part 1: {one} | Part 2: {two}")



if __name__ == '__main__':
    # data = get_input("./day_8/day_8_input.txt")
    data = get_input("./day_8/day_8_test.txt")
    solve(data)