/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package main

//go:generate sh -c "cd ./ui && npx react-scripts build"
//go:generate sh -c "rm -f ./assets/js/*.js"
//go:generate sh -c "cp ./ui/build/static/js/*.js ./assets/js/"
//go:generate sh -c "rm -f ./assets/css/*.css"
//go:generate sh -c "cp ./ui/build/static/css/*.css ./assets/css/"

