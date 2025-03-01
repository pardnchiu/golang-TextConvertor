.PHONY: build test clean

build:
	go build -o dist/TextConvertor -buildmode=c-shared src/TextConvertor.go

test:
	cd dist && python3 test.py

clean:
	rm -f TextConvertor TextConvertor.h