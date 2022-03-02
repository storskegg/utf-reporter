BIN=utf-reporter

all: build upx

build:
	go build -ldflags "-w -s" -a .

upx:
	upx -9 -k $(BIN)

clean:
	rm $(BIN)*
