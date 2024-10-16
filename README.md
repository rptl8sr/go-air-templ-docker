# Auto reload app and browser (GO + AIR + TEMPL + DOCKER)
Use this template if you need to auto-reload the go-application with web-interface

## How to use with Docker
1. Install Docker
2. Run containers with command ```docker-compose up```
3. Open your app in browser ```http://localhost:8080```
4. Make changes in .templ or .go files
5. See the changes
6. Enjoy it (or not)

## How to use without Docker
1. Install go 1.23, node 18
2. Install go dependencies ```go mod download && go mod verify```
3. Install node dependencies ```npm i``` 
4. Run ```make watch```
5. Make changes in .templ or .go files
6. See the changes
7. Also enjoy it (or not)


