var fs = require('fs');

var fname1=process.argv[2]
var fname2=process.argv[3]

fs.stat(fname1, function doneReading(err, stats) {
	stats1=stats["size"]
	fs.stat(fname2, function doneReading(err, stats) {
		stats2=stats["size"]
		if(size1 > size2) {
			console.log("The first file is bigger")
		} else if (size2 > size1) {
			console.log("The second file is bigger")
		} else {
			console.log("The files are the same size")
		}
	})
})


