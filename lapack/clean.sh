#!/bin/bash

ls | grep '\.go$' | grep '_cmplx[_.]' | xargs rm -f
