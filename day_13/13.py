#!/usr/bin/python3

def get_input(fn):
    with open(fn, "r") as f:
        return [line.strip() for line in f]


def make_grid(data):
    for row in range(len(data)):
        for col in range(len(data)):
            grid.append()





def main(data):

    grid = make_grid(data)



    part_1 = "solution"
    print(part_1)



if __name__ == "__main__":
    data = get_input("./day_13/13.in")
    main(data)