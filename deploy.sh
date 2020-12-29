#!/bin/sh
cp toplevel/about.html /var/www/doghouse >/dev/null
cp toplevel/home.css /var/www/doghouse >/dev/null
cp toplevel/index.html /var/www/doghouse >/dev/null
cp toplevel/main.css /var/www/doghouse >/dev/null
cp -r articles /var/www/doghouse >/dev/null
cp -r recipes /var/www/doghouse >/dev/null
cp -r img /var/www/doghouse >/dev/null
echo "Deployed"
