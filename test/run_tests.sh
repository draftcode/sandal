#!/bin/bash

tempfile=`mktemp`
trap "rm $tempfile" 0

for filename in `ls`; do
  case $filename in
    (*.sandal)
      expect_filename=${filename%.sandal}.expect
      err_output=`sandal $filename 2>&1 1>$tempfile`
      if [ $? -ne 0 ]; then
          echo "FAILED: $filename"
          echo "    "$err_output
      else
        if [ ! -e $expect_filename ]; then
          echo "FAILED: $filename"
          echo "    No expect file"
        else
          diff_output=`diff -u $expect_filename $tempfile`
          if [ $? -ne 0 ]; then
            echo "FAILED: $filename"
            echo "$diff_output" | awk '{ print "    " $0 }'
          else
            echo "PASSED: $filename"
          fi
        fi
      fi
      ;;
  esac
done
