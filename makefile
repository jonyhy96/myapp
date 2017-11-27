targetImagename = myapp
all: package

package: build createImage

build: 
	go build -o myapp main.go
createImage:
	docker build -t "$(targetImagename)" .
