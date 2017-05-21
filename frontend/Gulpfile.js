var gulp = require('gulp'),
	htmlmin = require('gulp-htmlmin'),
	sass = require('gulp-sass'),
	sourcemaps = require('gulp-sourcemaps'),
	autoprefixer = require('gulp-autoprefixer'),
	concat = require('gulp-concat'),
	uglify = require('gulp-uglify'),
	imagemin = require('gulp-imagemin'),
	plumber = require('gulp-plumber'),
	notify = require('gulp-notify'),
	livereload = require('gulp-livereload'),
	fs = require('node-fs');
	fse = require('fs-extra'),
	json = require('json-file'),
	nib = require('nib'),
	themeName = json.read('./package.json').get('themeName'),
	themeDir = '../' + themeName,
	ico = require('gulp-to-ico'),
	plumberErrorHandler = { errorHandler: notify.onError({
		title: 'Gulp',
		message: 'Error: <%= error.message %>'
		})
	};

gulp.task('html', function() {
  return gulp.src('src/html/*.html')
    .pipe(htmlmin({collapseWhitespace: true}))
    .pipe(gulp.dest('../public/'))
	.pipe(livereload());
});

gulp.task('styles', function() {
	gulp.src('./src/styles/*.scss')
		.pipe(sourcemaps.init())
		.pipe(plumber(plumberErrorHandler))
		.pipe(sass())
		.pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
		.pipe(sourcemaps.write())
		.pipe(autoprefixer())
		.pipe(gulp.dest('../public'))
		.pipe(livereload());
});

gulp.task('scripts', function () {
	gulp.src('./src/scripts/*.js')
		.pipe(plumber(plumberErrorHandler))
		.pipe(gulp.dest('../public'))
		.pipe(livereload());
});

gulp.task('watch', function() {
	livereload.listen();
	gulp.watch('./src/html/*.html', ['html']);
	gulp.watch('./src/styles/*.{scss,sass}', ['styles']);
	gulp.watch('./src/scripts/*.js', ['scripts']);
});

 
gulp.task('default', ['html', 'styles', 'scripts', 'watch']);