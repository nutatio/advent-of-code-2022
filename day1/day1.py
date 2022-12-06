def part_one():
    cur_sum, m = 0, 0
    while True:
        line = input()
        if line == "":
            m = max(m, cur_sum)
            cur_sum = 0
            continue
        if line == "amog":
            break
        cur_sum += int(line)
    print(m)



if __name__ == '__main__':
