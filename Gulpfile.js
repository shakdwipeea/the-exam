/**
 * Created by akash on 22/7/15.
 */

var gulp = require('gulp');
var child = require('child_process');
var util = require('gulp-util');
var notifier = require('node-notifier');
var reload = require('gulp-livereload');
var sync = require('gulp-sync')(gulp).sync;
var server = null;
/*
 * Build application server.
 */
gulp.task('server:build', function() {
    child.execSync('export GOPATH=/home/akash/Kode/Radeon/server/');
    child.execSync('export GOBIN=$GOPATH/bin');
    var build = child.spawnSync('go', ['install', 'server/src/github.com/shakdwipeea/main/server.go']);
    if (build.stderr.length) {
        var lines = build.stderr.toString()
            .split('\n').filter(function(line) {
                return line.length
            });
        for (var l in lines)
            util.log(util.colors.red(
                'Error (go install): ' + lines[l]
            ));
        notifier.notify({
            title: 'Error (go install)',
            message: lines
        });
    }
    return build;
});

/*
 * Restart application server.
 */
gulp.task('server:spawn', function() {
    if (server)
        server.kill();

    /* Spawn application server */
    server = child.spawn('./server/bin/server');

    /* Trigger reload upon server start */
    server.stdout.once('data', function() {
        reload.reload('/');
    });

    /* Pretty print server log output */
    server.stdout.on('data', function(data) {
        var lines = data.toString().split('\n')
        for (var l in lines)
            if (lines[l].length)
                util.log(lines[l]);

    });

    /* Print errors to stdout */
    server.stderr.on('data', function(data) {
        process.stdout.write(data.toString());
    });

    notifier.notify({
        title: 'server restarted',
        message: 'Hah mahanta'
    })
});

/*
 * Watch source for changes and restart application server.
 */
gulp.task('server:watch', function() {

    /* Rebuild and restart application server */
    gulp.watch([
        '*/**/*.go',
    ], sync([
        'server:build',
        'server:spawn'
    ], 'server'));
});

/* ----------------------------------------------------------------------------
 * Interface
 * ------------------------------------------------------------------------- */

/*
 * Build assets and application server.
 */
gulp.task('build', [
    'server:build'
]);

/*
 * Start asset and server watchdogs and initialize livereload.
 */
gulp.task('watch', [
    'server:build'
], function() {
    reload.listen();
    return gulp.start([
        'server:watch',
        'server:spawn'
    ]);
});

/*
 * Build assets by default.
 */
gulp.task('default', ['build', 'watch']);