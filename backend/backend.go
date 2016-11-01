package main

import (
	"fmt"
	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pktque"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
	"sync"
	"time"
)

func init() {
	format.RegisterAll()
}

type FrameDropper struct {
	Interval     int
	n            int
	skipping     bool
	DelaySkip    time.Duration
	lasttime     time.Time
	lastpkttime  time.Duration
	delay        time.Duration
	SkipInterval int
}

func (self *FrameDropper) ModifyPacket(pkt *av.Packet, streams []av.CodecData, videoidx int, audioidx int) (drop bool, err error) {
	if self.DelaySkip != 0 && pkt.Idx == int8(videoidx) {
		now := time.Now()
		if !self.lasttime.IsZero() {
			realdiff := now.Sub(self.lasttime)
			pktdiff := pkt.Time - self.lastpkttime
			self.delay += realdiff - pktdiff
		}
		self.lasttime = time.Now()
		self.lastpkttime = pkt.Time

		if !self.skipping {
			if self.delay > self.DelaySkip {
				self.skipping = true
				self.delay = 0
			}
		} else {
			if pkt.IsKeyFrame {
				self.skipping = false
			}
		}
		if self.skipping {
			drop = true
		}

		if self.SkipInterval != 0 && pkt.IsKeyFrame {
			if self.n == self.SkipInterval {
				self.n = 0
				self.skipping = true
			}
			self.n++
		}
	}

	if self.Interval != 0 {
		if self.n >= self.Interval && pkt.Idx == int8(videoidx) && !pkt.IsKeyFrame {
			drop = true
			self.n = 0
		}
		self.n++
	}

	return
}

func main() {

	// Enable debug
	//rtmp.Debug = true

	// Create rtmp server
	server := &rtmp.Server{}

	// Create a mutex for access to channel listing
	l := &sync.RWMutex{}
	type Channel struct {
		que *pubsub.Queue
	}
	channels := map[string]*Channel{}

	// HandlePlay function, this gets called when a new client
	// is connected to the server. In here we parse the channel,
	// and then serve the corresponding content to the user
	// TODO: Add authentication here to ensure that the user has access
	// TODO: to the matching channel, and reject if not
	server.HandlePlay = func(conn *rtmp.Conn) {
		// Find the matching channel from our channel list
		l.RLock()
		ch := channels[conn.URL.Path]
		l.RUnlock()

		// If we do not have a channel, we have nothing to server
		// TODO: Add authentication here
		if ch == nil {
			return
		}

		// For that channel get the latest position in the byte channel
		cursor := ch.que.Latest()
		query := conn.URL.Query()

		// TODO: Figure out what this stuff does
		if q := query.Get("delaygop"); q != "" {
			n := 0
			fmt.Sscanf(q, "%d", &n)
			cursor = ch.que.DelayedGopCount(n)
		} else if q := query.Get("delaytime"); q != "" {
			dur, _ := time.ParseDuration(q)
			cursor = ch.que.DelayedTime(dur)
		}

		filters := pktque.Filters{}
		filters = append(filters, &pktque.FixTime{StartFromZero: true, MakeIncrement: true})

		// This is can compress and change the packets of the stream
		// This could be a place where we can add other quality versions
		demuxer := &pktque.FilterDemuxer{
			Filter:  filters,
			Demuxer: cursor,
		}

		// Finally start sending the file to the connected client
		fmt.Printf("Client Connected: %s\n", query)
		avutil.CopyFile(conn, demuxer)
	}

	// HandlePublish function allows for a client using OBS or other streaming
	// software to connect to the server and broadcast a stream. In the future
	// we will want to implement something such as stream_keys that only allow
	// verified people to stream to the server and serve the content
	server.HandlePublish = func(conn *rtmp.Conn) {
		l.Lock()
		ch := channels[conn.URL.Path]
		if ch == nil {
			ch = &Channel{}
			ch.que = pubsub.NewQueue()
			query := conn.URL.Query()
			if q := query.Get("cachegop"); q != "" {
				var n int
				fmt.Sscanf(q, "%d", &n)
				ch.que.SetMaxGopCount(n)
			}
			channels[conn.URL.Path] = ch
		} else {
			ch = nil
		}
		l.Unlock()
		if ch == nil {
			return
		}

		fmt.Printf("Channel Starting: %s\n",conn.URL.Path)
		avutil.CopyFile(ch.que, conn)

		l.Lock()
		delete(channels, conn.URL.Path)
		l.Unlock()
		ch.que.Close()
		fmt.Printf("Channel Closed: %s\n",conn.URL.Path)
	}

	// Set the address of the server
	// TODO: Change to config file value
	server.Addr = ":1935"
	// Listen for incoming data, and serve it
	server.ListenAndServe()

}
