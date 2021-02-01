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

### There are no stupid questions.

If you're struggling to understand something, either be it in this project or not, ask - ensure you've done some research
first. As a team and company we want to hear these questions as it makes us better as individuals and as a
whole.

### What is the meaning of life?
42.

## Ok get on with it

Fine. This project is going to define the specifications of a project that you'll write, and we'll review, simple right?
Don't fret, it's mainly for us, we want to know what we're missing for new starters, what you might need, and what you might
already know, **BONUS POINTS** (there are no points).

In this project we'll cover **Golang**, **Kubernetes**, and **Docker**. Everything is to be done on a local machine no
cloud account required. We'll be working in steps; first we'll set the spec for a purely local project, then we'll see
the differences between Kubernetes and Lambda.

### Install

For installation we'll assume you're running on a MacOS machine as all devs use one (sorry PC master race). For an IDE
you  can choose anything to your liking, but GymShark offers a licence for Intelij Ultimate - it has some helpers that
make life easy.

* [brew](https://brew.sh/)
* Go - `brew install go`
* [Kubernetes/Docker](https://www.docker.com/products/docker-desktop)
* kubectl - `brew install kubernetes-cli`

Once docker desktop has been  installed, kubernetes won't be enabled. To enable it, open docker desktop, open preferences,
kubernetes and click enable kubernetes.