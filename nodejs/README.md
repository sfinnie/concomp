#Node implementation of filecomp

There are actually 3 implementations here:

* cmpSync: using the synchronous file system calls.  Not strictly compliant with the spec, since files sizes are gathered sequentially.  However useful for reference.
* cmpAsync: implementation using async callbacks
* cmpPromise: implementation using ES6 promises

##To run:

1. Install [node](https://nodejs.org/en/) if you haven't already.
2. clone this repo
3. Run from command line as e.g. `node cmpSync.js file1 file2 file3`.

Substitute `cmpAsync` / `cmpPromise.js` for `cmpSync` in above as appropriate.


