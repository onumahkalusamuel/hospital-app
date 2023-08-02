#!/bin/sh

cd web 
yarn build
cd ../
cd web && yarn build:watch & air && fg