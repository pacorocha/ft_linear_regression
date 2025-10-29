run:
	go run trainer/trainer.go
	go build  -o bin/predictor predictor.go
	./bin/predictor

fclean:
	rm -rf bin/
	rm -f *.txt

.PHONY: run fclean