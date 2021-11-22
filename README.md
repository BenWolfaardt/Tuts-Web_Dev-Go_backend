# Create a serverless Telegram bot using Go and Vercel

Source: [Tutorial](https://dev.to/jj/create-a-serverless-telegram-bot-using-go-and-vercel-4fdb)

I wanted to do this tutorial so that I could learn how to 
Now I can use a template and host a basic website for myself! 

> Check out: benwolfaardt.com

# [Going Serveless With Vercel](https://dev.to/sumitkolhe/going-serveless-with-vercel-5b4o)

What exactly is serverless?

Serverless, as the name goes, means running code without a server. Serverless is more of an architecture that defines how the code is to be handled. In a traditional server environment, a piece of code is executed on the server and the requests and responses are transferred between the client and the server. In other words, it can be said as the server is the environment where the execution of the server takes place.

How serverless is different?
Well, Serverless is a misleading word as servers are still needed in this type of architecture but the developers don't have to explicitly worry about managing/setting up the servers in any way. Going Serverless enables the developers to think about the applications on a task level rather than having to worry about it on a server level.

Think of serverless as breaking down your applications in separate smaller modules that can run independently. This concept is similar to microservices, but serverless goes even one step further than microservices. Microservices demand dividing the application in smaller modules depending on the kind of services they carry out say, for example, an authentication module is a microservice for a social media website as it only handles the sign-in / sign-up functionality. Microservices can be thought of as a collection of multiple functions whereas serverless on the other hand demands dividing the application on a task/function level.

Serverless functions depend on the platform they are running on. AWS Lambda, Google Cloud, Microsoft Azure, Vercel these are some great environments to run your serverless functions.

This tutorial made use of the following `modules`
* [GeertJohan/go.rice](https://github.com/GeertJohan/go.rice)
  * It is used to serve the front end on `Go`;
* [gorilla/mux](https://github.com/gorilla/mux)
  * Implements a request router and dispatcher for matching incoming requests to their respective handler.

# Run The Code

Run the below commands to host the website on your `Go` backend on http://localhost:8080/

```go
    go build
    ./02-polygot-frontend_served_in_go 
```

// TODO: edit main repository README to include the below commmand
<!-- git branch --set-upstream-to=origin/03-dev-create_a_serverless_Telegram_bot_on_Vercel 03-dev-create_a_serverless_Telegram_bot_on_Vercel -->
<!-- Branch '03-dev-create_a_serverless_Telegram_bot_on_Vercel' set up to track remote branch '03-dev-create_a_serverless_Telegram_bot_on_Vercel' from 'origin'. -->