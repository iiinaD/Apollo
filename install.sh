#!/bin/bash

# Build the Go program
sudo go build -o apollo

# Prompt for Gemini API key
read -rp "Enter your Gemini API key: " GEMENI_API_KEY

# Prompt for theme (optional)
read -rp "Do you want to use a special theme (y/n)? " theme_choice
if [[ $theme_choice == "y" ]]; then
  read -rp "Enter your theme name: " APOLLO_OUTPUT_THEME
fi

# Move the built program
sudo mv apollo /usr/local/bin/apollo

# Set execute permissions on the binary
sudo chmod +x /usr/local/bin/apollo

# Check the SHELL variable and set environment variables
SHELL_RC_FILE=""
if [[ $SHELL == *"bash"* ]]; then
  SHELL_RC_FILE=".bashrc"
elif [[ $SHELL == *"zsh"* ]]; then
  SHELL_RC_FILE=".zshrc"
fi

if [[ -n "$SHELL_RC_FILE" ]]; then
  if ! grep -q "export GEMENI_API_KEY=" ~/$SHELL_RC_FILE; then
    {
      echo ""
      echo "# api for google's generative ai 'gemeni'"
      echo "export GEMENI_API_KEY=\"$GEMENI_API_KEY\""
    } >> ~/$SHELL_RC_FILE
  else
    echo "GEMENI_API_KEY is already set in $SHELL_RC_FILE"
  fi

  if [[ -n "$APOLLO_OUTPUT_THEME" ]]; then
    if ! grep -q "export APOLLO_OUTPUT_THEME=" ~/$SHELL_RC_FILE; then
      {
        echo ""
        echo "# this is the custom theme for Apollo"
        echo "export APOLLO_OUTPUT_THEME=\"$APOLLO_OUTPUT_THEME\""
      } >> ~/$SHELL_RC_FILE
    else
      echo "APOLLO_OUTPUT_THEME is already set in $SHELL_RC_FILE"
    fi
  fi

  echo "Installation complete!"
else
  echo "Unsupported shell. Please use bash or zsh."
fi