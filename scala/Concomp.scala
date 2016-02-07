import java.io.File

import scala.collection.mutable.ListBuffer
import scala.concurrent.ExecutionContext.Implicits.global
import scala.concurrent.duration._
import scala.concurrent.{Await, Future}
import scala.language.postfixOps

object Concomp {
  class FileInfo(nameC: String, sizeC: Long) {
    var name: String = nameC;
    var size: Long = sizeC;
  }

  class Result(biggestC: FileInfo, evensC: ListBuffer[String]) {
    var biggest: FileInfo = biggestC;
    var evens: ListBuffer[String] = evensC;
  }

  def main(args: Array[String]) = {
    val futures = new ListBuffer[Future[FileInfo]]

    // Start futures to get file sizes
    args foreach (arg => {
      futures += getFileSize(arg)
    })

    // Process futures
    val result = getBiggestAndEvens(futures)

    // Display result
    if (result.evens.length == args.length) {
      println("All files are even")
    } else if (result.evens.length > 1) {
      println("Biggest files are:")
      result.evens.foreach(e => println("  " + e))
    } else {
      println("The biggest is: " + result.biggest.name)
    }
  }

  def getFileSize(fileName: String): Future[FileInfo] = Future {
    val file = new File(fileName)
    val fileInfo = new FileInfo(fileName, file.length())
    fileInfo
  }

  def getBiggestAndEvens(futures: ListBuffer[Future[FileInfo]]): Result = {
    var biggest = new FileInfo("", 0)
    var evens = new ListBuffer[String]
    Future.sequence(futures).foreach(list => {
      list.foreach(fileInfo => {
        if (fileInfo.size > biggest.size) {
          biggest = fileInfo
          evens = new ListBuffer[String]
        }
        if (fileInfo.size == biggest.size) {
          evens += fileInfo.name
        }
      })
    })
    // Wait 60 seconds max for each future before to return the result
    futures.foreach(f => Await.result(f, 60 seconds))
    new Result(biggest, evens)
  }
}