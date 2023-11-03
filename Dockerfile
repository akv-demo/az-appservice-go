FROM busybox
copy ./az-appservice /az-appservice
CMD ["/az-appservice"]
EXPOSE 8001