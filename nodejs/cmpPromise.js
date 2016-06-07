var fs = require('fs');
var Promise = require('promise');

var fname1=process.argv[2]
var fname2=process.argv[3]

function fileSize(fname) {
	return new Promise(function (fulfill, reject) {
		fs.stat(fname2, function doneReading(err, stats) {
			if (err) {
				reject(err)
			} else {
				fulfill(stats["size"])
			}
		})
	})
}

p1=fileSize(fname1)
p2=fileSize(fname2)

size1=p1.done()
size2=p2.done()

console.log(size1)

if(size1 > size2) {
	console.log("The first file is bigger")
} else if (size2 > size1) {
	console.log("The second file is bigger")
} else {
	console.log("The files are the same size")
}

