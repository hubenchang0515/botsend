.PHONY: all install

all:
	go build

install: botsend
	install -m755 botsend /usr/bin/botsend