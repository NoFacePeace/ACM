default: reflex
run:
	go run main.go >> log
reflex:
	reflex -r '\.go' make run 
tail:
	tail -f log