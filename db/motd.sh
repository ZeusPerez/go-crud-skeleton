#!/usr/bin/env bash

cat <<"EOF"
 .----------------. .----------------. .----------------. .----------------. .----------------.
| .--------------. | .--------------. | .--------------. | .--------------. | .--------------. |
| | ____    ____ | | |  ____  ____  | | |    _______   | | |    ___       | | |   _____      | |
| ||_   \  /   _|| | | |_  _||_  _| | | |   /  ___  |  | | |  .'   '.     | | |  |_   _|     | |
| |  |   \/   |  | | |   \ \  / /   | | |  |  (__ \_|  | | | /  .-.  \    | | |    | |       | |
| |  | |\  /| |  | | |    \ \/ /    | | |   '.___`-.   | | | | |   | |    | | |    | |   _   | |
| | _| |_\/_| |_ | | |    _|  |_    | | |  |`\____) |  | | | \  `-'  \_   | | |   _| |__/ |  | |
| ||_____||_____|| | |   |______|   | | |  |_______.'  | | |  `.___.\__|  | | |  |________|  | |
| |              | | |              | | |              | | |              | | |              | |
| '--------------' | '--------------' | '--------------' | '--------------' | '--------------' |
 '----------------' '----------------' '----------------' '----------------' '----------------'
From Match with love.

Hi, welcome to the future (or to the past if you're reading this from the future)

This is an MySQL instance for Asset Journeys, if you want to connect to it with a console mysql client
this command might be useful for you:

   docker run -it mysql/mysql-server:5.7 mysql -hdocker.for.mac.localhost -uassetstates -passetstates assetstates

EOF
