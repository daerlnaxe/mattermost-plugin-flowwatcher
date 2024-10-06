# Presentation
FlowWatcher use `gofeed` package and allow to post on a `Mattermost` channela feed from RSS (wip)

In theory it will support JSON also... in all cases, at the end it will support. I'm not sure for the while but i will see to handle another types of flux, like APIs.

<br>

## Version
Alpha 01
- Not really operationnal, for the while it's only binded to a RSS Source.
- Both core subscribtion and posting are working.
- need to finish.

<br>

## Todo
- [ ] A cleaner for the channel
- [ ] A Cleaner based on a word
- [ ] Bind core to subs
- [ ] Unsub
- [ ] Verify it works with json feed
- [ ] Make a system to avoid to refresh if not needed
- [ ] Configuration
  - [ ] Debug
  - [ ] Set Sleeper
- [ ] Auth Login for rss with authentification (later)
- [ ]   Pass Cloud Flare verification (later)

<br>
<br
  
# Build
go build -o plugin.exe plugin.go activate.go command.go && tar -czvf plugin.tar.gz plugin.exe plugin.json && cp plugin.tar.gz /mnt/hgfs/Write/Dev/
