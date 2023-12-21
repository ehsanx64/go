run:
	go run .

test:
	go build -o flag .
	@echo

	# Invoke the program with no flags
	./flag 
	@echo

	# String flag
	./flag -name Adam
	@echo
	go run . --name=Adam
	@echo
	go run . -name "Adam"
	@echo
	go run . --name "Adam"
	@echo

	# Boolean flags require the '='
	go run . --name "Adam" -age 40 --smart false
	@echo
	go run . --name "Adam" -age 40 --smart="false"
	@echo
	go run . --name "Adam" -age 40 --smart=true
	@echo

	# Using StringVar
	go run . --skill "over thinking"
	@echo
	go run . --skill over-thinking
	@echo
	go clean
