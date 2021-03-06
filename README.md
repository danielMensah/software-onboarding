# Software Onboarding

This project roughly outlines what the software team uses often and how we as a team implement functionality and features.

## Why
 
### Why does this project exist? 

Simple, if you're new and want to have a better understanding of how we work on the
GymShark Software team, this should outline most of what we do.

## What

### What does it involve?

This project isn't going to show you how we work down to a T, more like how we like to work and with what technologies.
There is no written rule on what is best and what should  be, if you've found something better tell us, we're not
dicks.

### What is the meaning of life?
42.

## Stupid questions

### There are no stupid questions.

If you're struggling to understand something, either be it in this project or not, ask - ensure you've done some research
first. As a team and company we want to hear these questions as it makes us better as individuals and as a
whole.

## Ok get on with it

Fine. This project is going to define the specifications of a project that you'll write, and we'll review, simple right?
Don't fret, it's mainly for us, we want to know what we're missing for new starters, what you might need, and what you might
already know, **BONUS POINTS** (there are no points).

In this project we'll cover **Golang**, **Docker**, and **Docker-Compose**. Everything is to be done on a local machine
no cloud account required.

### Install

For installation, everything can be controlled and installed via docker/docker-compose. For an IDE
you  can choose anything to your liking, but GymShark offers a licence for IntelliJ Ultimate - it has some helpers that
make life easy.

* [Docker](https://www.docker.com/products/docker-desktop)


### The Spec
The project will make `GET` requests to HackerNew's API and store the objects into a DB (dealers choice), and then show
the objects via a web ui or terminal output (dealers choice).

Retrieve, store, and present any and all information you see fit, there isn't a set way we want to store or see the
objects from HackerNews.

The project should be composed of three services ([Microservice Architecture
pattern](https://microservices.io/patterns/microservices.html)), the services are as follows: Consumer, API, and Presenter.

All services should be written inside the same project but should all be able to run separately. Ensure that all code is
written with testing in mind - interfaces are your best friend. Note that this project doesn't need to be tested to
perfection.

#### Consumer
The Consumer will make requests to [HackerNews' API](https://github.com/HackerNews/API) and store the
objects in the DB. It may mutate the data to a format of your choice, it doesn't have to be the structure from
HackerNews.

There are two endpoints required to get data from HackerNews: 
[TopStories](https://github.com/HackerNews/API#new-top-and-best-stories) and 
[Items](https://github.com/HackerNews/API#items). Some items can be stale, meaning they have been deleted or are dead
(see item docs), we do not want to store these. If an item is deleted and/or dead, disregard it.

The Consumer should be run like a seeder for your own DB data. It doesn't need to run constantly, but it could be run
every X period to stay up-to-date.

The Consumer can be made to run concurrently by sending the TopStory IDs into a channel for a worker to pick up and get
the item. Note it doesn't need to be concurrent, but you will see quicker processing times when running concurrently.

#### API
The API will communicate with the DB that the Consumer seeds for you. 

There should be three endpoints:
* /All - should return all stories regardless of type
* /Stories - should only return type stories
* /Jobs - should only return type jobs

The API should cache responses from the DB - the cache should be invalidated every 5 minutes - this is to speed up
subsequent requests. You can choose any caching mechanism you want, redis is an easy service to use, but it could also
be as simple as an inmemory hashmap

#### Presenter
The Presenter will communicate with the API in order to display top stories that are stored in the DB. 

The data that is returned should be displayed either in a web UI or just simply printing out to the terminal.

### Project structure
There is no exact structure we all work with at GS, it's project specific, with that said, we do have some folders we
like to use for ease of navigation and context. Folders such as pkg, internal, and cmd are the common ones.

There is an [example](example) directory that shows how we might set this up but there are a few things missing
(purposely done), this is so we can see how you design. The base of the project is there and can be used already but has
no functional components in.

Note that `go.mod` should not be copied from here it should be recreated in your own project, this is because of how go
modules work.

The project should utilise Docker and Docker-Compose. Using Docker for your own images that you create (**Consumer**,
**API**, **Presenter**) and Docker-Compose for any other services you might want to use (e.g. Redis, MySQL, etc).