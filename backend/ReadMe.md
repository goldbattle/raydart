## Raydart Backend

This backend is a rtmp server that allows a streamer to broadcast to it, and also allow for a client to connect to get the rtmp stream.
This will eventually have authentication with the frontend database to ensure that only valid people can stream to the server,
and also ensure that people that want to play a stream are valid. Without the backend running the frontend site will not have any content
to play, nor have anything to exist.


## TODOs

* Some type of stat and monitoring system
* Add authentication for the client
* Add authentication for the streamer
* Added transcoding features for the incoming stream


## Libraries Used

* joy4 - https://github.com/nareix/joy4

