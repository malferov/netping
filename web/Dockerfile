FROM alpine AS tpl
ARG app
ARG ver
ARG date
ENV app $app.org
RUN apk add py-pip
RUN pip install jinja2-cli
COPY http.conf .
COPY index.html .
RUN jinja2 http.conf -D app=$app > http.conf.rendered
RUN jinja2 index.html -D app=$app -D ver=$ver -D date="$date" > index.html.rendered

FROM nginx:alpine
COPY --from=tpl http.conf.rendered /etc/nginx/conf.d/http.conf
COPY --from=tpl index.html.rendered /usr/share/nginx/html/index.html
COPY styles.css /usr/share/nginx/html/styles.css
