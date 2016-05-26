# Golang - React App Using Beego Framework

## Summary  

A golang-react app. 

## Setup  
  
Place repo in the `/src` folder of your GOPATH.  
  
Install [beego w/ cli tools](http://beego.me/) and [neoism](https://github.com/jmcvetta/neoism)  

Run `npm install && bower install`.  

Run `gulp` to run the build sequence and start the webserver on `localhost:8080`.  

## Todos  

- Create task to watch for changes and re-run build process.
- ~~Add support for modular code in dev folder.~~  
- Create tasks to update index.tpl when packages are installed
- ~~Create tasks to update index.tpl when new files are created~~ *not needed with browserify*  
- ~~Add task to run tests~~
- Add task for production/deploy 
- Create client side unit tests. 

## Tests  

- `gulp server-tests` will run all tests in `beego/tests`.