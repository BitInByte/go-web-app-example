#!/bin/bash

http -vv --session=./session.json POST :3000/v1/auth/signup email=test2@test.com username=test2 password=test123
