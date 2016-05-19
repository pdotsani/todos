# Golang - React App Using Beego Framework

## Summary  

Preliminary skeleton code for a golang-react app. Currently gulp only runs two tasks in default mode:  

1. Compiles `static_dev/js/app.js` using babel (react and es2015 packages).
2. Migrates compiled version to `static/js` folder.
3. Starts the beego localserver by running `bee go`.  

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