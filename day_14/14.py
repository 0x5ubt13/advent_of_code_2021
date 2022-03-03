#/usr/bin/python3

def get_input(filename):
    with open(filename, "r") as f:
        return [line.strip() for line in f]

def solve(data):
    print(data)

if __name__ == '__main__':
    data = get_input("./day_14/14_test.txt")
    solve(data)