#!/bin/bash
src=/go/src/github.com/chroju/nature-remo-cli
dist=pkg

if [[ ! -d "${src}/${dist}" ]]; then
  mkdir "${src}/${dist}"
fi

for dir in ${src}/bin/*; do
  echo ${dir}
  cd ${dir}
  echo "zip -r "${src}/${dist}/$(basename ${dir}).zip" ./*"
  zip -r "${src}/${dist}/$(basename ${dir}).zip" ./*
done