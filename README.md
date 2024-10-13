Experimenting with Golang threading.

* thSimple - One parent thread, multiple child threads. No data passed between them.
* thTalk -
  - One parent thread, multiple child threads.
  - Each child receives a request message, sleeps a bit, and then replies.
  - The parent thread waits to get a reply from each child.

Getting the O/S thread ID across platforms to work in package helpers was a lot of fun!

So, as Go low-level documentation states, they use O/S threads although they are shared across the launched Goroutines.
