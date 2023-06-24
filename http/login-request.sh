#!/bin/bash

http -vv --session=./session.json POST :3000/v1/auth/login email=test@test.com username=test password=test123
