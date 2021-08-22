#!/bin/zsh
repo='JSONhilder/strongbox'
default_path=$HOME'/.strongbox'
releases="https://api.github.com/repos/"$repo/"releases"
exe_path=$HOME'/.strongbox/strongbox'

echo "Determining latest release"

tag=$(curl --silent $releases/latest | grep -Po '"tag_name": "\K.*?(?=")')

VERSION=$(echo $tag|cut -c 2-6)

FILE="strongbox_"$VERSION"_linux_amd64.tar.gz"

download="https://github.com/"$repo"/releases/download/"$tag"/"$FILE

echo "Downloading latest release..."
wget -O $FILE $download

echo "Creating home directory at: "$default_path
mkdir -p $default_path

echo "Extracting contents to: "$default_path"/"
tar xf $FILE -C $default_path"/"

echo "Cleaning up..."
rm -rf $FILE

if [ -e ~/.zshrc ]; then
    if [ "$(grep -c "alias strongbox" ~/.zshrc)" -eq 0 ]; then
        echo "Adding zsh path to exe file..."
        echo 'alias strongbox="~/.strongbox/strongbox"' >> ~/.zshrc
    else
        echo "alias exists in zshrc"
    fi
fi

if [ -e ~/.bashrc ]; then
    if [ "$(grep -c "alias strongbox" ~/.bashrc)" -eq 0 ]; then
        echo "Adding bash path to exe file..."
        echo 'alias strongbox="~/.strongbox/strongbox"' >> ~/.bashrc
    else
        echo "alias exists in bashrc"
    fi
fi

#LINUX PREREQ
#[ curl ]