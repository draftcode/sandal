parser.go: parser.go.y
	go tool yacc -o parser.go -v parser.output parser.go.y

clean:
	rm -f parser.go parser.output
