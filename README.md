# maze-tool
Toy maze generator project

## About

Strictly hobby project. Attempting to refamiliarise myself with Golang whilst working my way through the excellent book [Mazes For Programmers](http://www.mazesforprogrammers.com/) by Jamis Buck

I haven't got too far through this yet. Don't use it as a reference, put your hand in your pocket and buy the book!

## Usage

Build using the Makefile
```
make
```

Generate a maze
```
./maze -width=25 -height=10 -algorithm=aldousbroder -pngfile=/tmp/mymaze.png
```

That will give you something like this
```
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|  |              |     |  |        |                 |           |        |
+  +--+  +  +--+--+  +--+  +  +  +--+  +--+  +  +--+--+  +--+--+  +  +--+--+
|        |  |                 |        |     |     |  |     |  |        |  |
+--+  +  +  +  +--+  +  +--+  +--+--+  +  +--+  +--+  +  +  +  +  +  +  +  +
|  |  |  |  |  |  |  |  |  |     |  |  |     |  |     |  |     |  |  |  |  |
+  +--+  +  +--+  +  +  +  +--+--+  +--+--+  +--+--+  +  +--+  +  +--+  +  +
|        |           |           |           |           |     |  |        |
+--+--+--+--+--+  +--+  +--+  +  +--+  +--+--+--+  +--+  +  +--+  +--+--+  +
|  |  |     |  |     |     |  |     |  |  |     |  |     |     |  |     |  |
+  +  +  +  +  +--+--+  +  +--+--+  +--+  +--+  +--+--+--+  +--+--+  +--+--+
|  |     |        |     |  |        |                                   |  |
+  +--+  +--+--+  +--+--+--+--+--+  +  +--+  +--+--+  +  +--+--+--+  +--+  +
|  |        |     |     |  |  |        |     |        |           |     |  |
+  +--+--+--+  +--+--+  +  +  +  +--+  +  +  +--+  +  +  +--+--+--+--+--+  +
|     |     |  |           |     |  |  |  |     |  |  |                    |
+  +  +  +  +  +  +--+  +--+  +--+  +  +--+  +--+--+--+--+--+--+--+--+  +  +
|  |  |  |     |     |        |     |  |  |     |        |     |        |  |
+  +--+  +  +--+  +--+  +  +  +  +--+--+  +  +  +--+  +  +  +--+  +  +--+  +
|        |           |  |  |     |           |  |     |           |  |     |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
```

## TODO

Entrance and exits, implement more algorithms, maze analysis, and bitfield maze representation
