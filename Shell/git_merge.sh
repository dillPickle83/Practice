#!/bin/bash

curr_branch=$(git rev-parse --abbrev-ref HEAD)

echo "Current branch: $curr_branch"
cd $PRAC_DIR
git checkout main
git pull origin main
git checkout $(curr_branch)
