import random
with open('dataSet/1year.arff','r') as source:
    data = [ (random.random(), line) for line in source ]
data.sort()
with open('dataSet/1yearV2.arff','w') as target:
    for _, line in data:
        target.write( line )

with open('dataSet/2year.arff','r') as source:
    data = [ (random.random(), line) for line in source ]
data.sort()
with open('dataSet/2yearV2.arff','w') as target:
    for _, line in data:
        target.write( line )

with open('dataSet/3year.arff','r') as source:
    data = [ (random.random(), line) for line in source ]
data.sort()
with open('dataSet/3yearV2.arff','w') as target:
    for _, line in data:
        target.write( line )

with open('dataSet/4year.arff','r') as source:
    data = [ (random.random(), line) for line in source ]
data.sort()
with open('dataSet/4yearV2.arff','w') as target:
    for _, line in data:
        target.write( line )

with open('dataSet/5year.arff','r') as source:
    data = [ (random.random(), line) for line in source ]
data.sort()
with open('dataSet/5yearV2.arff','w') as target:
    for _, line in data:
        target.write( line )