run: build
	docker run -it bomberman
	
build:
	docker build -t bomberman .
