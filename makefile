targetImagename = myapp
all: package

package: build createImage

build: 
	go build -o myapp main.go
createImage:
<<<<<<< HEAD
	docker build -t "$(targetImagename)" .
=======
	docker build -t "$(targetImagename)" .
>>>>>>> 91e6275e63eff5ced9f64411bfc6d335bb81d764
