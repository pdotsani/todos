'use strict';

var gulp = require('gulp');
var babel = require('gulp-babel');
var run = require('gulp-run');
var sequence = require('gulp-sequence');

gulp.task("build", function () {
  return gulp.src("dev/js/app.js")
    .pipe(babel())
    .pipe(gulp.dest("static/js"));
});

gulp.task("run-beego", function() {
	run('bee run', {silent: false, verbosity: 3}).exec()
});

gulp.task("default", sequence("build", "run-beego"));