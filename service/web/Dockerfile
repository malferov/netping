FROM python:slim AS tpl
ARG app
ARG sha
ARG ver
ARG date
ARG min
ENV app ${app}.org
ENV sass dart-sass-1.26.8-linux-x64.tar.gz

ADD https://github.com/sass/dart-sass/releases/download/1.26.8/$sass .
RUN tar -xzf $sass
COPY sass sass
RUN dart-sass/sass sass/main.scss main.css

RUN pip install jinja2-cli
COPY http.conf index.html app.js ./
RUN jinja2 http.conf -D app=$app > http.conf.rendered
RUN jinja2 index.html -D app=$app -D sha=$sha -D ver=$ver -D date="$date" \
  -D min=$min > index.html.rendered
RUN jinja2 app.js -D app=$app > app.js.rendered

FROM nginx:alpine
ARG min
WORKDIR /usr/share/nginx/html
ADD https://vuejs.org/js/vue.${min}js .
RUN chmod +r *.js
COPY --from=tpl main.css main.css
COPY --from=tpl http.conf.rendered /etc/nginx/conf.d/http.conf
COPY --from=tpl index.html.rendered index.html
COPY --from=tpl app.js.rendered app.js
COPY fonts fonts
COPY images images
COPY favicon.ico favicon.ico
