#!/bin/sh
cp -rv www/* /var/www/doghouse && echo "Deployed sucessfully!" || echo "Deployment failed :("
