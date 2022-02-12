docker run -d \
    --net=host \
    -e TZ=Europe/Madrid \
    --name "mailserver" \
    -h "mail.samuran.com" \
    -t analogic/poste.io