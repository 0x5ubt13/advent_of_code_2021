#!/usr/bin/python3

def get_input(filename):
    with open(filename, 'r') as f:
        return f.read().strip().splitlines()


def solve(data):
    parts = [0, 0]

    for line in data:
        signals = line.split(" | ")[0].split()
        numbers = line.split(" | ")[1].split()
        
        # Part 2
        digitsets = [None] * 10

        for sig in signals:
            signal = set(sig)
            if len(signal) == 2:
                digitsets[1] = signal
            elif len(signal) == 4:
                digitsets[4] = signal
            elif len(signal) == 3:
                digitsets[7] = signal
            elif len(signal) == 7:
                digitsets[8] = signal

        for sig in signals:
            signal = set(sig)
            if len(signal) == 6:
                if (digitsets[7] | digitsets[4]).issubset(signal):
                    digitsets[9] = signal
                elif digitsets[1].issubset(signal):
                    digitsets[0] = signal
                else:
                    digitsets[6] = signal

            elif len(signal) == 5:
                if digitsets[1].issubset(signal):
                    digitsets[3] = signal
                elif len(signal.difference(digitsets[7] | digitsets[4])) == 1:
                    digitsets[5] = signal
                else:
                    digitsets[2] = signal

        output = ""
        for n in numbers:
            for d, digit in enumerate(digitsets):
                if set(n) == digit:
                    output += str(d)
                if len(n) in [2, 4, 3, 7]: 
                    parts[0] += 1 ## Part 1

        parts[1] += int(output) ## Part 2

    print(f"Part 1: {parts[0]} | Part 2: {parts[1]}")     


if __name__ == '__main__':
    data = get_input("./day_8/day_8_input.txt")
    # data = get_input("./day_8/day_8_test.txt")
    solve(data)