#Node implementation of ConComp

There are actually 3 implementations here:

* `concompSync`: using the synchronous file system calls.  Not strictly compliant with the spec, since files sizes are gathered sequentially.  However useful for reference.
* `concompAsync`: implementation using async callbacks.  Note this isn't actually concurrent either.  Its purpose is to illustrate the "legacy" javascript approach.
* `concompPromise`: implementation using ES6 promises.  A concurrent implementation which, as of writing (2016-06-21) represents the idiomatic approach.  Note this will likely be superseded by Async/Await in ES7.

##To run:

1. Install [node](https://nodejs.org/en/) if you haven't already.
2. clone this repo
3. Run from command line as e.g. `node concompSync.js file1 file2 file3`.

Substitute `cmpAsync` / `concompPromise.js` for `concompSync` in above as appropriate.


