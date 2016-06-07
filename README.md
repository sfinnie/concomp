#filecomp
Example implementations of concurrent file size comparison in various languages.  This repo is somewhat inspired by [todoMVC](http://todomvc.com/).  It's useful to see how different languages approach a problem that's simple enough to understand, yet complex enough to get a feel for the language features and philosophy.

#The Problem

Is best described by example:

    $ filecmp file1.txt file2.txt file3.txt
    
    file2.txt is the biggest.
    
    $ filecmp file1.txt file3.txt
    
    The files are the same size.

That's it. Simple.  A program that:

* runs from the command line
* accepts two or more parameters, where each is the name of a file
* compares the sizes of each, and 
* reports which is largest.
* If they're all the same size, it should say so.
 
##Where's the interest?
The purpose is to demonstrate the concurrency features in the language.  So there's one other requirement:

* Fetching the size of the supplied files should be completed concurrently.

What that means will differ across languages - that's the point.

To be more credible, the files to compare should be remote resources rather than local - so fetching each takes a non-trivial amount of time.  However, that's not mandatory.  It's not a performance test.  It's about coding style.  As long as the code demonstrates fetching concurrently, it doesn't really matter where the files are located.


