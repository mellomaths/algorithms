alias t := test

test TEST:
	go test ./{{TEST}} -v

tests:
	go test ./... -v
