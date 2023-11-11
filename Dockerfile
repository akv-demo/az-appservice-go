FROM alpine:3.18
copy ./az-appservice /az-appservice
CMD ["/az-appservice"]
EXPOSE 8001