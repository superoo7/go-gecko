#/bin/sh

set -e

Green='\033[0;32m'        # Green
Color_Off='\033[0m'       # Text Reset


for filename in $(ls _example/v3)
do 
    echo "$Green========================"
    echo $filename
    echo "========================$Color_Off"
    go run _example/v3/$filename >/dev/null
    echo "SUCCESSFUL"
done