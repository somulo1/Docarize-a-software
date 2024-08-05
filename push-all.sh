#!/bin/sh
git push origin "$1"
git push gitea-style "$1"
git push github-style "$1"