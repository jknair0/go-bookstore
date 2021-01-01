build-image:
	docker build -t jayakrishnan1236/go_bookstore .
run-container:
	docker run -it -d -p 8000:8000 jayakrishnan1236/go_bookstore
push-image:
	docker push jayakrishnan1236/go_bookstore