#!/bin/bash

# Need to download breach-parse and torrent the password list from http://www.github.com/hmaverickadams brech-parse repo [magnet link in readme]


if [ "$#" != "2" ]; then
	echo "Breach-Parse v2: A Breached Domain Parsing Tool by Heath Adams    [Rip off]"
	echo " "
	echo "Usage: breach-parse <domain to search> <file to output>"
	echo "Example: breach-parse @gmail.com gmail.txt"
	echo " "
	echo "For multiple domains: breach-parse '<domain to search>|<domain to search>' <file to output>"
	echo "Example: breach-parse '@gmail.com|@yahoo.com' multiple.txt"
	exit 1

else
	fullfile=$2
	fbname=$(basename "$$fullfile" | cut -d. -f1)
	master=$fbname-master.txt
	users=$fbname-user.txt
	passwords=$fbname-passwords.txt

	touch $master
	total_files=$(find /opt/breach-parse/BreachCompilation/data -type f | wc -l)
	file_count=0

	function ProgressBar {
		let _progress=(${file_count}*100/${total_files}*100)/100
		let _done=(${_progress}*4)/10
		let _left=40-$_done

		_fill=$(printf "%${_done}s")
		_empty=$(printf "%${_left}s")

	printf "\rProgress : [${_fill// /\#}${_empty// /-}] ${_progress}%%"
	}

	find /opt/breach-parse/BreaachCompilation/data -type f -print0 | while read -d $'\0' file
	do
		grep -a -E "$1" "$file" >> $master
		((++file_count))
		ProgressBar ${number} ${total_Files}
	done
fi

sleep 3

awk -F':' '{print $1}' $master > $users

sleep 1

awk -F':' '{print $2}' $master > $passwords
echo
exit 0
