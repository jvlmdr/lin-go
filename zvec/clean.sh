#!/bin/bash

ls | grep '\.go$' | grep -v '_cmplx[_.]' | xargs rm -f
