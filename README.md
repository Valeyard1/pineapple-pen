## Pineapple Pen

The main purpose of this program is *drawing*, but directly on the screen.  

If you are confused, one famous program with the same purpose is [Epic Pen](https://epic-pen.com/), but it's not available for Linux.


### Features
The only feature available in the moment is drawing on the screen moving the mouse. But with some delay, I don't know why.

### Download
```
git clone https://github.com/Valeyard1/pineapple-pen.git && cd pineapple-pen
cargo run
```
## Contributing
There's some features that needs to be implemented:
* Improve the drawing to make it faster
* Undo
* Erase

## Where should I begin contributing?
As the program grows, I'll create issues with subjects you can help, but for now,
you need to have a minimum knowledge in GTK (how it works, signals and callbacks functions, etc.),
after that, you will probably understand the `src/main.rs` file, then you can try to develop a callback function
in the _src/callbacks.rs_

At first I tried to write this with Go, but all the GTK libraries in Go aren't finished, you can see it [here](https://github.com/Valeyard1/pineapple-pen/commit/d15205f1744000a632f6b087c784dcb2dc8c40ba)

