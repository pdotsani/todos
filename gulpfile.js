'use strict';

var gulp = require('gulp');
var babel = require('gulp-babel');
var run = require('gulp-run');
var sequence = require('gulp-sequence');
var clean = require('gulp-clean');

gulp.task("clean-static", function() {
	return gulp.src("static/", {read: false})
		.pipe(clean());
});

gulp.task("build-js", function () {
  return gulp.src("dev/js/app.js")
    .pipe(babel())
    .pipe(gulp.dest("static/js"));
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

gulp.task("default", sequence("clean-static", "build-js", "build-css", "build-imgs", "run-beego"));