#!/usr/bin/python3

def get_data(input_file):
    # Create dictionary with every consecutive digit to perform the operations with
    bin_list = []
    bin_dict = {}
    for i in range(12):
        bin_dict[i] = []
    
    with open(input_file, "r") as f:
        x = 0
        for line in f:
            for digit in range(len(line.strip())):
                bin_dict[digit].append(line[digit])
            x += 1
            bin_list.append(line.strip())

    return bin_dict, bin_list

def solve(bin_dict, bin_list):
    # Part 1: Determine most and least common bits
    most_common = []
    for i in range(len(bin_dict)):
        if bin_dict[i].count('0') >= bin_dict[i].count('1'):
            most_common.append('0')
        else:
            most_common.append('1')

    print(f"Part 1: {int(''.join(most_common), 2) * int(''.join(most_common).replace('0', 'x').replace('1', '0').replace('x', '1'), 2)}")

    # Part 2
    oxygen_valid_numbers = bin_list.copy()
    co2_valid_numbers = bin_list.copy()
    
    for column in range(len(oxygen_valid_numbers[0])):
        column_bits_oxygen = [i[column] for i in oxygen_valid_numbers]
        column_bits_co2 = [i[column] for i in co2_valid_numbers]

        if column_bits_oxygen.count('1') >= column_bits_oxygen.count('0'):
            oxygen_valid_numbers = [i for i in oxygen_valid_numbers if i[column] == '1']
        else:
            oxygen_valid_numbers = [i for i in oxygen_valid_numbers if i[column] == '0']
        
        if column_bits_co2.count('1') < column_bits_co2.count('0'):
            co2_valid_numbers = [i for i in co2_valid_numbers if i[column] == '1']
        else:
            co2_valid_numbers = [i for i in co2_valid_numbers if i[column] == '0']

        if len(oxygen_valid_numbers) == 1:
            break
        
        if len(co2_valid_numbers) == 1:
            break

    print(f"Part 2: {int(oxygen_valid_numbers[0], 2) * int(co2_valid_numbers[0], 2)}")

if __name__ == '__main__':
    d, l = get_data("./day_3/day_3_input.txt")
    solve(d, l)
    