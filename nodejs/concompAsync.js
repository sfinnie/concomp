var fs = require('fs');

var fname1=process.argv[2]
var fname2=process.argv[3]

fs.stat(fname1, function doneReading(err, stats) {
	size1=stats["size"]
	fs.stat(fname2, function doneReading(err, stats) {
		size2=stats["size"]
		if(size1 > size2) {
			console.log(fname1 + " is bigger")
		} else if (size2 > size1) {
			console.log(fname2 + " is bigger")
		} else {
			console.log("The files are the same size")
		}
	})
})


