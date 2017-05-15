#!/bin/bash
export PATH=$PATH:/usr/local/go/bin
python3 shuffle.py
go build
./hello
