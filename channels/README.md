# Status Checker application

## Go Routines and Channels

This application will make an http request and check the status of a group of websites.

                               -> http://google.com
                               -> http://facebook.com
Status Checker -> http request -> http://stackoverflow.com
                               -> http://golang.org
                               -> http://amazon.com

The program loops through each link and checks if they are operational and throws errors if not. 

using the `go` keyword it will iteratively run a go routine will only run the checkLink() method that was retrieved.

---

Application architecture will automatically ping the websites after the routine is run for each child routine

Program starts                                                                  Time
|--------------------------------------------------------------------------------->
| ------------------------ ------------------------ ------------------------ 
| |  google.com routine | |  google.com routine   | | google.com routine   |---->
| ------------------------ ------------------------ ------------------------
|
| ------------------------------------------ ------------------------------ 
| |         stackoverflow.com routine      | |  stackoverflow.com routine |---->
| ------------------------------------------ ------------------------------
|
| ------------------------------------- ------------------------------ 
| |         facebook.com routine      | |  facebook.com routine      |---->
| ------------------------------------- ------------------------------










---

Behind the scenes:

go is going to use one CPU core and only one Go Routine will run at one time. The Go Scheduler is the manager of the routines and will execute them via the Go Scheduler.

Routines work different when your computer has multiple CPU Cores. For computers with one core, the scheduler works very quickly behind the scenes and by default GO only uses one CPU core. 

                    -----------------
                    |    One CPU    |
                    |      Core     |
                    -----------------
                        |      ^
                        V      |
            -----------------------------------
            |           Go Scheduler          |
            -----------------------------------
              ^              ^               ^
              |              |               |
              V              V               V
        -------------- -------------- --------------
        | Go Routine | | Go Routine | | Go Routine |
        -------------- -------------- --------------

If we use multiple CPU cores, each core can work/run each routine within the Go language.

        -----------------   -----------------  -----------------
        |    One CPU    |   |    One CPU    |  |    One CPU    |
        |      Core     |   |      Core     |  |      Core     |
        -----------------   -----------------  -----------------
            |      ^            |      ^            |      ^
            V      |            V      |            V      |
            ----------------------------------------------------
            |                   Go Scheduler                    |
            ----------------------------------------------------
              ^                     ^                    ^
              |                     |                    |
              V                     V                    V
        --------------      --------------       ---------------
        | Go Routine |      | Go Routine |       |  Go Routine |
        --------------      --------------       ---------------

Schedule runs one thread one each "logical" core - but by default Go tries to use one core!

*Concurency is not parallelism.*

Concurrency - We can have multiple threads executing code. If one thread blocks, another one is picked up and worked on.

Parallelism - Multiple threads executed at the exact same time to the nanosecond. Requires multiple CPU's.

## Channels

Channels communicate between Go Routines and routines can only be done between the routines.
Created the same way we create other ‘values’ in Go.
These are typed and the data between them must be the same type


Go will throw errors if you put bool values within a channel that expects a string or int.

