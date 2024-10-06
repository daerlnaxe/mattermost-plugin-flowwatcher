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
- [ ] A Cleaner based on a word
- [ ] Bind a command to force flow (enable/disable work)
- [ ] Verify it works with json feed
- [ ] Configuration
  - [ ] Debug
  - [ ] Set Sleeper
  - [ ] Set reactions that could stop the cleaner
  - [ ] Disabling cleaner for other users than bot
- [ ] Auth Login for rss with authentification (later)
- [ ] Pass Cloud Flare verification (later)
- [x] Make a system to avoid to refresh if not needed
- [x] Bind core to subs
- [x] Unsub
- [x] Flow Activation
  - [x] Filter if flow disabled
  - [x] Indicate status in ls
  - [x] Activate flow
  - [x] Stop flow
  - [x] Force flow (useful when flow is activted)  
  - [x] Automatize stop flow in case of error
- [x] A cleaner for the channel (each user must leave and back to the channel, to avoir "deleted message")
  - [x] Don't clean if there is reaction (white_check_mark for the while)
  - [x] Can clean for others User


<br>
<br
  
# Build
go build -o plugin.exe plugin.go activate.go command.go && tar -czvf plugin.tar.gz plugin.exe plugin.json && cp plugin.tar.gz /mnt/hgfs/Write/Dev/
