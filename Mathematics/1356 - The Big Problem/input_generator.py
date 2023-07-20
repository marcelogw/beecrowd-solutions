import random

with open('main/input.txt', 'w') as file:
    for _ in range(10000):
        number = random.randint(10**7, 10**8)
        file.write(str(number) + '\n')
