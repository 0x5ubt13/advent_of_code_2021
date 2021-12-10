chars = {
    '(': ')',
    '{': '}',
    '[': ']',
    '<': '>'
}
match = []

def get_input(filename):
    with open(filename, "r") as f:
        return [line.strip() for line in f]

def get_points(result):
    points = 0
    for ch in result:
        if ch == ')': points += 3
        if ch == ']': points += 57
        if ch == '}': points += 1197
        if ch == '>': points += 25137
    return points
        
def main(data):
    
    def part_1():
        to_delete = []
        x = -1
        for line in data:
            x += 1
            parse = []
            for char in line:
                if char in chars.keys():
                    parse.append(char)
                else:
                    if list(chars.keys())[list(chars.values()).index(char)] == parse[-1]:
                        parse.pop()
                    elif list(chars.keys())[list(chars.values()).index(char)] != parse[-1]:
                        match.append(char)
                        print("line", x)
                        to_delete.append(x)
                        break

        for i in to_delete[::-1]:
            del data[i]

        return get_points(match)
    

    def part_2():
        print(len(data))

    print(f"Part 1: {part_1()} | Part 2: {part_2()}")
        

if __name__ == '__main__':
    data = get_input("./day_10/10.in")
    # data = get_input("./day_10/10.test")
    main(data)