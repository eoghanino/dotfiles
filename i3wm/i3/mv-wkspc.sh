#!/bin/bash
CURRENT=$(i3-msg -t get_workspaces|jq -r ".[]|select(.visible and .focused).name")
NEW=$(zenity --text="Enter new name:" --entry --title="Rename workspace $CURRENT" --entry-text="$CURRENT")
#i3-input -F "Rename workspace $CURRENT to %s" -P "New Name: "
i3-msg "rename workspace \"$CURRENT\" to \"$NEW\""
