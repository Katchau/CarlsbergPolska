import random, argparse

def shuffle(year, ignore_null):
    with open('dataSet/' + year + 'year.arff', 'r') as source:
        bankrupt = []
        non_bankrupt = []

        for line in source:
            if not ignore_null or (ignore_null and line.find('?')) == -1:
                bankrupt.append(line) if line[-2:-1] == '1' else non_bankrupt.append(line)

        random.shuffle(bankrupt)
        random.shuffle(non_bankrupt)

        min_length = min(len(bankrupt), len(non_bankrupt))

        train = []
        test = []

        for i in range(min_length):
            if i % 3 == 0:
                test.append(bankrupt[i])
                test.append(non_bankrupt[i])
            else:
                train.append(bankrupt[i])
                train.append(non_bankrupt[i])

        random.shuffle(train)
        random.shuffle(test)

    with open('dataSet/' + year + 'yearV2.arff', 'w') as target:
        target.write(str(len(train)) + '\n')
        for target_train in train:
            target.write(target_train)
        for target_test in test:
            target.write(target_test)

parser = argparse.ArgumentParser()
parser.add_argument('--ignore', help='ignore null values', action='store_true')
args = parser.parse_args()
for i in range(1, 6):
    shuffle(str(i), args.ignore)
