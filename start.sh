#! /bin/sh

if [ `ls -1 go.mod 2> go.sum` ]; 
then
    echo "The go.mod and go.sum files are already present, no need to run: go mod init"
    go mod tidy
else
    echo "The go.mod and go.sum files are going to be generated by the 'go mod init' command"
    go mod init github.com/BenWolfaardt/Tuts-Web_Dev-Go_backend/tree/03-Dev.to-Create_a_serverlesss_Telegram_bot_on_Vercel
    go mod tidy
fi

# cd api
# go build
echo "The Go Wiki is ready and can be accessed at http://localhost:8080/view/ANewPage"
./01-Go-Writing_web_apps