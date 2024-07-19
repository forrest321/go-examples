# go-examples
Examples of various Go constructs  
**_This is for demo purposes only and not how a production project should be organized._**  
Basics of declaring variables, structs, slices, interfaces, mutexes, creating and consuming REST endpoints,
concurrency and parallelism. Code found online is noted in source (sorting algorithms), all else was hand written.

## Basics
Variable, constant, struct declarations, variable switching without intermediates, printing to console, 
simple math, if / else, switch, for loop with counters, range, for (while) loop.

## Concurrency
Use of channels, go routines, waitgroups, select statements, parallelism.

## Inheritance
Demonstrates Go's version of inheritance of properties from anonymous structs.

## Interfaces
Extends the inheritance examples through use of interfaces, which are implemented differently than  
other languages. Go's interfaces are implemented through what is sometimes called "duck typing". "If it walks like 
a duck, talks like a duck, then it must be a duck." 

## Mutexes
Basic demonstration of declaring and using mutexes and sync.Map.

## REST
Demonstrates creating a REST API with standard GET, POST, PUT, and DELETE endpoints using the standard libraries and the 
new routing capabilities introduced in Go 1.22. Once the endpoints are created, they are consumed using Go's built in http 
client functionality.

## Slices
Use of slices from sorting slices of integers to ranging over slices of strings. Generics used to reduce code by handling 
any type of integers. Sorting algorithms taken from https://github.com/0xAX/go-algorithms/tree/master/sorting and adapted
to generics. 

## To run
clone repo  
cd into folder  
run code  
```shell
git clone https://github.com/forrest321/go-examples.git
cd go-examples
```
Then run with Go  
```shell
go run . //need to run the whole directory. go run main.go will not work.
```
or  
Build and run  
```shell
go build
./go-examples
```

Example output:  
```text
Hello World
-=-=-=-=-=-=-=-=-=-=-=-
Basics
-=-=-=-=-=-=-=-=-=-=-=-
package level variable
I'm a constant
{x:1 y:2 name:b Exported:Exported notexported:Not Exported}
e, f, g, h := 1, 2, 3, 4
e, f, g, h = h, g, f, e
4 3 2 1
e, f, g, h = e*e, f*f, g*g, h*h
16 9 4 1
e, f, g, h = e/√e, f/√f, g/√g, h/√h
4 3 2 1
e, f, g, h = h, g, f, e
1 2 3 4
e is not greater than f
e == 1
For loop
i=0,1,2,3,4,5,6,7,8,9
Range
It is easy to range over a string as a byte slice
While loop
random number = 17
random number = 17
random number = 8
random number = 8
random number = 13
random number = 2
random number = 5
It matches!

-=-=-=-=-=-=-=-=-=-=-=-
Slices
-=-=-=-=-=-=-=-=-=-=-=-
ints slice: [1 2 3 4 5]
len is: 5
cap is: 5
sorted: true
Shuffling and sorting...
Sorting complete
gnome sort took 96 nanoseconds
counting sort took 212 nanoseconds
heap sort took 231 nanoseconds
default sort took 411 nanoseconds
bubble sort took 208 nanoseconds
quick sort took 860 nanoseconds
cocktail sort took 178 nanoseconds
bigSlice len is: 15000
bigSlice cap is: 15000
new bigSlice sorted: false
Shuffling and sorting...
Sorting complete
bubble sort took 222754 microseconds
counting sort took 4247198 microseconds
heap sort took 973 microseconds
quick sort took 4561 microseconds
default sort took 1149 microseconds
gnome sort took 88042 microseconds
cocktail sort took 91048 microseconds
-=-=-=-=-=-=-=-=-=-=-=-
Inheritance
-=-=-=-=-=-=-=-=-=-=-=-
basic animal: {name:animal legs:0 fur:false feathers:false sound:}
fido: {animal:{name:dog legs:4 fur:true feathers:false sound:bark} leash:yaaay lets go outside}
felix: {animal:{name:cat legs:4 fur:true feathers:false sound:meow} leash:you will pay for this hairballs:true}
eagle: {bird:{animal:{name:eagle legs:2 fur:false feathers:true sound:screech} wings:2} talons:sharp}
-=-=-=-=-=-=-=-=-=-=-=-
Interfaces
-=-=-=-=-=-=-=-=-=-=-=-
bark
yaaay lets go outside
all goes well
meow
you will pay for this
this was a bad idea
screech
-=-=-=-=-=-=-=-=-=-=-=-
Mutexes
-=-=-=-=-=-=-=-=-=-=-=-
created mutexThing
unsafe map before: map[0:value 0 1:value 1]
safe map before:
key: 0, value: value 0
key: 1, value: value 1
unsafe map after: map[0:safely changed values 1:value 1]
safe map after:
key: 0, value: changed values
key: 1, value: this is fine

-=-=-=-=-=-=-=-=-=-=-=-
Concurrency
-=-=-=-=-=-=-=-=-=-=-=-
Channel #2 counted from 2501 to 3750 in 0.00954654 seconds
Channel #7 counted from 8751 to 10000 in 0.010090777 seconds
Channel #8 counted from 10001 to 11250 in 0.010619358 seconds
Channel #10 counted from 12501 to 13750 in 0.011058665 seconds
Channel #9 counted from 11251 to 12500 in 0.012215469 seconds
Channel #5 counted from 6251 to 7500 in 0.012504371 seconds
Channel #0 counted from 1 to 1250 in 0.012801832 seconds
Channel #3 counted from 3751 to 5000 in 0.013412862 seconds
Channel #1 counted from 1251 to 2500 in 0.013536852 seconds
Channel #4 counted from 5001 to 6250 in 0.013555448 seconds
Channel #6 counted from 7501 to 8750 in 0.014137903 seconds
Channel #11 counted from 13751 to 15000 in 0.014242592 seconds

Started: 2024-07-19 13:02:54.03581744 -0500 CDT m=+4.658083416 
Ended: 2024-07-19 13:02:54.050116132 -0500 CDT m=+4.672382109 
Duration: 14.298693ms
Used 12 processors
Counted to: 15000
-=-=-=-=-=-=-=-=-=-=-=-
REST
-=-=-=-=-=-=-=-=-=-=-=-
Router created
Starting server
Calling GET endpoint
Response received: {"id":1,"name":"data-1","description":"This is data #1"}

Calling GET endpoint
Response received: {"id":5,"name":"data-5","description":"This is data #5"}

Getting data for id=11. This should get a 404:
Calling GET endpoint
Status : 404 Not Found
Calling POST endpoint
Response received: 201 Created
Getting data for id=11. This should succeed now
Calling GET endpoint
Response received: {"id":11,"name":"data-11","description":"This is data 11"}

Object to update:
Calling GET endpoint
Response received: {"id":7,"name":"data-7","description":"This is data #7"}

Calling PUT endpoint
Response received: 200 OK
Getting data for id=7. This should be updated now
Calling GET endpoint
Response received: {"id":7,"name":"data-777","description":"Updated data 777"}

Calling GET endpoint
Response received: {"id":4,"name":"data-4","description":"This is data #4"}

Calling Delete endpoint
Response received: 200 OK
Getting data for id=4 This should get a 404
Calling GET endpoint
Status : 404 Not Found
Calling GET endpoint
Response received: {"id":6,"name":"data-6","description":"This is data #6"}

Calling Delete endpoint
Response received: 200 OK
Getting data for id=6 This should get a 404
Calling GET endpoint
Status : 404 Not Found

Process finished with the exit code 0


```