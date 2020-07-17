# url-shortener-service
Service for URL Shortener in Go


## TODOs:
- [x] Implement Redis
- [x] Implement url checking before creating url 
- [x] Migrate to Echo
- [ ] Main page UI using React/HTML Template
- [ ] Implement OAUTH2 for user loginI
- [ ] Implement Rate-limiter. 
- [ ] Implement concurrent using go chan and workers
- [ ] Add expiration date and implement CRON to automatically delete expired urls. Or lazy delete by checking whenever a user tries to access an expired link.
- [ ] Add visited count
