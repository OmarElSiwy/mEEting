#! /bin/sh

# To Run on Linux
# bash Expotest.sh
# or To Run on Windows or Linux
# ./Expotest.sh

# CD into directory of Go Files
cd /usr/src/app/backend || exit

# Download Toals if not present
go get golang.org/x/mobile/cmd/gomobile
gomobile init

# Bind Files
gomobile bind -target=android -o /usr/src/app/android/app/libs/backend.aar . || exit
# ONLY UNCOMMENT IF RUNNING ON MAC
# gomobile bind -target=ios -o /usr/src/app/ios/backend.framework . || exit

# Go one directory back
cd /usr/src/app

# Start the Application
npx expo start

