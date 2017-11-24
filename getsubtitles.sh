IFS_BAK=$IFS
IFS="
"

for line in $NAUTILUS_SCRIPT_SELECTED_FILE_PATHS; do
	getsubtitles #your program compiled
        notify-send $line
done

IFS=$IFS_BAK