#!/bin/sh
git push stylize "$1"
git push gitea "$1"
git push github "$1"