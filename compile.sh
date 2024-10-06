#/bin/bash

set -e

read -e -p "Give the name of the plugin: " plugin
# get absolute path
pluginpath=$(realpath $plugin)
pluginName=$(basename $plugin)


echo "Work for $pluginName ..."
echo  -n "Verification for: '$pluginpath' ... "


if [ -d "$pluginpath" ]; then
    echo "folder exists."
else

    echo -e "\n >>> This folder doesn't exist"
    exit 30
fi


cd $pluginpath/Source

# Dependency file
gomodPath=$pluginpath/Source/go.mod
echo "Checking the presence of the dependency file: $gomodPath ..."
if [ -f $gomodPath ]; then
    echo "go.mod exists"
else
    echo "Creating dependancy file go.mod"
    read -p "Enter git/github url for the plugin: " gitURL
    
    gitURL=$(echo $gitURL| sed 's,https://,,g')
    gitURL=$(echo $gitURL| sed 's,http://,,g')

    go mod init $gitURL
    
fi

# Dependencies
echo "Building dependencies... "
go mod tidy

echo "Dependencies built."


# Building

cd $pluginpath/Source/server

echo "Compiling $gofiles ... "

gofiles=$(ls *.go)


mkdir -p ../../dist
go build -o ../../dist/plugin.exe $gofiles
cp ../plugin.json ../../dist/

echo "Compiling done."

echo "Compression of the plugin"



cd ../../dist
tar -czvf "mattermost-plugin-$pluginName".tar.gz plugin.exe plugin.json 

echo "Move plugin"
mv "mattermost-plugin-$pluginName.tar.gz" /mnt/hgfs/Write/Dev/




echo "Copy of the sources"
cd ..
target="/mnt/hgfs/Write/Dev/"
#mkdir -p $target/Source/

#cd FlowWatcher

#cp $gofiles $target/Source/
#cp go.mod go.sum plugin.json $target/

cp -Rd $pluginpath $target

set +e
