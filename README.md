# go-chuck

> Chuck Norris' random facts app in Go

## How to run app locally ?

Before you get started make sure you have an up-to-date version of Go and Node, on macOS with [Homebrew][brew] installed you can run :

```bash
$ brew install node
$ brew install go
```

After that you can run:

```bash
$ go get -d github.com/nurlansu/go-chuck
$ cd $GOPATH/src/github.com/nurlansu/go-chuck
```

If you haven't [Wellington][wt], you should install it on your computer by running :

```bash
$ go get -u github.com/wellington/wellington/wt
```

Now we can install all necessary `node` modules :

```bash
$ npm install
```

To build front-end stuff, you should run:

```bash
$ npm build
```

And at the end you can run the app using this command:

```bash
$ go run *.go
```

And that's all. Normally you could see the app running on http://localhost:8080

## Dependencies

This app uses [httprouter][router] for routing.

## Demo

The live version can be found on https://go-chuck.herokuapp.com/

## License

<p align="center">
  <a href="./LICENSE"><img src="https://i.nurlan.co/logo.svg" width="100%" height="128"></a>
  <a href="./LICENSE"><strong>MIT</strong></a>
</p>



[brew]: https://brew.sh
[json]: https://github.com/mailru/easyjson
[router]: https://github.com/julienschmidt/httprouter
[wt]: https://github.com/wellington/wellington
