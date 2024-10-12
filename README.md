Experimenting with Golang threading.

* thSimple - One parent thread, multiple child threads. No data passed between them.
* thTalk -
  - One parent thread, multiple child threads.
  - Each child receives a request message, sleeps a bit, and then replies.
  - The parent thread waits to get a reply from each child.
