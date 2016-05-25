# Golang - React App Using Beego Framework

## Summary  

A golang-react app. 

## Setup  
  
[beego](http://beego.me/) and [neoism](https://github.com/jmcvetta/neoism) are required go packages.  

Run `npm install && bower install` to install node and bower packages.

Run `gulp` to build the project and start the webserver on `localhost:8080`.

## Todos  

- Create task to watch for changes and re-run build process.
- Add support for modular code in dev folder. 
- Create tasks to update index.tpl when packages are installed
- Create tasks to update index.tpl when new files are created
- Add task to run tests
- Add task for production/deploy 
- Create client side unit tests. 

## Tests  

- `gulp server-tests` will run all tests in `beego/tests`.