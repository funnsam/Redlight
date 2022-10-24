@echo off
set commit=%1
set version=%2
git add --all
git commit -m %commit%
git tag %version%
git push -f -u origin HEAD:main
@echo on