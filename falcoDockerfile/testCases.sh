#!/bin/sh

echo ""
echo "Starting script"
echo ""
echo "Installing sudo"
echo ""
apt-get update
apt-get install sudo
echo ""

# Test Case 1
echo "Case 1 - Writing below binary directory"
cd /		# go into root directory
cd bin		# go into binary directory
sudo touch znew	# attempt to edit znew file

# Test Case 2
echo ""
echo "Case 2 - Deleting bash history with 'history -c && history -w'"
echo "Current bash history:"
history				# show bash history
history -c && history -w	# delete bash history
echo "History after deleting:"
history				# display history after deleting

# Test Case 3
echo ""
echo "Case 3 - Deleting bash history with 'wipe ~/.bash_history'"
echo "Install wipe:"
echo Yes | sudo apt install wipe		# install wipe in case not there
echo "Wipe bash history"
echo Yes | wipe ~/.bash_history		# wipe bash history

# Test Case 4 - don't need for testing in minikube
# echo Case 4 - Running docker ubuntu container
# docker run --rm -ti ubuntu bash

# Test Case 4
echo ""
echo "Case 4 - Creating hidden directory and hidden file within such 
directory"
cd /
sudo mkdir .hiddenDir
cd .hiddenDir
sudo touch .hiddenFile.txt
cd /

# Test Case 5
echo ""
echo "Case 5 - Attempting to clear log activities (auth.log) using 
truncate -s 0"
cd /
cd var/log			# navigate to log directory
echo "auth.log before clearing:"
echo ""
sudo cat auth.log		# display log file before clearing
echo ""
sudo truncate -s 0 auth.log	# setting file size to 0 and removing 
contents
echo "auth.log after clearing:"
sudo cat auth.log	

# Test Case 6
echo ""
echo "Case 6 - Finding log files and piping output to a text file"
echo ""
echo "logFind.txt:"
find /var/log -type f -iname '*log' > logFind.txt
cat logFind.txt

# Test Case 7
echo ""
echo "Case 7 - Searching for passwords"
echo ""
find "id_rsa" > passwordFind.txt
cat passwordFind.txt
