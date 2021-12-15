# Deploying a Simple Golang Web App on Heroku

Source: [Tutorial](https://dzone.com/articles/deploying-a-simple-golang-webapp-on-heroku)

I wanted to do this tutorial so that I could learn how to host a website on a `Go` backend. Now I can apply this knowledge and use it to serve a basic frontend template on a `Go` backend for a place holder to my first itteration of my website's temporary landing page! In addition to this, I needed to learn how one goes about hosting a `Go` binary on [Heroku](https://dashboard.heroku.com/apps). Whilst performing this I was also exposed to how much easier it is deploying a website / project as a Docker container. 

> Check out: [benwolfaardt.com](https://BenWolfaardt.com)

--- 

This tutorial made use of the following `modules`:

* [gin-gonic/gin](https://github.com/gin-gonic/gin)
  * Gin is a web framework written in `Go`. It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter).
* [/gin-contrib/static](https://github.com/gin-contrib/static)
  * Static middleware - it allows you to have static front end content "hosted" by your backend.

## Pre-requisites

As well as having had the below pre-requisites:

* A recent version of [Go](https://go.dev/) (I'm using 1.17.2)
* [Docker](https://www.docker.com/)
* A [Heroku](https://heroku.com/) account (the free account works fine for this example.)
* The [Heroku command-line client](https://devcenter.heroku.com/articles/heroku-cli)
* [git](https://git-scm.com/)

## Run The Code Locally

Run the below commands to correctly configure your `go.mod` and `go.sum` files and to run the frontend, **locally**, of your website on your `Go` backend at http://localhost:8080/ (default port).

```sh
  ./start.sh
```

## Run The Code on Heroku

### [Creating a Heroku remote](https://devcenter.heroku.com/articles/git#creating-a-heroku-remote)

#### `heroku create`

The `heroku create` CLI command creates a new empty application on Heroku, along with an associated empty Git repository. If you run this command from your app’s root directory, the empty Heroku Git repository is automatically set as a remote for your local repository.

```
Creating app... done, ⬢ murmuring-hamlet-57084
https://murmuring-hamlet-57084.herokuapp.com/ | https://git.heroku.com/murmuring-hamlet-57084.git
```

#### `heroku stack`

Tell Heroku we want to build this project using a Dockerfile, rather than a buildpack: `heroku stack:set container`

To do this, we also need to create a `heroku.yml` file like this:

```yml
build: 
  docker:
    web: Dockerfile

run:
  web: ./helloworld
```

#### `git remote -v`

You can use the `git remote` command to confirm that a remote named heroku has been set for your app:

```
heroku  https://git.heroku.com/murmuring-hamlet-57084.git (fetch)
heroku  https://git.heroku.com/murmuring-hamlet-57084.git (push)
origin  git@github.com:BenWolfaardt/Tuts-Web_Dev-Go_backend.git (fetch)
origin  git@github.com:BenWolfaardt/Tuts-Web_Dev-Go_backend.git (push)
```

##### For an existing Heroku app (if it already exists)

If you have already created your Heroku app, you can easily add a remote to your local repository with the heroku git:remote command. All you need is your Heroku app’s name:

`heroku git:remote -a murmuring-hamlet-57084`

```
set git remote heroku to https://git.heroku.com/murmuring-hamlet-57084.git
```

##### `git remote rename` (optional)

By default, the Heroku CLI names all of the Heroku remotes it creates for your app heroku. You can rename your remotes with the `git remote rename` command:

`git remote rename heroku heroku-staging`

Renaming your Heroku remote can be handy if you have multiple Heroku apps that use the same codebase (for example, the staging and production versions of an app). In this case, each Heroku app has its own remote in your local repository.

### Deploying code

To deploy your app to Heroku, you typically use the `git push` command to push the code from your local repository’s master or main branch to your heroku remote, like so:

`git push heroku main`

However, in our case we aren't pushing from our main branch, therefore: 

#### [Pushing a none "main" branch to Heroku](https://stackoverflow.com/questions/14593538/make-heroku-run-non-master-git-branch)

You can push an alternative branch to Heroku using Git.

`git push heroku 04-DZone-Deploying_a_simple_Go_web_app_on_Heroku:main`

This pushes your local test branch to the remote's master branch (on Heroku).

Worth noting also, when you're ready to go back to master you need to do

`git push -f heroku main:main`

# Interesting Information

Below are some interesting snippets of information that I learnt whilst completing this tutorial. 

## Dockerize

Our Go web application has been written in `Go`, now let's package it up as a Docker image. We could create it as a Heroku buildpack, but one of the nice features of `Go` is that you can distribute your software as a single binary file, and using a Docker-based Heroku deployment. Docker makes deployment easier by addressing the below issues: 

* Different Environments: Docker allows you to come up with your “production” environment on your development machine (your laptop, for example) and then deploy this exact configuration to your production server. There are no differences between these two environments!
* Configuration Management: The Docker configuration files (Dockerfiles and docker-compose.yml) are included in your git repository and they allow you to exactly define the configuration of your overall web application. You can specify the exact version of each image to use, such as Nginx 1.11.3.
* Collaboration: Since the Docker configuration files are stored in your git repository on GitLab (or GitHub or BitBucket), it is easy for a colleague/friend to clone the repository and set up the same application environment on their machine using Docker.

An additional advantage of Docker is having to think about the overall architecture of your web application early on in the development process. Docker forces you to understand how all the pieces in your web application need to work together in a production environment earlier on in the development process

## Serving Static files

The static files are not compiled into the binary, so if you put the `helloworld` file in a different directory, it will not find the `views` directory to serve our HTML and CSS content.

## Remember to rebuild Docker when editing the static files

> The below was stipulated in the guide, but when testing this it didn't seem to be the case!

The dockerized webserver should behave just like the `go run hello.go` and `./helloworld` versions except that it has its own copies of the static files. So, if you change any of the files in `views/` you won't see the changes until you rebuild the Docker image and restart the container.

## `heroku.yml` Overview

##### Source: [Building Docker Images with heroku.yml](https://devcenter.heroku.com/articles/build-docker-images-heroku-yml)

A heroku.yml manifest has 4 top-level sections:

* setup - Specifies the add-ons and config vars to create during app provisioning
* build - Specifies the Dockerfile to build
* release - Specifies the release phase tasks to execute
* run - Specifies process types and the commands to run for each
Details for each of these sections are described below.

Here’s an example that illustrates using a `heroku.yml` manifest to build Docker images:

```yml
setup:
  addons:
    - plan: heroku-postgresql
      as: DATABASE
  config:
    S3_BUCKET: my-example-bucket
build:
  docker:
    web: Dockerfile
    worker: worker/Dockerfile
  config:
    RAILS_ENV: development
    FOO: bar
release:
  command:
    - ./deployment-tasks.sh
  image: worker
run:
  web: bundle exec puma -C config/puma.rb
  worker: python myworker.py
  asset-syncer:
    command:
      - python asset-syncer.py
    image: worker
```

## Miscellaneous

While Docker solves a lot of problems with a traditional deployment process, there are still a number of technologies that you need to understand to do a deployment:

* How to secure a web server
* What Nginx is and how to configure it
* Why you need a WSGI server, such as Gunicorn, in your web application
* How to configure and work with Postgres

## Useful links:

* [Deploying with Git](https://devcenter.heroku.com/articles/git)

## Sources:

Most links are provided where they have been used. Some information is, however, scattered throughout this document, these include, but are not limited to, the following: 

* [Why I Switched from a Traditional Deployment to Using Docker](https://www.patricksoftwareblog.com/why-i-switched-from-a-traditional-deployment-to-using-docker/).
