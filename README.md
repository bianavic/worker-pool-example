# Worker pool to encode videos

Worker pools is a design in which a fixed number of X workers (Go goroutines) works on Y tasks in a work queue (Go channel).
The work resides in a queue until a worker finish its current task and pull a new one.


### tools:
- ffmpeg is a powerful tool to encode and decode videos. It is a command line tool that can be used to encode and decode videos in different formats.
- goffmpeg is a Go wrapper for FFmpeg. It allows you to encode and decode videos in Go.
HLS: multiples resolutions in a single m3u8 file (1080p, 720p, 480p)
