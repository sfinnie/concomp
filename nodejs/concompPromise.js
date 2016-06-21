'use strict';

const co = require('co');
const fs = require('fs-promise');

var fname1=process.argv[2]
var fname2=process.argv[3]

co(function* () {
    let res = yield [fs.stat(fname1), fs.stat(fname2)];
	let size1 = res[0]["size"]
	let size2 = res[1]["size"]
	if(size1 > size2) {
		console.log(fname1 + " is bigger")
	} else if (size2 > size1) {
		console.log(fname2 + " is bigger")
	} else {
		console.log("The files are the same size")
	}
})
