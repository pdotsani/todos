'use strict';

var gulp = require('gulp');
var run = require('gulp-run');
var sequence = require('gulp-sequence');
var clean = require('gulp-clean');
var browserify = require('browserify');
var source = require('vinyl-source-stream');
var gutil = require('gulp-util');
var babelify = require('babelify');

var bundleApp = browserify({
	entries: "dev/js/app.js",
	debug: true
})

gulp.task("clean-static", function() {
	return gulp.src("static/", {read: false})
		.pipe(clean());
});

gulp.task("build-js", function() {
	return bundleApp
		.transform("babelify", {presets: ["es2015", "react"]})
		.bundle()
		.on('error', gutil.log)
		.pipe(source('app.js'))
		.pipe(gulp.dest("./static/js"));
});

gulp.task("build-css", function() {
	return gulp.src("dev/css/*")
		.pipe(gulp.dest("static/css"));
});

gulp.task("build-imgs", function() {
	return gulp.src("dev/img/*")
		.pipe(gulp.dest("static/img"))
});

gulp.task("server-tests", function() {
	run('cd tests && go test -v && cd ..',
		{silent: false, verbosity: 3}).exec()
});

gulp.task("run-beego", function() {
	run('bee run', {silent: false, verbosity: 3}).exec()
});

gulp.task("build-static", sequence("build-js", "build-css", "build-imgs"));

gulp.task("default", sequence("clean-static", "build-static", "run-beego"));