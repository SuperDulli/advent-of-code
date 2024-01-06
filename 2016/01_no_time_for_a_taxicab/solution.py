instructions = open("input", "r").read()[:-1].split(", ")

def cab_distance(destination):
	return abs(-destination[0]) + abs(-destination[1])

duplicates = []

north = 1
east = 0
destination = [0, 0]
visited = [[0,0]]
for instruction in instructions:
	tmp = north
	if instruction[0] == "R":
		north = east * -1
		east = tmp * 1
	else:
		north = east * 1
		east = tmp * -1
	destination[0] += east * int(instruction[1:])
	destination[1] += north * int(instruction[1:])
	a, b = visited[-1][0] - destination[0], visited[-1][1] - destination[1]
	if a == 0:
		for i in range(visited[-1][1], destination[1], 1 if b < 0 else -1):
			new_dest = [destination[0], i + 1 if b < 0 else i-1]
			if new_dest in visited:
				duplicates.append(new_dest)
			visited.append(new_dest)
	if b == 0:
		for i in range(visited[-1][0], destination[0], 1 if a < 0 else -1):
			new_dest = [i + 1 if a < 0 else i-1, destination[1]]
			if new_dest in visited:
				duplicates.append(new_dest)
			visited.append(new_dest)
print(cab_distance(destination)) # Answer Part 1

print(cab_distance(duplicates[0])) # Answer Part 2
