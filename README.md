## Polymer Template

> A starting point for building web applications with Polymer 1.0 and GopherJS

### Get the code

Visit (https://github.com/PalmStoneGames/polymer-template/) and fetch the Polymer Template.

### Install dependencies

#### Quick-start (for experienced users)

With Node.js installed, run the following one liner from the root of your Polymer Template download:

```sh
npm install -g gulp bower && npm install && bower install
```

#### Prerequisites (for everyone)

Polymer template requires the following dependencies:

- Node.js, used to run JavaScript tools from the command line.
- npm, the node package manager, installed with Node.js and used to install Node.js packages.
- gulp, a Node.js-based build tool.
- bower, a Node.js-based package manager used to install front-end packages (like Polymer).
- go, the programming language
- gopherJS, a compiler for Go to Javascript

**To install dependencies:**

1)  Check your Node.js version.

```sh
node --version
```

The version should be at or above 0.12.x.

2)  If you don't have Node.js installed, or you have a lower version, go to [nodejs.org](https://nodejs.org) and click on the big green Install button.

3)  Install `gulp` and `bower` globally.

```sh
npm install -g gulp bower
```

This lets you run `gulp` and `bower` from the command line.

4)  Install the local `npm` and `bower` dependencies.

```sh
cd polymer-template && npm install && bower install
```

This installs the element sets (Paper, Iron, Platinum) and tools required to build and serve apps.

5)  Download and install the Go programming language.

Please follow the instructions for your system at https://golang.org/doc/install

6)  Grab the GopherJS compiler.

```sh
go get -u github.com/gopherjs/gopherjs
```

### Build

Build the current project, for development and testing.

```sh
gulp build:dev && go build ./...
```

Build and optimize the current project, ready for deployment.

```sh
gulp build:prod && go build ./...
```

## Dependency Management

Polymer uses [Bower](http://bower.io) for package management. This makes it easy to keep your elements up to date and versioned. For tooling, we use npm to manage Node.js-based dependencies.