'use strict';

// Include Gulp & tools we'll use
var gulp = require('gulp');
var $ = require('gulp-load-plugins')();
var del = require('del');
var exec = require('child_process').execSync;
var merge = require('merge-stream');
var path = require('path');

// Load tasks for web-component-tester
// Adds tasks for `gulp test:local` and `gulp test:remote`
require('web-component-tester').gulp.init(gulp);

// Clean output directory
gulp.task('clean', function() {
	return del(['.tmp', 'server/static', 'server/templates']);
});

// Compile Go down to JS
var buildGopherJS = function (cb) {
	exec('mkdir -p .tmp && cd .tmp && gopherjs build -m code.psg.io/polymer-template/client && cd ..', function (err, stdout, stderr) {
		if(stdout.length != 0) {
			console.log(stdout);
		}
		if(stderr.length != 0) {
			throw stderr.toString()
		}
	});
}

// Build assets from resource directories
var buildGoBinData = function (cb) {
	exec('cd server && go-bindata-assetfs static/... templates/... && cd ..', function (err, stdout, stderr) {
	if(stdout.length != 0) {
			console.log(stdout);
		}
		if(stderr.length != 0) {
			throw stderr.toString()
		}
	});
}

var validate = function() {
	// TODO: Figure out how to make polylint work with our setup
}

// Build for development
gulp.task('build-step:dev', ['clean'], function() {
	// Compile gopherJS
	buildGopherJS()

	// First level html files get prefixed and copied seperately to templates/
	var firstLevelHtml = gulp.src('client/*.html')
		.pipe($.htmlAutoprefixer())
		.pipe(gulp.dest('server/templates'));

	// Prefix and copy over all other html files
	var otherHtml = gulp.src(['client/**/*.html', '!client/*.html', '!client/bower_components/**/*'])
		.pipe($.htmlAutoprefixer())
		.pipe(gulp.dest('server/static'));

	// Prefix and copy over css
	var css = gulp.src(['client/**/*.css', '!client/bower_components/**/*'])
		.pipe($.autoprefixer())
		.pipe(gulp.dest('server/static'));

	// Copy over javascript
	var js = gulp.src(['client/**/*.js', '!client/bower_components/**/*'])
		.pipe(gulp.dest('server/static'));

	// Copy over images and optimize them
	var images = gulp.src(['client/**/*.png', 'client/**/*.jpg', 'client/**/*.gif', 'client/**/*.svg'])
		.pipe(gulp.dest('server/static'));

	// Copy over bower components, don't bother trying to prefix those
	var bowerComponents = gulp.src('client/bower_components/**/*')
		.pipe(gulp.dest('server/static/bower_components'));

	// Copy gopherJS output to dist as well
	var gopherJS = gulp.src('.tmp/**/*')
		.pipe(gulp.dest('server/static'));

	return merge(firstLevelHtml, otherHtml, css, js, images, bowerComponents, gopherJS)
});

gulp.task('validate-step:dev', ['clean', 'build-step:dev'], function() {
	return validate();
})

gulp.task('build:dev', ['clean', 'build-step:dev', 'validate-step:dev'], function() {
	// Embed static and template files
	buildGoBinData()
});

// Build .tmp dir for production
// We use .tmp as a temporary staging directory for all code files
// files will get picked up by either vulcanize or useref if it needs to move to dist
gulp.task('build-step-1:prod', ['clean'], function() {
	// Compile gopherJS
	buildGopherJS()

	// Move all js, css and html files
	var processCSS = gulp.src('client/**/*.css').pipe(gulp.dest('.tmp'));
	var processJS = gulp.src(['client/**/*.js', '!client/bower_components/webcomponentsjs/webcomponents-lite.js']).pipe(gulp.dest('.tmp'));
	var processHTML = gulp.src('client/**/*.html').pipe(gulp.dest('.tmp'));

	// Uglify webcomponents-lite.js
	var uglifyWebComponents = gulp.src('client/bower_components/webcomponentsjs/webcomponents-lite.js').pipe($.uglify()).pipe(gulp.dest('.tmp/bower_components/webcomponentsjs/'));

	return merge(processCSS, processJS, processHTML, uglifyWebComponents);
});

// Build for production
// Note: We use .tmp as an intermediary folder for all html/css/js
gulp.task('build-step-2:prod', ['clean', 'build-step-1:prod'], function() {
	var minifyInlineOpts = {css: false};

	// Optimize images
	var processImages = gulp.src(['client/**/*.png', 'client/**/*.jpg', 'client/**/*.gif', 'client/**/*.svg', '!client/bower_components/**/*'])
		.pipe($.imagemin())
		.pipe(gulp.dest('server/static'));

	// Process elements/elements.html
	var processElements = gulp.src('.tmp/elements/elements.html')
		.pipe($.vulcanize({
		  stripComments: true,
		  inlineCss: true,
		  inlineScripts: true
		}))
		.pipe($.htmlAutoprefixer())
		.pipe($.minifyHtml())
		.pipe($.minifyInline(minifyInlineOpts))
		.pipe($.rename('elements.min.html'))
		.pipe(gulp.dest('server/static'));

	// process templates and their deps
	var processRootHtml = gulp.src('.tmp/*.html')
		.pipe($.replace('elements/elements.html', 'elements.min.html'))
		.pipe($.useref())
		.pipe($.if("*.html", $.htmlAutoprefixer()))
		.pipe($.if("*.html", $.minifyHtml()))
		.pipe($.if("*.html", $.minifyInline()))
		.pipe($.if("*.html", gulp.dest('server/templates')))
		.pipe($.if("!*.html", gulp.dest('server/static')));

	return merge(processImages, processElements, processRootHtml)
});

gulp.task('validate-step:prod', ['clean', 'build-step-1:prod', 'build-step-2:prod'], function() {
	return validate();
})

gulp.task('build:prod', ['clean', 'build-step-1:prod', 'build-step-2:prod', 'validate-step:prod'], function() {
	// Embed static and template files
	buildGoBinData()
});
