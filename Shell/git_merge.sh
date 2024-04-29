#!/bin/bash

cd $PRAC_DIR
curr_branch=$(git rev-parse --abbrev-ref HEAD)

echo "Current branch: $curr_branch"
git checkout main
git pull origin main
git checkout $curr_branch
