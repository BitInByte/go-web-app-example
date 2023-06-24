#!/bin/bash

http -vv --session=./session.json GET localhost:3000/v1/auth/
