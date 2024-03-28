# Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`