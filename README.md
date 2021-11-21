# Writing Web Applications

Source: [Tutorial](https://golang.org/doc/articles/wiki/)

I wanted to do this tutorial so that I could learn about the internal `Go` modules that are used for websites built on a `Go` backend. I will then use this as the fundamental knowledge that I will carry with me through this repository where I will start to understand the building blocks required to build a website. This knowledge will then be used to build the first itteration of my website - a temporary landing page! 

> Check out: [benwolfaardt.com](https://BenWolfaardt.com)

---

This tutorial made use of the following internal `Go`  `modules`
* [errors](https://pkg.go.dev/errors?utm_source=gopls)
  * Implements functions to manipulate errors.
* [html/template](https://pkg.go.dev/html/template?utm_source=gopls)
  * Implements data-driven templates for generating HTML output safe against code injection. It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.
* [io/ioutil](https://pkg.go.dev/io/ioutil?utm_source=gopls)
  * Implements some I/O utility functions.
* [log](https://pkg.go.dev/log?utm_source=gopls)
  * Implements a simple logging package. It defines a type, Logger, with methods for formatting output. It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than creating a Logger manually. That logger writes to standard error and prints the date and time of each logged message. Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one. The Fatal functions call os.Exit(1) after writing the log message. The Panic functions call panic after writing the log message.
* [net/http](https://pkg.go.dev/net/http?utm_source=gopls)
  * Provides HTTP client and server implementations.
* [regexp](https://pkg.go.dev/regexp?utm_source=gopls)
  * Implements regular expression search.

# Run The Code

Run the below commands to correctly configure your `go.mod` and `go.sum` files and run the frontend of your website on your `Go` backend at: http://localhost:8080/view/ANewPage

```sh
  ./start.sh
```

# A few other tasks / TODO

Here are some simple tasks you might want to tackle on your own:

* Store templates in tmpl/ and page data in data/.
* Add a handler to make the web root redirect to /view/FrontPage.
* Spruce up the page templates by making them valid HTML and adding some CSS rules.
* Implement inter-page linking by converting instances of [PageName] to  
  <a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)