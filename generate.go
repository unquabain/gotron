/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package main

//go:generate sh -c "cd ./front-end && npx react-scripts build"
//go:generate sh -c "rm -f ./assets/js/*.js"
//go:generate sh -c "cp ./front-end/build/static/js/*.js ./assets/js/"
//go:generate sh -c "rm -f ./assets/css/*.css"
//go:generate sh -c "cp ./front-end/build/static/css/*.css ./assets/css/"

