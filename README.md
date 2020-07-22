## Pineapple Pen

The main purpose of this program is *drawing*, but directly on the screen.  

If you are confused, one famous program with the same purpose is [Epic Pen](https://epic-pen.com/), but it's not available for Linux.


### Features
The only feature available in the moment is drawing on the screen moving the mouse. Like the image below.
![feature-example](./img/example.png)

### Download
```
git clone https://github.com/Valeyard1/pineapple-pen.git
go run cmd/pineapple-pen/main.go
```
## Contributing
There's some features that needs to be implemented:
* Improve the drawing to make it faster
* Undo
* Erase

The project structure is following [this](https://github.com/golang-standards/project-layout) layout, so I think the struct won't be a problem.  
Where to start: callbacks.go, that's the main file, where all the important stuff is.


