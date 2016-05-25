# Golang - React App Using Beego Framework

## Summary  

Preliminary skeleton code for a golang-react app. 

1. `gulp 'default'` now builds the js, img, and css files, from `dev/`then migrates them to `static/`.  

2. `bee run` then starts the beego build, then starts localhost.  

NPM and Bower packages are linked by setting static paths via [server side router](https://github.com/pdotsani/todos/blob/master/routers/router.go)

## Setup  

Run `npm install && bower install`. This should install all necessary dependencies to run the app as is.  

To start the app. Run `gulp` at the root folder where gulpfile.js lives.  

## Todos  

- Create task to watch for changes and re-run build process.
- Add support for sourcemaps during js build: [gulp-sourcemaps](https://github.com/floridoo/gulp-sourcemaps)  
- Create tasks to update index.tpl when packages are installed
- Create tasks to update index.tpl when new files are created
- Add task to run tests
- Add task for production/deploy