Welcome to snow stake!

# what?

The snow stake is just siting there, lonely.  It looks like this:

![image](https://www.mtbachelor.com/webcams/snowstake.jpg?44862)

Let's make a recording of the snow stake each time it changes.  Cause we can.

# how to install?

this is golang, motherfucker!  It is cross-compiled!  Choose your arch from `builds/*` and you're gtg.  The following builds are available.


```
darwin-amd64/stake
linux-amd64/stake
linux-arm/stake
```

Pick your executable and run it.

# what does it do?

Each time `./stake*` is run, we

* do a HEAD call to the snow stake jpg
* compare the etag of the HEAD call to the last call we made
* if the etag has changed, save the jpg in the format `2017-10-13-05-56.jpg`

# what about cron?

Assuming you moved your binary to `/opt/stake/stake`, you'd 

### edit your cron

to look like this to make it run every 10th minute

```
*/10 * * * * cd /opt/stake && ./stake
```