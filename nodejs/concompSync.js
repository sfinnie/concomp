var fs = require('fs');

var fname1=process.argv[2]
var fname2=process.argv[3]

var stats1 = fs.statSync(fname1)
size1=stats1["size"]

var stats2 = fs.statSync(fname2)
size2=stats2["size"]

if(size1 > size2) {
	console.log(fname1 + " is bigger")
} else if (size2 > size1) {
	console.log(fname2 + " is bigger")
} else {
	console.log("The files are the same size")
}


