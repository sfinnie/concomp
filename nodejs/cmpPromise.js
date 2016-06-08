var fs = require('fs');
var Promise = require('promise');

var fname1=process.argv[2]
var fname2=process.argv[3]
var size1 = undefined
var size2 = undefined

console.log("\n~~~~~ Start ~~~~~\n")
console.log("1. Creating promises to obtain file sizes... ")
p1=fileSize(fname1)
p2=fileSize(fname2)
console.log("2. Promises created.")

console.log("3. Waiting on promises to complete... ")
p1.then((size) => console.log("    3.1. Size of file " + fname1 + ": " + size)) 
p1.then((size) => size2 = size) 
p2.then((size) => console.log("    3.1. Size of file " + fname2 + ": " + size)) 
console.log("4. Promises Resolved. ")

console.log("5. Comparing file sizes...")
if(size1 > size2) {
	console.log("    The first file is bigger")
} else if (size2 > size1) {
	console.log("    The second file is bigger")
} else {
	console.log("    The files are the same size")
}
console.log("6. Comparison complete.")
console.log("\n~~~~~ End ~~~~~\n")

function fileSize(fname) {
	return new Promise(function (resolve, reject) {
		fs.stat(fname, function doneReading(err, stats) {
			if (err) {
				reject(err)
			} else {
				resolve(stats["size"])
			}
		})
	})
}





