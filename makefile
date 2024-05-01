
tmplgen:
	templ generate 

tmplformate: 
	go run ./server/cmd/ui fmt .


run:
	go run ./server/app/cmd/pdfsync/main.go