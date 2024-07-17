# go-examples
Examples of various Go constructs  
**_This is for demo purposes only and not how a production project should be organized._**  
Basics of declaring variables, structs, 

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
go run . //need to run the whole directory. go run main.go will not work
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
{1 2 b Exported Not Exported}
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
random number = 9
random number = 14
random number = 16
random number = 5
It matches!

-=-=-=-=-=-=-=-=-=-=-=-
Slices
-=-=-=-=-=-=-=-=-=-=-=-
ints slice: [1 2 3 4 5]
len is: 5
cap is: 5
sorted: true
shuffled slice: [2 5 1 4 3]
sorted: false
bubble sorted: [1 2 3 4 5]
sorted: true
shuffled slice: [1 4 5 2 3]
sorted: false
quick sorted: [1 2 3 4 5]
sorted: true
bigSlice len is: 15000
bigSlice cap is: 15000
new bigSlice sorted: false
bubble sort of bigSlice took: 0.213256469 seconds
bigSlice sorted: true
shuffled bigSlice sorted: false
quick sort of bigSlice took: 0.004424522 seconds
bigSlice sorted: true
quick sort faster by 0.208831947 seconds
-=-=-=-=-=-=-=-=-=-=-=-
Inheritance
-=-=-=-=-=-=-=-=-=-=-=-
basic animal: {name:animal legs:0 fur:false feathers:false sound:}
fido: {animal:{name:dog legs:4 fur:true feathers:false sound:bark} leash:yaaay lets go outside}
felix: {animal:{name:cat legs:4 fur:true feathers:false sound:meow} leash:you will pay for this hairballs:true}
eagle: {bird:{animal:{name:eagle legs:2 fur:false feathers:true sound:screech} wings:2} talons:sharp}
-=-=-=-=-=-=-=-=-=-=-=-
Mutexes
-=-=-=-=-=-=-=-=-=-=-=-
created mutexThing: &{mu:{state:0 sema:0} unsafeMap:map[0:value 0 1:value 1] safeMap:{mu:{state:0 sema:0} read:{_:[] _:{} v:0xc000016050} dirty:map[0:0xc00011a050 1:0xc00011a058] misses:0}}

-=-=-=-=-=-=-=-=-=-=-=-
Concurrency
-=-=-=-=-=-=-=-=-=-=-=-
Channel #1 counted from 1251 to 2500 in 0.008074704 seconds
Channel #6 counted from 7501 to 8750 in 0.00878178 seconds
Channel #5 counted from 6251 to 7500 in 0.009834425 seconds
Channel #2 counted from 2501 to 3750 in 0.013455482 seconds
Channel #7 counted from 8751 to 10000 in 0.015092619 seconds
Channel #8 counted from 10001 to 11250 in 0.01536675 seconds
Channel #10 counted from 12501 to 13750 in 0.016154528 seconds
Channel #0 counted from 1 to 1250 in 0.016422908 seconds
Channel #11 counted from 13751 to 15000 in 0.016747667 seconds
Channel #3 counted from 3751 to 5000 in 0.017348844 seconds
Channel #9 counted from 11251 to 12500 in 0.01773115 seconds
Channel #4 counted from 5001 to 6250 in 0.017988914 seconds

Started: 2024-07-16 13:38:58.507672955 -0500 CDT m=+0.218573938 
Ended: 2024-07-16 13:38:58.525688587 -0500 CDT m=+0.236589569 
Duration: 18.015631ms
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

```