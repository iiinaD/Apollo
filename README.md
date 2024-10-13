# Apollo

## Description
Apollo is a commandline client for prompting Google's
Generative AI 'Gemeni'. This client is made purely in go using
Gemeni's [API](https://ai.google.dev/gemini-api) and [glamour](https://github.com/charmbracelet/glamour).

## Before Installation
Before installing Apollo you need to create you own Gemeni API key. 
This can be done [here](https://aistudio.google.com/app/apikey).
This key is need for the program to run.

## Installation
To install Apollo clone the repository and cd into the cloned folder.
Run ``` chmod +x ./install.sh; ./install.sh ``` inside the folder.
When prompted input your api-key.
Apollo is also going to ask you if you want to choose a custom theme.
You should be all done now.

## Usage
There are currently two ways to use Apollo. You can write ``` apollo "<your-prompt-here>" ``` if you just want a quick answer to something without starting a conversation.
Otherwise, you can just write ``` apollo ``` which will open the client and you can now have a conversation with Gemeni through your terminal.

## Themes
The themes Apollo has are specified by [glamour](https://github.com/charmbracelet/glamour) and can be found [here](https://github.com/charmbracelet/glamour/tree/master/styles/gallery).
Just take a look at each of the pngs and choose pick the one you love :)
To set the theme write the name of the png you like (without .png) when asked for in the installer.
If you already have Apollo installed you can manually set the theme with ``` export APOLLO_OUTPUT_THEME="<name-of-your-theme>" ```.
For this to be persistent write this line into your .bashrc or .zshrc file.