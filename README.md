# japfp-go - Just Another PixelFlut Programm - GoLang
A client for the "[Pixelflut: Multiplayer canvas](https://github.com/defnull/pixelflut)".

## Ideas
- Edge detection and then draw the Edges in a seperate thread from the rest to make the image more visible quicker
- Starting multiple goroutines with some delay to fix anything that's been erased
- Drawing in sections (like splitting the image into 4 parts) and drawing each section in a seperate thread concurrently
- Using Watcher-Threads that detect if the integrity of the image has been damaged that then dispatch fixer threads to fix the damaged part(s)
- Start multiple connections

Note: I'm throwing around the terms "thread" and "concurrency" with little regard to their technical accuracy. I feel like this is ok in this case, because this is just supposed to be a small collection of some ideas I've had on how to (help) implement this.

## Features
- [ ] Basic features:
  - [ ] Read image from filesystem
  - [ ] Write image to pixelflut server
  - [ ] Scale image
- [ ] Advanced stuff
  - [ ] Decide which ideas to implement
  - [ ] Implement them
- [ ] REPL to:
  - [ ] Get stats
  - [ ] Send manual commands to the server
  - [ ] Change image path
  - [ ] Change offset
  - [ ] Change scaling
- [ ] Arguments for specifying:
  - [ ] Port
  - [ ] Server IP
  - [ ] Image path
  - [ ] Image scale
  - [ ] Offset of where to put the image

This is just everything that came to mind, so some might not make sense.
