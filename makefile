
tmplgen:
	templ generate 

tmplformate: 
	go run ./server/cmd/ui fmt .


run:
	go run ./cmd/sync/*.go

doc:
	swag init -d ./cmd/sync/ -g ./main.go -o ./cmd/sync/docs

docpritty:
	swag fmt