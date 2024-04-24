#!/bin/bash

curr_branch=$(git rev-parse --abbrev-ref HEAD)

cd ~/Projects/Practice
git checkout main
git pull origin main
git checkout $(curr_branch)
