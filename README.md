

go build -o plugin.exe plugin.go activate.go command.go && tar -czvf plugin.tar.gz plugin.exe plugin.json && cp plugin.tar.gz /mnt/hgfs/Write/Dev/
