FROM nginx:1.27.0-alpine

RUN rm /etc/nginx/conf.d/*

COPY ./default.conf /etc/nginx/conf.d/
COPY ./local.cert.pem /etc/letsencrypt/live/budgetbuddy.tvgelderen.nl/fullchain.pem
COPY ./local.key.pem /etc/letsencrypt/live/budgetbuddy.tvgelderen.nl/privkey.pem

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]
