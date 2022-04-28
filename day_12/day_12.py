#!/usr/bin/python3

from collections import defaultdict, deque
import time

tic = time.perf_counter()

def trace(map, dbls):
    ct = 0
    tracker = deque([("start", set(["start"]), False)])
    while tracker:
        p, seen, visited = tracker.popleft()
        if p == "end":
            ct += 1
            continue
        for c in map[p]:
            if c not in seen:
                seen_cp = set(seen)
                if c.islower():
                    seen_cp.add(c)
                tracker.append((c, seen_cp, visited))
            elif c in seen and not visited and c not in ["start", "end"] and dbls:
                tracker.append((c, seen, c))
    return ct


data = open("./day_12/12.in").read().strip().split("\n")
map = defaultdict(list)
for line in data:
    p, c = line.split("-")
    map[p].append(c)
    map[c].append(p)
print(f"Part 1: {trace(map, False)}")
print(f"Part 2: {trace(map, True)}")
toc = time.perf_counter()
print(f"Puzzle 12 python -> {toc - tic:0.4f} seconds")