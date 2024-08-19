
tmplgen:
	templ generate 

tmplformate: 
	go run ./server/cmd/ui fmt .


run:
	go run ./cmd/sync/*.go