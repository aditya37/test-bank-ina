## PREREQUISITE
### REGISTER ENDPOINT OAUTH2 CALLBACK
- kunjungi halaman google cloud console (https://console.cloud.google.com/apis/credentials)
- kemudian tambahkan credential dengan click menu *+ Crate credentials*
- setalah itu pilih OAUTH2 CLIENT ID dan pilih application type web application
- dan musakan url redirect atau callback di menu Authorized redirect URIs

untuk project ini untuk url redirect-nya adalah http://localhost:2024/users/callback

### SETUP ENVIRONMENT GOOGLE OUATH2
```
OAUTH_REDIRECT="{OAUTH_REDIRECT_URL}"
GOOGLE_CLIENT_ID="{GOOGLE_OAUTH_CLIENT_ID}"
GOOGLE_CLIENT_SECRET="{GOOGLE_OUATH_SECRET}"
```

OAUTH_REDIRECT => url atau endpoint untuk redirect setelah login dengan google, untuk project ini url redirect-nya adalah http://localhost:2024/users/callback

GOOGLE_CLIENT_ID & GOOGLE_CLIENT_SECRET => Key ini di dapat ketika berhasil register oauth di step awal (REGISTER ENDPOINT OAUTH2 CALLBACK)

### Setup env
untuk env yang digunakan, bisa rename file .env.example menjadi .env dan setup sesuai database di mesin anda.

### how to
- register dengan endpoint /users/
- setelah berhasil register lakukan login melalu web browser dengan endpoint 127.0.0.1:2024/users/auth, nanti akan di redirect ke login google
- jika berhasil maka akan ada informasi token dan di redirect ke http://localhost:2024/users/callback
- jika email yg di gunakan tidak terdaftar di service maka tidak mendapatkan access token. 
