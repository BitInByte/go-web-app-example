#!/bin/bash

http -vv --session=./session.json POST :3000/v1/todo/ title="Test todo title" body="Test todo body"
