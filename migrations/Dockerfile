FROM python:3.7-alpine

# gcc jpeg-dev and zlib-dev required for Pillow (required by some library in requirements.txt)
# gcc libc-dev and linux headers required to install uwsgi
RUN set -e; \
    apk add --no-cache zlib-dev postgresql-dev; \
	apk add --no-cache --virtual .build-deps \
		gcc \
		libc-dev \
		linux-headers \
		musl-dev \
		git \
	; \
    pip install --upgrade pip;
# Install all python dependency libs
# copy requirement only, not whole code so docker can cache it
COPY requirements.txt /app/
RUN LIBRARY_PATH=/lib:/usr/lib /bin/sh -c "pip3 install -r /app/requirements.txt"
COPY . /app/
WORKDIR /app/