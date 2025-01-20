module myapp

go 1.23.3

replace github.com/ali-ahadi1105/Quokka => ../Quokka

require github.com/ali-ahadi1105/Quokka v0.0.0-20250119052130-1a93aa5f0414

require (
	github.com/go-chi/chi/v5 v5.2.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
