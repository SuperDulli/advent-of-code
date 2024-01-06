brackets = open("input", "r").read()
# brackets = "(()(()("

# Part 1 - which floor
floor = 0
for bracket in brackets:
	if bracket == "(":
		floor += 1
	elif bracket == ")":
		floor -= 1
print(floor)

# Part 2 - how many instruction until basement
floor = 0
count = 0
for bracket in brackets:
	if bracket == "(":
		floor += 1
	elif bracket == ")":
		floor -= 1
	count += 1
	if floor == -1:
		break
print(count)
