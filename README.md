# Bundle HTML, CSS, And JavaScript To Be Served In A Golang Application

Source: [Tutorial](https://www.thepolyglotdeveloper.com/2017/03/bundle-html-css-javascript-served-golang-application/)

I wanted to do this tutorial so that I could learn how to host a website on a `Go` backend. Now I can apply this knowledge and use it to serve a basic frontend template on a `Go` backend for a place holder to my first itteration of my website;s temporary landing page! 

> Check out: [benwolfaardt.com](https://BenWolfaardt.com)

--- 

This tutorial made use of the following `modules`
* [GeertJohan/go.rice](https://github.com/GeertJohan/go.rice)
  * It is used to serve the front end on `Go`;
* [gorilla/mux](https://github.com/gorilla/mux)
  * Implements a request router and dispatcher for matching incoming requests to their respective handler.

# Run The Code

Run the below commands to correctly configure your `go.mod` and `go.sum` files and run the frontend of your website on your `Go` backend at http://localhost:8080/

```sh
  ./start.sh
```
