#!/bin/bash
export PATH=$PATH:/usr/local/go/bin
python3 shuffle.py --ignore
go build
./hello ignore
