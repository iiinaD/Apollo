#!/bin/bash

# Remove the apollo binary
sudo rm -f /usr/local/bin/apollo

# Check the SHELL variable and remove environment variables and comments
SHELL_RC_FILE=""
if [[ $SHELL == *"bash"* ]]; then
  SHELL_RC_FILE=".bashrc"
elif [[ $SHELL == *"zsh"* ]]; then
  SHELL_RC_FILE=".zshrc"
fi

if [[ -n "$SHELL_RC_FILE" ]]; then
  if grep -q "export GEMENI_API_KEY=" ~/$SHELL_RC_FILE; then
    sed -i '/# api for google'"'"'s generative ai '"'"'gemeni'"'"'/d' ~/$SHELL_RC_FILE
    sed -i '/export GEMENI_API_KEY=/d' ~/$SHELL_RC_FILE
    echo "Removed GEMENI_API_KEY and its comment from $SHELL_RC_FILE"
  else
    echo "GEMENI_API_KEY not found in $SHELL_RC_FILE"
  fi

  if grep -q "export APOLLO_OUTPUT_THEME=" ~/$SHELL_RC_FILE; then
    sed -i '/# this is the custom theme for Apollo/d' ~/$SHELL_RC_FILE
    sed -i '/export APOLLO_OUTPUT_THEME=/d' ~/$SHELL_RC_FILE
    echo "Removed APOLLO_OUTPUT_THEME and its comment from $SHELL_RC_FILE"
  else
    echo "APOLLO_OUTPUT_THEME not found in $SHELL_RC_FILE"
  fi

  echo "Uninstallation complete!"
  exec $SHELL
else
  echo "Unsupported shell. Please use bash or zsh."
fi