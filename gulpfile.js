var gulp = require('gulp');
var uglify = require('gulp-uglify');
var sourcemaps = require('gulp-sourcemaps');
var shell = require('gulp-shell');
var flatten = require('gulp-flatten');

var paths = {
  minify_scripts: ['bower_components/**/dist/*.min.js', 'bower_components/**/dist/js/*.min.js'],
  minify_styles: ['bower_components/**/dist/css/*.min.css'],
  fonts: ['bower_components/**/dist/fonts/*']
}

gulp.task('minify-scripts', function() {
  return gulp.src(paths.minify_scripts)
    .pipe(flatten())
    .pipe(gulp.dest('public/js'));
});

gulp.task('minify-styles', function() {
  return gulp.src(paths.minify_styles)
    .pipe(flatten())
    .pipe(gulp.dest('public/css'));
});

gulp.task('fonts', function() {
  return gulp.src(paths.fonts)
    .pipe(flatten())
    .pipe(gulp.dest('public/fonts'));
});

// TODO gulp.watch 적용


gulp.task('run', shell.task('go run main.go'));

// The default task (called when you run `gulp` from cli)
gulp.task('default', ['minify-scripts', 'minify-styles', 'fonts', 'run']);
