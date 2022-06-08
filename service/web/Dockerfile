FROM python:slim AS tpl
ARG app
ARG sha
ARG ver
ARG date
ENV app ${app}.org
ENV sass dart-sass-1.52.2-linux-x64.tar.gz

ADD https://github.com/sass/dart-sass/releases/download/1.52.2/$sass .
RUN tar -xzf $sass
COPY sass sass
RUN dart-sass/sass sass/main.scss main.css --style compressed --no-source-map

RUN pip3 install jinja2-cli
COPY *.j2.* ./
WORKDIR rendered
RUN jinja2 ../http.j2.conf -D app=$app > http.conf
RUN jinja2 ../index.j2.html -D app=$app -D sha=$sha -D ver=$ver -D date="$date" > index.html
RUN jinja2 ../app.j2.js -D app=$app > app.js

FROM nginx:alpine
WORKDIR /usr/share/nginx/html
ADD https://unpkg.com/vue@3.2.37/dist/vue.runtime.global.prod.js vue.js
RUN chmod +r *.js
COPY --from=tpl main.css main.css
COPY --from=tpl rendered/http.conf /etc/nginx/conf.d/
COPY --from=tpl rendered/index.html .
COPY --from=tpl rendered/app.js .
COPY fonts fonts
COPY images images
COPY favicon.ico favicon.ico
