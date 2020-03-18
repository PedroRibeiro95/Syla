![alt text](assets/syla.png "syla")

`syla`, fancy acronym for "Share your listening activity", is a simple project implemented in Golang that allows one to share their listening activity (duh!) by fetching profile information directly from Spotify, similarly to LastFM. It implements a RESTful API, making it very easy to integrate in a personal website.

## Goals

This is still a very raw personal project, with development being done whenever spare time arises. However, the main intent is to have `syla` supporting several music stream providers. As of now, the main focus goes to Spotify (as it is the one I personally use), with support for Apple Music to be added afterwards.

Additionally, `syla` is only capable of providing "static" information; to be more precise, this is information that can be fetched at any time from Spotify API. There is a personal interest in developing a _watcher_ component for `syla`, which would follow the user's activity on Spotify and generate interesting data accordingly.

## Why should I ever use this instead of solutions such as LastFM?

Well, I think it really boils down to allowing you to control how you share your information. You have the code, you run the binary. Also, it's open source so you can add always contribute with more features! Yay!

## Next steps

* [ ] Fully implement the Spotify provider
  * [ ] Fetching favorite albums
  * [ ] Fetching favorite artists
  * [ ] Fetching favorite songs
* [ ] Refactor code
* [ ] Implement good abstractions
* [ ] Fully implement the Apple Music provider
  * [ ] Fetching favorite albums
  * [ ] Fetching favorite artists
  * [ ] Fetching favorite songs
* [ ] Conceptualize `watcher` (_big if_)