.PHONY: build test clean

build:
	go build -o TextConvertor -buildmode=c-shared TextConvertor.go

test:
	python3 test.py

clean:
	rm -f TextConvertor TextConvertor.h