include .env

init:
	go get -u github.com/fatih/color
	go get -u github.com/wcharczuk/go-chart
	go build src/train/train.go
	go build src/predict/predict.go

clean: 
	rm $(TRAINNAME)
	rm $(RESNAME)

all: init