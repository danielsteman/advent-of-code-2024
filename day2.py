def is_safe(report: list[int]) -> bool:
    first_lvl = report[0]
    second_lvl = report[1]

    if abs(first_lvl - second_lvl) > 3 or first_lvl == second_lvl:
        return False

    if first_lvl < second_lvl:
        increasing = True
    else:
        increasing = False

    last_lvl = second_lvl

    for lvl in report[2:]:
        if lvl == last_lvl:
            return False
        if abs(lvl - last_lvl) > 3:
            return False
        if lvl < last_lvl and increasing:
            return False
        if lvl > last_lvl and not increasing:
            return False
        last_lvl = lvl

    return True

with open("data.txt", "r") as file:
    safe_report_count = 0
    for line in file:
        report = list(map(int, line.strip().split(" ")))
        print(report)
        print(is_safe(report))

        if is_safe(report):
            safe_report_count += 1

    print(safe_report_count)