#!/bin/bash
function selectWorkspaceToDisplay(){
	echo "Which of the currently active workspaces do you want to share?";
	echo "$(tput setaf 2) $(i3-msg -t get_outputs| jq -r '.[]|select(.active == true )|.current_workspace') $(tput sgr0)"
	read CURRENT_WORKSPACE;
	export CURRENT_WORKSPACE;
}
function retrieveClipInfoForWorkspace(){
	local WORKSPACE=\"$CURRENT_WORKSPACE\";
	local WORKSPACE_DETAILS=$(i3-msg -t get_outputs| jq ".[]|select(.current_workspace == "$WORKSPACE")|.")
	local OFFSETX=$(echo $WORKSPACE_DETAILS|jq '.rect.x')
	local OFFSETY=$(echo $WORKSPACE_DETAILS|jq '.rect.y')
	local X=$(echo $WORKSPACE_DETAILS|jq '.rect.width')
	local Y=$(echo $WORKSPACE_DETAILS|jq '.rect.height')
	echo "${X}x${Y}+${OFFSETX}+${OFFSETY}"
}
selectWorkspaceToDisplay;
CLIP_DETAILS=$(retrieveClipInfoForWorkspace);
echo "Clip details are: $CLIP_DETAILS"
x11vnc -display :0 -clip $CLIP_DETAILS -noxinerama -passwdfile ~/.vnc/passwd
#x11vnc -display :0 -clip 1920x1200+1920+0 -noxinerama -passwdfile .vnc/passwd
#x11vnc -display :0 -clip 3440x1440+1920+0 -noxinerama -passwdfile .vnc/passwd
#x11vnc -display :0 -clip 1920x1200+1920+0 -noxinerama -passwdfile .vnc/passwd
#x11vnc -display :0 -clip 1920x1200+0+0 -noxinerama -passwdfile .vnc/passwd
