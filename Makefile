include .env

init:
	go get -u github.com/fatih/color
	go get -u github.com/wcharczuk/go-chart
	go get -u github.com/mattn/go-pairplot
	go get -u gonum.org/v1/plot
	go get -u gonum.org/v1/plot/vg
	go get -u github.com/nfnt/resize

	go build src/describe.go
	go build src/histogram.go
	go build src/scatter_plot.go
	go build src/pair_plot.go

clean: 
	rm $(P1)
	rm $(P2)
	rm $(P3)
	rm $(P4)

all: init