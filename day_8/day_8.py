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

def parse_signals(raw_instructions, sample):
    keys, number = {}, []
    for char in range(10):
        keys[char] = ''
    instructions = [i for i in raw_instructions.strip().split(" ")]
    print(f"BEFORE = keys: {keys}, number: {number}, instructions: {instructions}, sample: {sample}")

    sorting = True
    while sorting:
    # First run to fill up the easy ones
        if len(instructions) == 0:
            sorting = False
        for ins in instructions:
            print(f"for ins({ins}) in instructions({instructions})")
            print(f"current keys: {keys}")
            if len(ins) == 2: # 1
                keys[1] += ins
                print(f"{ins} found to be in keys[1]: {keys[1]}")
                instructions.remove(ins)
            elif len(ins) == 3: # 7
                keys[7] += ins
                print(f"{ins} found to be in keys[7]: {keys[7]}")
                instructions.remove(ins)
            elif len(ins) == 4: # 4
                keys[4] += ins
                print(f"{ins} found to be in keys[4]: {keys[4]}")
                instructions.remove(ins)
            elif len(ins) == 7: # 8
                keys[8] += ins
                print(f"{ins} found to be in keys[8]: {keys[8]}")
                instructions.remove(ins)
        # Second run to get the rest
        for ins in instructions:
            if len(ins) == 6: # 6, 9, 0
                ding = 0
                for char in ins:
                    for k, v in keys.items():
                        if k == 4:
                            for ch in v:
                                if ch == char:
                                    ding += 1
                if ding == 4 and len(keys[9]) < 1:
                    keys[9] = ins
                    ding = 0
                    print(f"{ins} found to be in keys[9]: {keys[9]}")
                    instructions.remove(ins)
                else: 
                    for char in ins:
                        for k, v in keys.items():
                            if k == 1:
                                for ch in v:
                                    if ch == char:
                                        ding += 1
                    if ding == 2 and len(keys[0]) < 1:
                        keys[0] += ins
                        ding = 0
                        print(f"{ins} found to be in keys[0]: {keys[0]}")
                        instructions.remove(ins)
                    else:
                        ding = 0
                        for char in ins:
                            for k, v in keys.items():
                                if k == 8:
                                    for ch in v:
                                        if ch == char:
                                            ding += 1
                        if ding == 6 and len(keys[6]) < 1:
                            keys[6] += ins
                            ding = 0
                            print(f"{ins} found to be in keys[6]: {keys[6]}")
                            instructions.remove(ins)
            elif len(ins) == 5: # 5, 3, 2
                ding = 0
                for char in ins:
                    for k, v in keys.items():
                        if k == 6:
                            for ch in v:
                                if ch == char:
                                    ding += 1
                if ding == 5 and len(keys[5]) < 1:
                    keys[5] += ins
                    ding = 0
                    print(f"{ins} found to be in keys[5]: {keys[5]}")
                    instructions.remove(ins)
                else:
                    ding = 0
                    for char in ins:
                        for k, v in keys.items():
                            if k == 9:
                                for ch in v:
                                    if ch == char:
                                        ding += 1
                    if ding == 5 and len(keys[3]) < 1:
                        keys[3] += ins
                        ding = 0
                        print(f"{ins} found to be in keys[3]: {keys[3]}")
                        instructions.remove(ins)
                    else:
                        for char in ins:    
                            for k, v in keys.items():
                                if k == 9:
                                    for ch in v:
                                        if ch == char:
                                            ding += 1
                        if ding == 4 and len(keys[2]) < 1:
                            keys[2] = ins
                            print(f"{ins} found to be in keys[2]: {keys[2]}")
                            instructions.remove(ins)
                        else:
                            continue

        else:
            continue        
    print(f"AFTER = keys: {keys}, number: {number}, instructions: {instructions}, sample: {sample}")
    # for str in sample:
    #     for k, v in keys.items():
    #         if str == v:
    #             number.append(k)

    # print(f"keys: {keys}")
    # print(number)
    return number


def solve(data):
    one, two = 0, 0
    for line in data:
        two_pool = ''
        # print(line[1].split(' '))
        instructions = line[0]
        for signals in line[1].split(' '):
            # print(signals)
            if parse_signals_one(signals):
                one += 1
            
            parse_signals(instructions, signals)
        print(''.join(two_pool))
        # two += int(''.join(two_pool))        
    
    print(f"Part 1: {one} | Part 2: {two}")


if __name__ == '__main__':
    # data = get_input("./day_8/day_8_input.txt")
    data = get_input("./day_8/day_8_test.txt")
    solve(data)