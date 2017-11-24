targetImagename = myapp
all: package

package: build createImage

build: 
	go build -o myapp main.o
createImage:
	docker build -t "$(targetImagename)"