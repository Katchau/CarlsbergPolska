import random

def shuffle(year):
    with open('dataSet/' + year + 'year.arff', 'r') as source:
        bankrupt = []
        non_bankrupt = []

        for line in source:
            if line.find('?') == -1:
                bankrupt.append(line) if line[-2:-1] == '1' else non_bankrupt.append(line)

        random.shuffle(bankrupt)
        random.shuffle(non_bankrupt)

        min_length = min(len(bankrupt), len(non_bankrupt))

        data = bankrupt[:min_length] + non_bankrupt[:min_length]

        random.shuffle(data)

    with open('dataSet/' + year + 'yearV2.arff', 'w') as target:
        for target_line in data:
            target.write(target_line)

for i in range(1, 6):
    shuffle(str(i))
